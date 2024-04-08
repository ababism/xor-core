package financesclient

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/repository/financesclient/generated"
	"xor-go/services/courses/internal/service/adapters"
)

var _ adapters.FinancesClient = Client{}

type Client struct {
	httpDoer *generated.ClientWithResponses
}

func (c Client) RegisterProducts(ctx context.Context, lessonIDs []uuid.UUID, teacherID uuid.UUID, price float64) ([]domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (c Client) CreatePurchase(ctx context.Context, productIDs []domain.Product) (domain.PaymentRedirect, error) {
	//TODO implement me
	panic("implement me")
}

func NewClient(client *generated.ClientWithResponses) *Client {
	return &Client{httpDoer: client}
}

func GetRequestIDFromContext(ctx context.Context) (string, bool) {
	requestID, ok := ctx.Value(domain.KeyRequestID).(string)
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
