package bank_account_api

import (
	"go.uber.org/zap"
	"time"
	bankaccount "xor-go/services/finances/internal/handler/generated/bank-account"
	"xor-go/services/finances/internal/service/adapters"
)

var _ bankaccount.ServerInterface = &BankAccountHandler{}

type BankAccountHandler struct {
	logger             *zap.Logger
	bankAccountService adapters.BankAccountService
	WaitTimeout        time.Duration
}

func NewBankAccountHandler(logger *zap.Logger, bankAccountService adapters.BankAccountService, socketTimeout time.Duration) *BankAccountHandler {
	return &BankAccountHandler{logger: logger, bankAccountService: bankAccountService, WaitTimeout: socketTimeout}
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
