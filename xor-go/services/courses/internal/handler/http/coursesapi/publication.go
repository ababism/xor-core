package coursesapi

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	global "go.opentelemetry.io/otel"
	"net/http"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/handler/generated"
	"xor-go/services/courses/internal/handler/http/models"
)

func (h CoursesHandler) PostPublicationRequests(ginCtx *gin.Context, params generated.PostPublicationRequestsParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.PostPublicationRequests")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	var payload generated.PublicationRequest
	h.bindRequestBody(ginCtx, &payload)

	publication, err := h.coursesService.RequestCoursePublication(ctx, params.Actor.ToDomain(), payload.ToDomain())
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	resp := models.ToPublicationRequestResponse(publication)

	ginCtx.JSON(http.StatusOK, resp)
}

func (h CoursesHandler) PutPublicationRequestsRequestID(ginCtx *gin.Context, requestID uuid.UUID, params generated.PutPublicationRequestsRequestIDParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.PutPublicationRequestsRequestID")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	var payload generated.PublicationRequest
	h.bindRequestBody(ginCtx, &payload)

	payload.ID = requestID

	publication, err := h.coursesService.UpdatePublicationRequest(ctx, params.Actor.ToDomain(), payload.ToDomain())
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	resp := models.ToPublicationRequestResponse(publication)

	ginCtx.JSON(http.StatusOK, resp)
}
