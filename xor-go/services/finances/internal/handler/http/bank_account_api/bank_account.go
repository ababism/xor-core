package bank_account_api

import (
	"github.com/gin-gonic/gin"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/context"
	"net/http"
	bankaccount "xor-go/services/finances/internal/handler/generated/bank-account"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanDefaultBankAccount = "bank-account/handler."
)

var _ bankaccount.ServerInterface = &BankAccountHandler{}

type BankAccountHandler struct {
	bankAccountService adapters.BankAccountService
}

func NewBankAccountHandler(bankAccountService adapters.BankAccountService) *BankAccountHandler {
	return &BankAccountHandler{bankAccountService: bankAccountService}
}

func getAccountTracerSpan(ctx *gin.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNameBankAccount)
	newCtx, span := tr.Start(ctx, spanDefaultBankAccount+name)
	return tr, newCtx, span
}

func (h *BankAccountHandler) GetList(c *gin.Context) {
	_, newCtx, span := getAccountTracerSpan(c, ".GetList")
	defer span.End()

	trip, err := h.bankAccountService.List(newCtx, params.UserId)
	if err != nil {
		AbortWithBadResponse(ginCtx, h.logger, MapErrorToCode(err), err)
		return
	}
	resp := models.ToTripResponse(*trip)

	ginCtx.JSON(http.StatusOK, resp)
}

func (h *BankAccountHandler) Create(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *BankAccountHandler) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *BankAccountHandler) Get(c *gin.Context, login string) {
	//TODO implement me
	panic("implement me")
}

func (h *BankAccountHandler) Change(c *gin.Context, login string) {
	//TODO implement me
	panic("implement me")
}

//func (h *BankAccountHandler) GetTripByID(ginCtx *gin.Context, tripId openapitypes.UUID, params generated.GetTripByIDParams) {
//	tr := global.Tracer(domain.ServiceName)
//	ctxTrace, span := tr.Start(ginCtx, "bankAccount/bankAccount_api.GetTripByID")
//	defer span.End()
//
//	ctx := zapctx.WithLogger(ctxTrace, h.logger)
//
//	trip, err := h.bankAccountService.GetTripByID(ctx, params.UserId, tripId)
//	if err != nil {
//		AbortWithBadResponse(ginCtx, h.logger, MapErrorToCode(err), err)
//		return
//	}
//	resp := models.ToTripResponse(*trip)
//
//	ginCtx.JSON(http.StatusOK, resp)
//}

//func (h *BankAccountHandler) AcceptTrip(ginCtx *gin.Context, tripId openapitypes.UUID, params generated.AcceptTripParams) {
//	tr := global.Tracer(domain.ServiceName)
//	ctxTrace, span := tr.Start(ginCtx, "bankAccount/bankAccount_api.AcceptTrip")
//	defer span.End()
//
//	ctx := zapctx.WithLogger(ctxTrace, h.logger)
//
//	err := h.bankAccountService.AcceptTrip(ctx, params.UserId, tripId)
//	if err != nil {
//		AbortWithBadResponse(ginCtx, h.logger, MapErrorToCode(err), err)
//		return
//	}
//
//	ginCtx.JSON(http.StatusOK, http.NoBody)
//}
