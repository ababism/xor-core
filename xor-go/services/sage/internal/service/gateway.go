package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"regexp"
	"strings"
	idmproto "xor-go/proto/idm"
	"xor-go/services/sage/internal/config"
	"xor-go/services/sage/internal/domain"
	"xor-go/services/sage/internal/service/adapter"
)

const xorAuthorizationHeader = "Xor-Authorization"
const xorRequestUuidHeader = "Xor-Request-Uuid"
const xorAuthorizationHeaderAuthorized = "AUTHORIZED"
const xorAuthorizationHeaderNotAuthorized = "NOT_AUTHORIZED"

var _ adapter.GatewayService = &gatewayService{}

type gatewayService struct {
	resourceToConfig map[string]*config.ResourceConfig
	restyClient      *resty.Client
	idmGrpcClient    idmproto.IdmClient
}

func NewGatewayResourceService(
	resourceToConfig map[string]*config.ResourceConfig,
	idmGrpcClient idmproto.IdmClient,
) adapter.GatewayService {
	return &gatewayService{
		resourceToConfig: resourceToConfig,
		restyClient:      resty.New(),
		idmGrpcClient:    idmGrpcClient,
	}
}

func (r *gatewayService) PassSecure(
	ctx context.Context,
	passSecureResourceRequest *domain.PassSecureResourceRequest,
) (*domain.InternalResourceResponse, error) {
	resourceConfig, ok := r.resourceToConfig[passSecureResourceRequest.Resource]
	if !ok {
		return nil, fmt.Errorf("failed to find config for resource=%v", passSecureResourceRequest.Resource)
	}

	xorHeaders := map[string]string{
		xorRequestUuidHeader:   passSecureResourceRequest.RequestUuid.String(),
		xorAuthorizationHeader: xorAuthorizationHeaderAuthorized,
		"Xor-Account-Uuid":     passSecureResourceRequest.AccountUuid.String(),
		"Xor-Account-Email":    passSecureResourceRequest.AccountEmail,
		"Xor-Account-Roles":    strings.Join(passSecureResourceRequest.Roles, ","),
	}

	var internalResourceResponse domain.InternalResourceResponse
	restyRequest := r.restyClient.R().
		SetHeaders(xorHeaders).
		SetBody(passSecureResourceRequest.Body).
		SetResult(&internalResourceResponse.Body)

	restyResponse, err := executeRestyRequest(
		restyRequest,
		passSecureResourceRequest.Method,
		resourceConfig.Host+passSecureResourceRequest.Route,
	)
	if err != nil {
		return nil, err
	}

	internalResourceResponse.Status = restyResponse.StatusCode()

	return &internalResourceResponse, nil
}

func (r *gatewayService) Verify(
	ctx context.Context,
	passSecureResourceInfo *domain.PassSecureResourceInfo,
) (*domain.IdmVerifyResponse, error) {
	verifyRequest := &idmproto.VerifyRequest{AccessToken: passSecureResourceInfo.AccessToken}
	verifyResponse, err := r.idmGrpcClient.Verify(ctx, verifyRequest)
	if err != nil {
		return nil, err
	}

	assignedRoles := make(map[string]bool)
	for _, role := range verifyResponse.Roles {
		assignedRoles[role] = true
	}

	requiredRoles, err := r.getRequiredRoles(passSecureResourceInfo.Resource, passSecureResourceInfo.Route)
	if err != nil {
		return nil, err
	}

	for role, _ := range requiredRoles {
		if _, ok := assignedRoles[role]; !ok {
			return nil, fmt.Errorf("required role=%v is not assigned", role)
		}
	}

	assignedRolesSlice := make([]string, len(assignedRoles))
	for role, _ := range requiredRoles {
		assignedRolesSlice = append(assignedRolesSlice, role)
	}

	parsedAccountUuid, err := uuid.Parse(verifyResponse.AccountUuid)
	if err != nil {
		return nil, err
	}
	return &domain.IdmVerifyResponse{
		AccountUuid:  parsedAccountUuid,
		AccountEmail: verifyResponse.AccountEmail,
		Roles:        assignedRolesSlice,
	}, nil
}

func (r *gatewayService) getRequiredRoles(resourceName string, resourceRoute string) (map[string]bool, error) {
	resourceConfig, ok := r.resourceToConfig[resourceName]
	if !ok {
		return nil, fmt.Errorf("failed to resolve resource=%v in resources config", resourceName)
	}

	matchedCount := 0
	var matchedRoute config.ResourceRoute
	for _, route := range resourceConfig.Routes {
		matched, err := regexp.MatchString(route.Pattern, resourceRoute)
		if err != nil {
			return nil, err
		}
		if matched {
			matchedCount += 1
			matchedRoute = route
		}
		if matchedCount > 1 {
			return nil, fmt.Errorf(
				"ambiguous routes match for resource=%v, route=%v",
				resourceName,
				resourceRoute,
			)
		}
	}

	if matchedCount == 0 {
		return nil, fmt.Errorf("failed to match any route for resource=%v, route=%v",
			resourceName,
			resourceRoute,
		)
	}
	return getRequiredRoles(matchedRoute), nil
}

func getRequiredRoles(resourceRoute config.ResourceRoute) map[string]bool {
	requiredRoles := make(map[string]bool)
	for _, role := range resourceRoute.RequiredRoles {
		requiredRoles[role] = true
	}
	return requiredRoles
}

func executeRestyRequest(request *resty.Request, method string, url string) (*resty.Response, error) {
	if method == "GET" || method == "POST" || method == "UPDATE" || method == "DELETE" || method == "PATCH" {
		return request.Execute(method, url)
	}
	return nil, fmt.Errorf("failed to execute resty request with unsupported http method=%v", method)
}

func (r *gatewayService) PassInsecure(
	ctx context.Context,
	passInsecureResourceRequest *domain.PassInsecureResourceRequest,
) (*domain.InternalResourceResponse, error) {
	requiredRoles, err := r.getRequiredRoles(passInsecureResourceRequest.Resource, passInsecureResourceRequest.Route)
	if err != nil {
		return nil, err
	}

	if len(requiredRoles) != 0 {
		return nil, errors.New("resource route requires IDM verification")
	}

	resourceConfig, ok := r.resourceToConfig[passInsecureResourceRequest.Resource]
	if !ok {
		return nil, fmt.Errorf("failed to find config for resource=%v", passInsecureResourceRequest.Resource)
	}

	xorHeaders := map[string]string{
		xorRequestUuidHeader:   passInsecureResourceRequest.RequestUuid.String(),
		xorAuthorizationHeader: xorAuthorizationHeaderNotAuthorized,
	}

	var internalResourceResponse domain.InternalResourceResponse
	restyRequest := r.restyClient.R().
		SetHeaders(xorHeaders).
		SetBody(passInsecureResourceRequest.Body).
		SetResult(&internalResourceResponse.Body)
	resp, err := executeRestyRequest(
		restyRequest,
		passInsecureResourceRequest.Method,
		resourceConfig.Host+passInsecureResourceRequest.Route,
	)
	if err != nil {
		return nil, err
	}

	internalResourceResponse.Status = resp.StatusCode()
	if resp.StatusCode() < 500 {
		var parsedResponseBody map[string]any
		err = json.Unmarshal([]byte(resp.String()), &parsedResponseBody)
		if err != nil {
			return nil, err
		}
		internalResourceResponse.Body = parsedResponseBody
	}

	return &internalResourceResponse, nil
}
