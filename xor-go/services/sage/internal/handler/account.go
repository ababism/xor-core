package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	xorhttp "xor-go/pkg/xorhttp/response"
	"xor-go/services/sage/internal/handler/dto"
	"xor-go/services/sage/internal/service"
)

type AccountHandler struct {
	responseWrapper *xorhttp.HttpResponseWrapper
	accountService  service.AccountService
}

func NewAccountHandler(responseWrapper *xorhttp.HttpResponseWrapper, accountService service.AccountService) *AccountHandler {
	return &AccountHandler{responseWrapper: responseWrapper, accountService: accountService}
}

func (r *AccountHandler) InitAccountRoutes(g *gin.RouterGroup) {
	account := g.Group("/account")
	account.POST("/register", r.Register)
	account.PUT("/update-password", r.UpdatePassword)
}

func (r *AccountHandler) Register(ctx *gin.Context) {
	var registerAccountDto dto.RegisterAccountDto
	err := ctx.BindJSON(&registerAccountDto)
	if err != nil {
		r.responseWrapper.HandleErrorWithMessage(ctx, http.StatusBadRequest, err)
		return
	}

	err = r.accountService.Create(ctx, registerAccountDto.ToRegisterAccount())
	if err != nil {
		r.responseWrapper.HandleXorErrorWithMessage(ctx, err)
		return
	}
	r.responseWrapper.HandleSuccessWithMessage(ctx, http.StatusOK, "account is registered")
}

func (r *AccountHandler) UpdatePassword(ctx *gin.Context) {
	uuidParam := "uuid"
	passwordParam := "password"
	params := ctx.Request.URL.Query()
	if !params.Has(uuidParam) || !params.Has(passwordParam) {
		r.responseWrapper.HandleErrorWithMessage(ctx, http.StatusBadRequest, errors.New("query params are not provided"))
		return
	}
	parsedUuid, err := uuid.Parse(params.Get(uuidParam))
	if err != nil {
		r.responseWrapper.HandleErrorWithMessage(ctx, http.StatusBadRequest, errors.New("failed to parse account uuid"))
		return
	}
	password := params.Get(passwordParam)

	err = r.accountService.UpdatePassword(ctx, parsedUuid, password)
	if err != nil {
		r.responseWrapper.HandleXorErrorWithMessage(ctx, err)
		return
	}
	r.responseWrapper.HandleSuccessWithMessage(ctx, http.StatusOK, "account password is updated")
}
