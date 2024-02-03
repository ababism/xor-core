package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"xor-go/pkg/xorerror"
	xorhttp "xor-go/pkg/xorhttp/response"
	"xor-go/services/sage/internal/model"
	"xor-go/services/sage/internal/service"
)

type AccountHandler struct {
	responseWrapper *xorhttp.HttpResponseWrapper
	accountService  *service.AccountService
}

func NewAccountHandler(responseWrapper *xorhttp.HttpResponseWrapper, accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{responseWrapper: responseWrapper, accountService: accountService}
}

func (h *AccountHandler) InitAccountRoutes(g *gin.RouterGroup) {
	account := g.Group("/account")
	account.POST("/register", h.Register)
	account.PUT("/update-password", h.UpdatePassword)
}

func (h *AccountHandler) Register(ctx *gin.Context) {
	var registerAccountDto model.RegisterAccountDto
	err := ctx.BindJSON(&registerAccountDto)
	if err != nil {
		h.responseWrapper.HandleErrorWithMessage(ctx, http.StatusBadRequest, err)
		return
	}

	err = h.accountService.Create(ctx, registerAccountDto.ToRegisterAccountEntity())
	if err != nil {
		xorerror.HandleInternalErrorWithMessage(ctx, h.responseWrapper, err)
		return
	}
	h.responseWrapper.HandleSuccessWithMessage(ctx, http.StatusOK, "account is registered")
}

func (h *AccountHandler) UpdatePassword(ctx *gin.Context) {
	uuidParam := "uuid"
	passwordParam := "password"
	params := ctx.Request.URL.Query()
	if !params.Has(uuidParam) || !params.Has(passwordParam) {
		h.responseWrapper.HandleErrorWithMessage(ctx, http.StatusBadRequest, errors.New("query params are not provided"))
		return
	}
	parsedUuid, err := uuid.Parse(params.Get(uuidParam))
	if err != nil {
		h.responseWrapper.HandleErrorWithMessage(ctx, http.StatusBadRequest, errors.New("failed to parse account uuid"))
		return
	}
	password := params.Get(passwordParam)

	err = h.accountService.UpdatePassword(ctx, parsedUuid, password)
	if err != nil {
		h.responseWrapper.HandleErrorWithMessage(ctx, http.StatusInternalServerError, err)
		return
	}
	h.responseWrapper.HandleSuccessWithMessage(ctx, http.StatusOK, "account password is updated")
}
