package payments

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"net/http"
	"xor-go/pkg/xcommon"
	"xor-go/services/finances/internal/domain"
	"xor-go/services/finances/internal/log"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanPaymentsDefault = "payment/client"
)

var _ adapters.PaymentsClient = &paymentsClient{}

type paymentsClient struct {
	httpDoer *ClientWithResponses
}

func NewPaymentsClient(client *ClientWithResponses) adapters.PaymentsClient {
	return &paymentsClient{httpDoer: client}
}

func CtxRequestID(ctx context.Context) (string, bool) {
	requestID, ok := ctx.Value(domain.KeyRequestID).(string)
	return requestID, ok
}

func (c *paymentsClient) GetStatus(ctx context.Context, id uuid.UUID) (*domain.PaymentsGetStatus, error) {
	tr := global.Tracer(adapters.ServiceNamePayment)
	_, span := tr.Start(ctx, spanPaymentsDefault+".GetStatus")
	defer span.End()

	requestID, ok := CtxRequestID(ctx)
	var (
		resp *GetStatusResponse
		err  error
	)
	if ok {
		span.AddEvent(
			"passed requestId for GetStatus handler from Payments Service",
			trace.WithAttributes(attribute.String(domain.KeyRequestID, requestID)),
		)

		reqEditor := func(newCtx context.Context, req *http.Request) error {
			req.Header.Set(domain.KeyRequestID, requestID)
			return nil
		}
		resp, err = c.httpDoer.GetStatusWithResponse(ctx, GetStatusJSONRequestBody{PaymentId: id}, reqEditor)
	} else {
		log.Logger.Error("can't find RequestID in ctx")
		resp, err = c.httpDoer.GetStatusWithResponse(ctx, GetStatusJSONRequestBody{PaymentId: id})
	}
	if err != nil {
		log.Logger.Error("error while getting drivers from Payments Service:", zap.Error(err))
		return nil, err
	}

	var paymentsErrorMessage Error
	if resp.HTTPResponse.StatusCode != http.StatusOK {
		if err = json.Unmarshal(resp.Body, &paymentsErrorMessage); err != nil {
			log.Logger.Error("error while decoding Payments Service error message JSON:", zap.Error(err))
			return nil, err
		}
		log.Logger.Error(
			"can't get payment status from Payments Service ended:", zap.Int("status", resp.HTTPResponse.StatusCode),
			zap.Error(paymentsErrorMessage),
		)
		return nil, paymentsErrorMessage
	}

	var response GetStatusResponse
	err = json.Unmarshal(resp.Body, &response)
	if err != nil {
		log.Logger.Error("error while decoding status from Payments Service JSON:", zap.Error(err))
		return nil, err
	}

	return &domain.PaymentsGetStatus{
		UUID:   *response.JSON200.Id,
		Status: *response.JSON200.Status,
	}, nil
}

func (c *paymentsClient) CreatePurchase(ctx context.Context, purchase *domain.PaymentsCreatePurchase) error {
	tr := global.Tracer(adapters.ServiceNamePayment)
	_, span := tr.Start(ctx, spanPaymentsDefault+".CreatePurchase")
	defer span.End()

	requestID, ok := CtxRequestID(ctx)
	var (
		resp *CreatePurchaseRequestResponse
		err  error
	)
	if ok {
		span.AddEvent(
			"passed requestId for CreatePurchase handler from Payments Service",
			trace.WithAttributes(attribute.String(domain.KeyRequestID, requestID)),
		)

		reqEditor := func(newCtx context.Context, req *http.Request) error {
			req.Header.Set(domain.KeyRequestID, requestID)
			return nil
		}
		resp, err = c.httpDoer.CreatePurchaseRequestWithResponse(ctx, CreatePurchaseRequestJSONRequestBody{}, reqEditor)
	} else {
		log.Logger.Error("can't find RequestID in ctx")
		resp, err = c.httpDoer.GetStatusWithResponse(ctx, GetStatusJSONRequestBody{PaymentId: id})
	}
	if err != nil {
		log.Logger.Error("error while getting drivers from Payments Service:", zap.Error(err))
		return nil, err
	}

	var paymentsErrorMessage Error
	if resp.HTTPResponse.StatusCode != http.StatusOK {
		if err = json.Unmarshal(resp.Body, &paymentsErrorMessage); err != nil {
			log.Logger.Error("error while decoding Payments Service error message JSON:", zap.Error(err))
			return nil, err
		}
		log.Logger.Error(
			"can't get payment status from Payments Service ended:", zap.Int("status", resp.HTTPResponse.StatusCode),
			zap.Error(paymentsErrorMessage),
		)
		return nil, paymentsErrorMessage
	}

	var response GetStatusResponse
	err = json.Unmarshal(resp.Body, &response)
	if err != nil {
		log.Logger.Error("error while decoding status from Payments Service JSON:", zap.Error(err))
		return nil, err
	}

	return &domain.PaymentsGetStatus{
		UUID:   *response.JSON200.Id,
		Status: *response.JSON200.Status,
	}, nil
}

func (c *paymentsClient) CreatePayout(ctx context.Context, purchase *domain.PaymentsCreatePayout) error {
	//TODO implement me
	panic("implement me")
}

func (c *paymentsClient) Get(ctx context.Context, filter *domain.PaymentFilter) (*domain.PaymentGet, error) {
	tr := global.Tracer(adapters.ServiceNamePayment)
	_, span := tr.Start(ctx, spanPaymentsDefault+".Get")
	defer span.End()

	accounts, err := c.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	return xcommon.EnsureSingle(accounts)
}
