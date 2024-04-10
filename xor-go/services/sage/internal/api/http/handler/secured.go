package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"xor-go/pkg/xhttp/response"
	"xor-go/services/sage/internal/api/http/model"
	"xor-go/services/sage/pkg/idm"
)

//var _ ServerInterface = &Handler{}

type SecuredHandler struct {
	logger    *zap.Logger
	responser *response.HttpResponseWrapper
	idmClient *idm.Client
	//idmService
}

//	type Handler struct {
//		productService adapters.ProductService
//	}
func NewSecuredHandler(logger *zap.Logger, responser *response.HttpResponseWrapper, client *idm.Client) *SecuredHandler {
	return &SecuredHandler{
		logger:    logger,
		responser: responser,
		idmClient: client,
	}
}

func (r *SecuredHandler) InitRoutes(g *gin.RouterGroup) {
	account := g.Group("/secured")
	//account.GET("/list", r.List)
	account.POST("/request", r.Request)
	//account.PUT("/update-password", r.UpdatePassword)
	//account.PUT("/deactivate/:uuid", r.Deactivate)
}

func (r *SecuredHandler) Request(ctx *gin.Context) {
	var accessRequest model.SecuredAccessRequest
	if err := ctx.BindJSON(&accessRequest); err != nil {
		r.responser.HandleErrorWithMessage(ctx, http.StatusBadRequest, err)
		return
	}

	vR := &idm.VerifyRequest{AccessToken: accessRequest.AccessToken}
	verifyResponse, err := r.idmClient.Verify(vR)
	if err != nil {
		r.logger.Error("Failed verification", zap.String("access_token", accessRequest.AccessToken))
		r.logger.Debug(err.Error())
		ctx.AbortWithStatus(500)
		//r.responser.HandleXorErrorWithMessage(ctx, err)
		return
	}

	fmt.Println(verifyResponse)

	var totalResp model.SecuredAccessResponse

	systemHeaders := map[string]string{
		"Xor-Account-Uuid":  verifyResponse.AccountUuid.String(),
		"Xor-Account-Email": verifyResponse.AccountEmail,
		"Xor-Account-Roles": strings.Join(verifyResponse.Roles, ","),
	}

	restyClient := resty.New()
	resp, err := restyClient.R().
		SetBody(accessRequest.Body).
		SetResult(&totalResp.Body).
		SetHeaders(systemHeaders).
		Get("http://localhost:8478" + accessRequest.ApiUrl)

	if err != nil {
		r.responser.HandleXorErrorWithMessage(ctx, err)
		return
	}

	if resp.IsError() {
		//return nil, errors.New(resp.Error())
		r.responser.HandleXorErrorWithMessage(ctx, errors.New(resp.Error().(string)))
		return
	}
	totalResp.Status = resp.StatusCode()

	ctx.JSON(200, totalResp)

	//
	//fmt.Println(accessRequest)
	//
	//vR := &idm.VerifyRequest{AccessToken: "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJmb3hsZXJlbkB5YW5kZXgucnUiLCJpYXQiOjE3MTI3MjA4OTMsImV4cCI6MTcxMjgwNzI5M30.y6dbguGqVf_OsfiFASx7wTWneje8ErdPIaD2HIGC6PM"}
	//response, err := r.IdmClient.Verify(vR)
	//if err != nil {
	//	fmt.Println(err)
	//	ctx.AbortWithStatus(400)
	//	return
	//}
	//fmt.Println(response)
	//
	//ctx.JSON(http.StatusOK, response)

	// Адрес хоста, на который будет перенаправлен запрос
	//targetHost := "http://localhost:8478"

	// Создаем новый HTTP клиент
	//httpClient := &http.Client{}
	//
	//requestUrl := accessRequest.ApiUrl
	//
	//// Создаем копию оригинального запроса
	//req, err := http.NewRequest("GET", requestUrl, ctx.Request.Body)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//// Копируем заголовки запроса
	//for k, v := range ctx.Request.Header {
	//	req.Header.Set(k, v[0])
	//}
	//
	//// Выполняем запрос к целевому хосту
	//resp, err := httpClient.Do(req)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//defer resp.Body.Close()
	//
	//// Копируем заголовки ответа
	//for k, v := range resp.Header {
	//	ctx.Header(k, v[0])
	//}
	//
	//// Копируем статус ответа
	//ctx.Status(resp.StatusCode)
	//
	//var targetResponse model.SecuredAccessResponse
	//err = json.NewDecoder(resp.Body).Decode(&targetResponse.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response body"})
	//	return
	//}
	//
	//// Отправляем структуру TargetResponse клиенту
	//ctx.JSON(resp.StatusCode, targetResponse)

	//---------

	// Считываем тело ответа от целевого хоста и отправляем его клиенту
	//responseBody, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response accessRequest"})
	//	return
	//}

	//ctx.Writer.Write(responseBody)

	//// Копируем тело ответа
	//_, err = ctx.Writer.Write([]byte{})
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}

	//ctx.JSON(http.StatusOK, resp.Body)

	//req, err := http.NewRequest(ctx.Request.Method, targetHost+ctx.Request.URL.String(), ctx.Request.Body)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//_, newCtx, span := getPaymentTracerSpan(ctx, ".Request")
	//defer span.End()
	//
	//var accessRequest *PaymentFilter
	//if err := ctx.BindJSON(&accessRequest); err != nil && err != io.EOF {
	//	http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
	//	return
	//}
	//
	//domains, err := r.paymentService.List(newCtx, FilterToDomain(accessRequest))
	//if err != nil {
	//	http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
	//	return
	//}
	//
	//list := make([]PaymentGet, len(domains))
	//for i, item := range domains {
	//	list[i] = DomainToGet(item)
	//}
	//
	//ctx.JSON(http.StatusOK, list)
}
