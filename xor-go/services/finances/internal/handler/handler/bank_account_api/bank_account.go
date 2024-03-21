package bank_account_api

import (
	"context"
	"github.com/gin-gonic/gin"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	bankaccount "xor-go/services/finances/internal/handler/generated/bank-account"
	http2 "xor-go/services/finances/internal/handler/handler"
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

func (h *BankAccountHandler) Get(c *gin.Context, login string) {
	_, newCtx, span := getAccountTracerSpan(c, ".Get")
	defer span.End()

	domain, err := h.bankAccountService.Get(newCtx, login)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	response := DomainToGet(*domain)

	c.JSON(http.StatusOK, response)
}

func (h *BankAccountHandler) GetList(c *gin.Context, params bankaccount.GetListParams) {
	_, newCtx, span := getAccountTracerSpan(c, ".GetList")
	defer span.End()

	domains, err := h.bankAccountService.List(newCtx, FilterToDomain(params.Filter))
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	list := make([]bankaccount.BankAccountGet, len(domains))
	for i, item := range domains {
		list[i] = DomainToGet(item)
	}

	c.JSON(http.StatusOK, list)
}

func (h *BankAccountHandler) Create(c *gin.Context, params bankaccount.CreateParams) {
	_, newCtx, span := getAccountTracerSpan(c, ".Create")
	defer span.End()

	domain := CreateToDomain(params.Model)
	err := h.bankAccountService.Create(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}

func (h *BankAccountHandler) Update(c *gin.Context, params bankaccount.UpdateParams) {
	_, newCtx, span := getAccountTracerSpan(c, ".Update")
	defer span.End()

	domain := UpdateToDomain(params.Model)
	err := h.bankAccountService.Update(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}

func (h *BankAccountHandler) Change(c *gin.Context, login string, params bankaccount.ChangeParams) {
	_, newCtx, span := getAccountTracerSpan(c, ".Change")
	defer span.End()

	err := h.bankAccountService.ChangeFunds(newCtx, login, params.NewFunds)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}
