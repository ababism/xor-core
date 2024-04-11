package bank

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
	"xor-go/services/finances/internal/handler/http/dto"
	http2 "xor-go/services/finances/internal/handler/http/utils"
	"xor-go/services/finances/internal/log"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanDefaultBankAccount = "bank-account/handler."
)

var _ ServerInterface = &Handler{}

type Handler struct {
	bankAccountService adapters.BankAccountService
}

func NewBankAccountHandler(bankAccountService adapters.BankAccountService) *Handler {
	return &Handler{bankAccountService: bankAccountService}
}

func getAccountTracerSpan(ctx *gin.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNameBankAccount)
	newCtx, span := tr.Start(ctx, spanDefaultBankAccount+name)

	return tr, newCtx, span
}

func (h *Handler) GetBankAccountsLoginLogin(ctx *gin.Context, login string) {
	_, newCtx, span := getAccountTracerSpan(ctx, ".GetByLogin")
	defer span.End()

	domain, err := h.bankAccountService.GetByLogin(newCtx, login)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	response := DomainToGet(*domain)

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetBankAccountsId(ctx *gin.Context, id openapi_types.UUID) {
	_, newCtx, span := getAccountTracerSpan(ctx, ".GetBankAccountsId")
	defer span.End()

	domain, err := h.bankAccountService.GetById(newCtx, id)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	response := DomainToGet(*domain)

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetBankAccounts(ctx *gin.Context) {
	_, newCtx, span := getAccountTracerSpan(ctx, ".GetList")
	defer span.End()

	var body *BankAccountFilter
	if err := ctx.BindJSON(&body); err != nil && err != io.EOF {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	domains, err := h.bankAccountService.List(newCtx, FilterToDomain(body))
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	list := make([]BankAccountGet, len(domains))
	for i, item := range domains {
		list[i] = DomainToGet(item)
	}

	ctx.JSON(http.StatusOK, list)
}

func (h *Handler) PostBankAccounts(ctx *gin.Context) {
	_, newCtx, span := getAccountTracerSpan(ctx, ".Create")
	defer span.End()

	var body BankAccountCreate
	if err := ctx.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	domain := CreateToDomain(body)
	log.Logger.Info(fmt.Sprintf("%v", domain))
	id, err := h.bankAccountService.Create(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, dto.ModelUUID{UUID: *id})
}

func (h *Handler) PutBankAccounts(ctx *gin.Context) {
	_, newCtx, span := getAccountTracerSpan(ctx, ".Update")
	defer span.End()

	var body BankAccountUpdate
	if err := ctx.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	domain := UpdateToDomain(body)
	err := h.bankAccountService.Update(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, http.NoBody)
}

func (h *Handler) PutBankAccountsIdChangeFunds(
	c *gin.Context,
	id openapi_types.UUID,
	params PutBankAccountsIdChangeFundsParams,
) {
	_, newCtx, span := getAccountTracerSpan(c, ".Change")
	defer span.End()

	err := h.bankAccountService.ChangeFunds(newCtx, id, params.NewFunds)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}
