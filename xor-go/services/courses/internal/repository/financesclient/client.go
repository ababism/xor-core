package financesclient

import (
	"context"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"net/http"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/domain/keys"
	"xor-go/services/courses/internal/repository/financesclient/generated"
	"xor-go/services/courses/internal/service/adapters"
)

var _ adapters.FinancesClient = Client{}

type Client struct {
	httpDoer *generated.ClientWithResponses
}

func (c Client) RegisterProducts(initialCtx context.Context, products []domain.Product) ([]domain.Product, error) {
	_ = zapctx.Logger(initialCtx)

	tr := otel.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "driver/repository/financesClient.RegisterProducts")
	defer span.End()

	productsCreate := make([]generated.ProductCreate, len(products))
	for i, product := range products {
		productsCreate[i] = ToProductCreateRequest(product)
	}

	requestID, ok := GetRequestIDFromContext(ctx)

	// Inject the trace context into the request's headers to send trace
	reqEditor := func(ctx context.Context, req *http.Request) error {
		otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
		if ok {
			req.Header.Set(keys.KeyRequestID, requestID)
		}
		return nil
	}

	_, _ = c.httpDoer.PostProductsListWithResponse(ctx, productsCreate, reqEditor)
	//if err != nil {
	//	logger.Error("error while creating register products request:", zap.Error(err))
	//	return nil, xapperror.New(http.StatusInternalServerError,
	//		"error while creating register products request",
	//		"error while creating register products to finance microservice", err)
	//}

	//logger.Info("register products response", zap.Int("status", resp.HTTPResponse.StatusCode))
	//logger.Info("register products response", zap.ByteString("res", resp.Body))
	//logger.Info("register products response", zap.Any("res", resp))
	//
	//if resp.HTTPResponse.StatusCode == http.StatusOK {
	//	 productsResponse := make([]domain.Product, 0)
	//
	//	err = json.Unmarshal(resp.Body, &productsResponse)
	//	if err != nil {
	//		logger.Error("error while decoding products JSON:", zap.Error(err))
	//		return nil, err
	//	}
	//	return productsResponse, nil
	//} else {
	//	var productsErrorMessage Error
	//	err = json.Unmarshal(resp.Body, &productsErrorMessage)
	//	if err != nil {
	//		logger.Error("error while decoding products error message JSON:", zap.Error(err))
	//		return nil, err
	//	}
	//	logger.Error("can't create register products ended:", zap.Int("status", resp.HTTPResponse.StatusCode), zap.Error(productsErrorMessage))
	//	return nil, xapperror.New(http.StatusInternalServerError,
	//		"error while creating register products", "error while creating register products microservice", productsErrorMessage)
	//}

	for i, _ := range products {
		products[i].ID = uuid.New()
	}
	return products, nil
}

func NewClient(client *generated.ClientWithResponses) *Client {
	return &Client{httpDoer: client}
}

func GetRequestIDFromContext(ctx context.Context) (string, bool) {
	requestID, ok := ctx.Value(keys.KeyRequestID).(string)
	return requestID, ok
}

//// TODO SendTrace
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
