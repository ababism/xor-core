package payments

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"net/http"
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
		log.Logger.Error("error while getting status from Payments Service:", zap.Error(err))
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

	return &domain.PaymentsGetStatus{
		UUID:   *resp.JSON200.Id,
		Status: *resp.JSON200.Status,
	}, nil
}

func (c *paymentsClient) CreatePurchase(
	ctx context.Context,
	purchase *domain.PaymentsCreatePurchase,
) (*domain.CreatePurchaseResponse, error) {
	tr := global.Tracer(adapters.ServiceNamePayment)
	_, span := tr.Start(ctx, spanPaymentsDefault+".CreatePurchase")
	defer span.End()

	requestID, ok := CtxRequestID(ctx)
	var (
		resp *PostPurchaseResponse
		err  error
	)
	if ok {
		span.AddEvent(
			"passed requestId for CreatePurchase handler from Payments Service",
			trace.WithAttributes(attribute.String(domain.KeyRequestID, requestID)),
		)

		// Inject the trace context into the request's headers to send trace
		reqEditor := func(ctx context.Context, req *http.Request) error {
			global.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
			if ok {
				req.Header.Set(domain.KeyRequestID, requestID)
			}
			return nil
		}
		resp, err = c.httpDoer.PostPurchaseWithResponse(ctx, convertToCreatePurchase(*purchase), reqEditor)
	} else {
		log.Logger.Error("can't find RequestID in ctx")
		resp, err = c.httpDoer.PostPurchaseWithResponse(ctx, convertToCreatePurchase(*purchase))
	}
	if err != nil {
		log.Logger.Error("error while creating purchase from Payments Service:", zap.Error(err))
		return nil, err
	}

	var paymentsErrorMessage Error
	if resp.HTTPResponse.StatusCode != http.StatusOK {
		if err = json.Unmarshal(resp.Body, &paymentsErrorMessage); err != nil {
			log.Logger.Error("error while decoding Payments Service error message JSON:", zap.Error(err))
			return nil, err
		}
		log.Logger.Error(
			"can't create purchase from Payments Service ended:", zap.Int("status", resp.HTTPResponse.StatusCode),
			zap.Error(paymentsErrorMessage),
		)
		return nil, paymentsErrorMessage
	}

	response := resp.JSON200
	return &domain.CreatePurchaseResponse{
		Amount: domain.Amount{
			Currency: response.Amount.Currency,
			Value:    response.Amount.Value,
		},
		Confirmation: domain.Confirmation{
			ConfirmationUrl: response.Confirmation.ConfirmationUrl,
			ReturnUrl:       response.Confirmation.ReturnUrl,
			Type:            response.Confirmation.Type,
		},
		CreatedAt:   response.CreatedAt,
		Description: response.Description,
		Id:          response.Id,
		Metadata:    response.Metadata,
		Paid:        response.Paid,
		PaymentMethod: domain.PaymentMethod{
			Id:    response.PaymentMethod.Id,
			Saved: response.PaymentMethod.Saved,
			Type:  response.PaymentMethod.Type,
		},
		Recipient: domain.Recipient{
			AccountId: response.Recipient.AccountId,
			GatewayId: response.Recipient.GatewayId,
		},
		Refundable: response.Refundable,
		Status:     response.Status,
		Test:       response.Test,
	}, nil
}

func (c *paymentsClient) CreatePayout(
	ctx context.Context,
	payout *domain.PaymentsCreatePayout,
) (*domain.CreatePayoutResponse, error) {
	tr := global.Tracer(adapters.ServiceNamePayment)
	_, span := tr.Start(ctx, spanPaymentsDefault+".CreatePayout")
	defer span.End()

	requestID, ok := CtxRequestID(ctx)
	var (
		resp *PostPayoutResponse
		err  error
	)
	if ok {
		span.AddEvent(
			"passed requestId for CreatePayout handler from Payments Service",
			trace.WithAttributes(attribute.String(domain.KeyRequestID, requestID)),
		)

		reqEditor := func(newCtx context.Context, req *http.Request) error {
			req.Header.Set(domain.KeyRequestID, requestID)
			return nil
		}
		resp, err = c.httpDoer.PostPayoutWithResponse(ctx, convertToCreatePayout(*payout), reqEditor)
	} else {
		log.Logger.Error("can't find RequestID in ctx")
		resp, err = c.httpDoer.PostPayoutWithResponse(ctx, convertToCreatePayout(*payout))
	}
	if err != nil {
		log.Logger.Error("error while creating payout from Payments Service:", zap.Error(err))
		return nil, err
	}

	var paymentsErrorMessage Error
	if resp.HTTPResponse.StatusCode != http.StatusOK {
		if err = json.Unmarshal(resp.Body, &paymentsErrorMessage); err != nil {
			log.Logger.Error("error while decoding Payments Service error message JSON:", zap.Error(err))
			return nil, err
		}
		log.Logger.Error(
			"can't create payout from Payments Service ended:", zap.Int("status", resp.HTTPResponse.StatusCode),
			zap.Error(paymentsErrorMessage),
		)
		return nil, paymentsErrorMessage
	}

	response := resp.JSON200
	return &domain.CreatePayoutResponse{
		Amount: domain.Amount{
			Currency: response.Amount.Currency,
			Value:    response.Amount.Value,
		},
		CreatedAt:   response.CreatedAt,
		Description: response.Description,
		Id:          response.Id,
		Metadata:    response.Metadata,
		Status:      response.Status,
		Test:        response.Test,
	}, nil
}
