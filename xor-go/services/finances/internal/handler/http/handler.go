package http

import (
	"fmt"
	"strings"
	"xor-go/services/finances/internal/config"
	bank_account "xor-go/services/finances/internal/handler/generated/bank-account"
	"xor-go/services/finances/internal/handler/generated/discount"
	"xor-go/services/finances/internal/handler/generated/payment"
	payout_request "xor-go/services/finances/internal/handler/generated/payout-request"
	"xor-go/services/finances/internal/handler/generated/product"
	purchase_request "xor-go/services/finances/internal/handler/generated/purchase-request"
	"xor-go/services/finances/internal/handler/http/bank_account_api"
	"xor-go/services/finances/internal/handler/http/discount_api"
	"xor-go/services/finances/internal/handler/http/payment_api"
	"xor-go/services/finances/internal/handler/http/payout_request_api"
	"xor-go/services/finances/internal/handler/http/product_api"
	"xor-go/services/finances/internal/handler/http/purchase_request_api"
	"xor-go/services/finances/internal/service/adapters"

	"github.com/gin-gonic/gin"
)

const (
	httpPrefix = "api"
	version    = "1"
)

type MiddlewareFunc func(c *gin.Context)

type Handler struct {
	cfg                    *config.Config
	bankAccountService     adapters.BankAccountService
	discountService        adapters.DiscountService
	paymentService         adapters.PaymentService
	productService         adapters.ProductService
	purchaseRequestService adapters.PurchaseRequestService
	payoutRequestService   adapters.PayoutRequestService
}

func NewHandler(cfg *config.Config,
	bankAccountService adapters.BankAccountService,
	discountService adapters.DiscountService,
	paymentService adapters.PaymentService,
	productService adapters.ProductService,
	purchaseRequestService adapters.PurchaseRequestService,
	payoutRequestService adapters.PayoutRequestService,
) Handler {
	return Handler{
		cfg:                    cfg,
		bankAccountService:     bankAccountService,
		discountService:        discountService,
		paymentService:         paymentService,
		productService:         productService,
		purchaseRequestService: purchaseRequestService,
		payoutRequestService:   payoutRequestService,
	}
}

// HandleError is a sample error handler function
func HandleError(c *gin.Context, err error, statusCode int) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}

func ConvertToBankAccount(middlewareMainArr []MiddlewareFunc) []bank_account.MiddlewareFunc {
	result := make([]bank_account.MiddlewareFunc, len(middlewareMainArr))
	for i, middlewareMain := range middlewareMainArr {
		result[i] = func(c *gin.Context) {
			middlewareMain(c)
		}
	}
	return result
}

func ConvertToDiscount(middlewareMainArr []MiddlewareFunc) []discount.MiddlewareFunc {
	result := make([]discount.MiddlewareFunc, len(middlewareMainArr))
	for i, middlewareMain := range middlewareMainArr {
		result[i] = func(c *gin.Context) {
			middlewareMain(c)
		}
	}
	return result
}

func ConvertToProduct(middlewareMainArr []MiddlewareFunc) []product.MiddlewareFunc {
	result := make([]product.MiddlewareFunc, len(middlewareMainArr))
	for i, middlewareMain := range middlewareMainArr {
		result[i] = func(c *gin.Context) {
			middlewareMain(c)
		}
	}
	return result
}

func ConvertToPayment(middlewareMainArr []MiddlewareFunc) []payment.MiddlewareFunc {
	result := make([]payment.MiddlewareFunc, len(middlewareMainArr))
	for i, middlewareMain := range middlewareMainArr {
		result[i] = func(c *gin.Context) {
			middlewareMain(c)
		}
	}
	return result
}

func ConvertToPayoutRequest(middlewareMainArr []MiddlewareFunc) []payout_request.MiddlewareFunc {
	result := make([]payout_request.MiddlewareFunc, len(middlewareMainArr))
	for i, middlewareMain := range middlewareMainArr {
		result[i] = func(c *gin.Context) {
			middlewareMain(c)
		}
	}
	return result
}

func ConvertToPurchaseRequest(middlewareMainArr []MiddlewareFunc) []purchase_request.MiddlewareFunc {
	result := make([]purchase_request.MiddlewareFunc, len(middlewareMainArr))
	for i, middlewareMain := range middlewareMainArr {
		result[i] = func(c *gin.Context) {
			middlewareMain(c)
		}
	}
	return result
}

func InitHandler(
	handler Handler,
	router gin.IRouter,
	middlewares []MiddlewareFunc,
) {
	baseUrl := fmt.Sprintf("%s/%s", httpPrefix, getVersion())

	bank_account.RegisterHandlersWithOptions(router,
		bank_account_api.NewBankAccountHandler(handler.bankAccountService),
		bank_account.GinServerOptions{
			BaseURL:      baseUrl,
			Middlewares:  ConvertToBankAccount(middlewares),
			ErrorHandler: HandleError,
		})

	discount.RegisterHandlersWithOptions(router,
		discount_api.NewDiscountHandler(handler.discountService),
		discount.GinServerOptions{
			BaseURL:      baseUrl,
			Middlewares:  ConvertToDiscount(middlewares),
			ErrorHandler: HandleError,
		})

	product.RegisterHandlersWithOptions(router,
		product_api.NewProductHandler(handler.productService),
		product.GinServerOptions{
			BaseURL:      baseUrl,
			Middlewares:  ConvertToProduct(middlewares),
			ErrorHandler: HandleError,
		})

	payment.RegisterHandlersWithOptions(router,
		payment_api.NewPaymentHandler(handler.paymentService),
		payment.GinServerOptions{
			BaseURL:      baseUrl,
			Middlewares:  ConvertToPayment(middlewares),
			ErrorHandler: HandleError,
		})

	purchase_request.RegisterHandlersWithOptions(router,
		purchase_request_api.NewPurchaseRequestHandler(handler.purchaseRequestService),
		purchase_request.GinServerOptions{
			BaseURL:      baseUrl,
			Middlewares:  ConvertToPurchaseRequest(middlewares),
			ErrorHandler: HandleError,
		})

	payout_request.RegisterHandlersWithOptions(router,
		payout_request_api.NewPayoutRequestHandler(handler.payoutRequestService),
		payout_request.GinServerOptions{
			BaseURL:      baseUrl,
			Middlewares:  ConvertToPayoutRequest(middlewares),
			ErrorHandler: HandleError,
		})
}

func getVersion() string {
	return fmt.Sprintf("v%s", strings.Split(version, ".")[0])
}
