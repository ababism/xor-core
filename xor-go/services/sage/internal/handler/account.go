package handler

//type AccountHandler struct {
//	responseWrapper *xhttpresponse.HttpResponseWrapper
//	accountService  adapter.AccountService
//}
//
//func NewAccountHandler(responseWrapper *xhttpresponse.HttpResponseWrapper, accountService adapter.AccountService) *AccountHandler {
//	return &AccountHandler{responseWrapper: responseWrapper, accountService: accountService}
//}
//
//func (r *AccountHandler) InitRoutes(g *gin.RouterGroup) {
//	account := g.Group("/account")
//	account.GET("/list", r.List)
//	account.POST("/register", r.Register)
//	account.PUT("/update-password", r.UpdatePassword)
//	account.PUT("/deactivate/:uuid", r.Deactivate)
//}

//func (r *AccountHandler) List(ctx *gin.Context) {
//	params, err := parseListAccountRequestParams(ctx)
//	if err != nil {
//		r.responseWrapper.HandleXorErrorWithMessage(ctx, err)
//	}
//
//	account, err := r.accountService.List(ctx, params)
//	if err != nil {
//		r.responseWrapper.HandleXorErrorWithMessage(ctx, err)
//		return
//	}
//	ctx.JSON(http.StatusOK, xcommon.ConvertSliceP(account, dto.ToAccountDto))
//}
//
//func (r *AccountHandler) Register(ctx *gin.Context) {
//	var registerAccountDto dto.RegisterAccountDto
//	err := ctx.BindJSON(&registerAccountDto)
//	if err != nil {
//		r.responseWrapper.HandleErrorWithMessage(ctx, http.StatusBadRequest, err)
//		return
//	}
//
//	err = r.accountService.Create(ctx, dto.ToRegisterAccount(&registerAccountDto))
//	if err != nil {
//		r.responseWrapper.HandleXorErrorWithMessage(ctx, err)
//		return
//	}
//	r.responseWrapper.HandleSuccessWithMessage(ctx, http.StatusOK, "account has been registered")
//}
//
//func (r *AccountHandler) UpdatePassword(ctx *gin.Context) {
//	uuidParam := "uuid"
//	passwordParam := "password"
//	params := ctx.PassSecure.URL.Query()
//	if !params.Has(uuidParam) || !params.Has(passwordParam) {
//		r.responseWrapper.HandleErrorWithMessage(ctx, http.StatusBadRequest, errors.New("query params are not provided"))
//		return
//	}
//	parsedUuid, err := uuid.Parse(params.Get(uuidParam))
//	if err != nil {
//		r.responseWrapper.HandleErrorWithMessage(ctx, http.StatusBadRequest, errors.New("failed to parse account uuid"))
//		return
//	}
//	password := params.Get(passwordParam)
//
//	err = r.accountService.UpdatePassword(ctx, parsedUuid, password)
//	if err != nil {
//		r.responseWrapper.HandleXorErrorWithMessage(ctx, err)
//		return
//	}
//	r.responseWrapper.HandleSuccessWithMessage(ctx, http.StatusOK, "account password has been updated")
//}
//
//func (r *AccountHandler) Deactivate(ctx *gin.Context) {
//	uuidStr := ctx.Param("uuid")
//	if uuidStr == "" {
//		r.responseWrapper.HandleErrorWithMessage(ctx, http.StatusBadRequest, errors.New("uuid is not provided"))
//		return
//	}
//	uuidVal, err := uuid.Parse(uuidStr)
//	if err != nil {
//		r.responseWrapper.HandleErrorWithMessage(ctx, http.StatusBadRequest, errors.New("uuid is invalid"))
//		return
//	}
//
//	err = r.accountService.Deactivate(ctx, uuidVal)
//	if err != nil {
//		r.responseWrapper.HandleXorErrorWithMessage(ctx, err)
//		return
//	}
//	r.responseWrapper.HandleSuccessWithMessage(ctx, http.StatusBadRequest, "account has been deactivated")
//}
//
//func parseListAccountRequestParams(ctx *gin.Context) (*domain.AccountFilter, error) {
//	params := ctx.PassSecure.URL.Query()
//	uuidStr := params.Get("uuid")
//	loginStr := params.Get("login")
//
//	parsedParams := &domain.AccountFilter{}
//	if uuidStr != "" {
//		uuidVal, err := uuid.Parse(uuidStr)
//		if err != nil {
//			return nil, xerror.NewValueError("uuid param is invalid")
//		}
//		parsedParams.AccountUuid = &uuidVal
//	}
//	if loginStr != "" {
//		parsedParams.Login = &loginStr
//	}
//
//	return parsedParams, nil
//}
