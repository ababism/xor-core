package purchaseclient

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
	"net/http"
	"time"
	"xor-go/pkg/xapperror"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/domain/keys"
	"xor-go/services/courses/internal/repository/purchaseclient/generated"
	"xor-go/services/courses/internal/service/adapters"
)

var _ adapters.PurchaseClient = Client{}

type Client struct {
	httpDoer *generated.ClientWithResponses
}

func NewClient(client *generated.ClientWithResponses) *Client {
	return &Client{httpDoer: client}
}

func (c Client) CreatePurchase(initialCtx context.Context, products []domain.Product) (domain.PaymentRedirect, error) {
	logger := zapctx.Logger(initialCtx)

	tr := otel.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "driver/repository/purchaseClient.CreatePurchase")
	defer span.End()

	productIDs := make([]uuid.UUID, len(products))
	for i, product := range products {
		productIDs[i] = product.ID
	}

	req := generated.PurchaseRequestCreate{
		CreatedAt: time.Now(),
		Products:  productIDs,
		// TODO
		Receiver:   nil,
		Sender:     nil,
		WebhookURL: "",
	}

	requestID, ok := GetRequestIDFromContext(ctx)

	// Inject trace context in request's header to send trace
	reqEditor := func(ctx context.Context, req *http.Request) error {
		otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
		if ok {
			req.Header.Set(keys.KeyRequestID, requestID)
		}
		return nil
	}

	resp, err := c.httpDoer.CreateWithResponse(ctx, req, reqEditor)
	if err != nil {
		logger.Error("error while creating purchase request:", zap.Error(err))
		return domain.PaymentRedirect{}, xapperror.New(http.StatusInternalServerError,
			"error while creating purchase request", "error while creating purchase to finance microservice", err)
	}

	if resp.HTTPResponse.StatusCode == http.StatusOK {
		var paymentRedirect domain.PaymentRedirect
		//err = resp.DecodeJSON(&paymentRedirect)
		err = json.Unmarshal(resp.Body, &paymentRedirect)
		if err != nil {
			logger.Error("error while decoding payment redirect JSON:", zap.Error(err))
			return domain.PaymentRedirect{}, err
		}
		return paymentRedirect, nil
	} else {
		var paymentErrorMessage Error
		err = json.Unmarshal(resp.Body, &paymentErrorMessage)
		if err != nil {
			logger.Error("error while decoding payment error message JSON:", zap.Error(err))
			return domain.PaymentRedirect{}, err
		}
		logger.Error("can't create purchase ended:", zap.Int("status", resp.HTTPResponse.StatusCode), zap.Error(paymentErrorMessage))
		return domain.PaymentRedirect{}, xapperror.New(http.StatusInternalServerError,
			"error while creating purchase", "error while creating purchase microservice", paymentErrorMessage)
	}
}

func GetRequestIDFromContext(ctx context.Context) (string, bool) {
	requestID, ok := ctx.Value(keys.KeyRequestID).(string)
	return requestID, ok
}

//func (c Client) GetDrivers(ctx context.Context, driverLocation domain.LatLngLiteral, radius float32) ([]domain.DriverLocation, error) {
//	logger := zapctx.Logger(ctx)
//
//	tr := global.Tracer(domain.ServiceName)
//	newCtx, span := tr.Start(ctx, "driver/repository/locationclient.GetDrivers")
//	defer span.End()
//
//	requestID, ok := GetRequestIDFromContext(newCtx)
//	var (
//		resp *generated.GetDriversResponse
//		err  error
//	)
//
//	if ok {
//		span.AddEvent("passed requestId for GetDrivers handler from location service",
//			trace.WithAttributes(attribute.String(domain.KeyRequestID, requestID)))
//
//		reqEditor := func(newCtx context.Context, req *http.Request) error {
//			req.Header.Set(domain.KeyRequestID, requestID)
//			return nil
//		}
//		resp, err = c.httpDoer.GetDriversWithResponse(newCtx, &generated.GetDriversParams{
//			Lat:    driverLocation.Lat,
//			Lng:    driverLocation.Lng,
//			Radius: radius,
//		}, reqEditor)
//	} else {
//		logger.Error("can't find RequestID in ctx")
//		resp, err = c.httpDoer.GetDriversWithResponse(newCtx, &generated.GetDriversParams{
//			Lat:    driverLocation.Lat,
//			Lng:    driverLocation.Lng,
//			Radius: radius,
//		})
//	}
//	if err != nil {
//		logger.Error("error while getting drivers from location service:", zap.Error(err))
//		return nil, err
//	}
//
//	var locationErrorMessage Error
//	if resp.HTTPResponse.StatusCode != http.StatusOK {
//		if err = json.Unmarshal(resp.Body, &locationErrorMessage); err != nil {
//			logger.Error("error while decoding location error message JSON:", zap.Error(err))
//			return nil, err
//		}
//		logger.Error("can't get drivers from location service ended:", zap.Int("status", resp.HTTPResponse.StatusCode), zap.Error(locationErrorMessage))
//		return nil, locationErrorMessage
//	}
//
//	//var driverLocations GetDriversResponse
//	//var driverLocations []generated.Driver
//	type GetDriversResponse struct {
//		Drivers []generated.Driver `json:"drivers"`
//	}
//	var response GetDriversResponse
//
//	err = json.Unmarshal(resp.Body, &response)
//	if err != nil {
//		logger.Error("error while decoding driver location JSON:", zap.Error(err))
//		return nil, err
//	}
//
//	res, err := ToDriverLocationsDomain(response.Drivers)
//	if err != nil {
//		logger.Error("error while converting driver locations to domain:", zap.Error(err))
//		return nil, err
//	}
//	return res, nil
//}
