package driver_api

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/zaputil/zapctx"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"gitlab.com/ArtemFed/mts-final-taxi/projects/template/internal/domain"
	"gitlab.com/ArtemFed/mts-final-taxi/projects/template/internal/handler/generated"
	"gitlab.com/ArtemFed/mts-final-taxi/projects/template/internal/handler/http/models"
	"gitlab.com/ArtemFed/mts-final-taxi/projects/template/internal/service/adapters"
	global "go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var _ generated.ServerInterface = &DriverHandler{}

type DriverHandler struct {
	logger        *zap.Logger
	driverService adapters.DriverService
	WaitTimeout   time.Duration
}

func NewDriverHandler(logger *zap.Logger, driverService adapters.DriverService, socketTimeout time.Duration) *DriverHandler {
	return &DriverHandler{logger: logger, driverService: driverService, WaitTimeout: socketTimeout}
}

func (h *DriverHandler) GetTripByID(ginCtx *gin.Context, tripId openapi_types.UUID, params generated.GetTripByIDParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "driver/driver_api.GetTripByID")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	trip, err := h.driverService.GetTripByID(ctx, params.UserId, tripId)
	if err != nil {
		AbortWithBadResponse(ginCtx, h.logger, MapErrorToCode(err), err)
		return
	}
	resp := models.ToTripResponse(*trip)

	ginCtx.JSON(http.StatusOK, resp)
}

func (h *DriverHandler) AcceptTrip(ginCtx *gin.Context, tripId openapi_types.UUID, params generated.AcceptTripParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "driver/driver_api.AcceptTrip")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	err := h.driverService.AcceptTrip(ctx, params.UserId, tripId)
	if err != nil {
		AbortWithBadResponse(ginCtx, h.logger, MapErrorToCode(err), err)
		return
	}

	ginCtx.JSON(http.StatusOK, http.NoBody)
}

func (h *DriverHandler) GetTrips(c *gin.Context, params generated.GetTripsParams) {
	//TODO implement me
	panic("implement me")
}

func (h *DriverHandler) CancelTrip(c *gin.Context, tripId openapi_types.UUID, params generated.CancelTripParams) {
	//TODO implement me
	panic("implement me")
}

func (h *DriverHandler) EndTrip(c *gin.Context, tripId openapi_types.UUID, params generated.EndTripParams) {
	//TODO implement me
	panic("implement me")
}

func (h *DriverHandler) StartTrip(c *gin.Context, tripId openapi_types.UUID, params generated.StartTripParams) {
	//TODO implement me
	panic("implement me")
}
