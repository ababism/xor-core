package http

import (
	"fmt"
	"strings"
	"xor-go/services/finances/internal/config"
	"xor-go/services/finances/internal/handler/http/bank"
	"xor-go/services/finances/internal/handler/http/discount"
	"xor-go/services/finances/internal/handler/http/payment"
	"xor-go/services/finances/internal/handler/http/payout"
	"xor-go/services/finances/internal/handler/http/product"
	"xor-go/services/finances/internal/handler/http/purchase"
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

func ConvertToBankAccount(middlewareMainArr []MiddlewareFunc) []bank.MiddlewareFunc {
	result := make([]bank.MiddlewareFunc, len(middlewareMainArr))
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

func ConvertToPayoutRequest(middlewareMainArr []MiddlewareFunc) []payout.MiddlewareFunc {
	result := make([]payout.MiddlewareFunc, len(middlewareMainArr))
	for i, middlewareMain := range middlewareMainArr {
		result[i] = func(c *gin.Context) {
			middlewareMain(c)
		}
	}
	return result
}

func ConvertToPurchaseRequest(middlewareMainArr []MiddlewareFunc) []purchase.MiddlewareFunc {
	result := make([]purchase.MiddlewareFunc, len(middlewareMainArr))
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

	bank.RegisterHandlersWithOptions(router,
		bank.NewBankAccountHandler(handler.bankAccountService),
		bank.GinServerOptions{
			BaseURL:      baseUrl,
			Middlewares:  ConvertToBankAccount(middlewares),
			ErrorHandler: HandleError,
		})

	discount.RegisterHandlersWithOptions(router,
		discount.NewDiscountHandler(handler.discountService),
		discount.GinServerOptions{
			BaseURL:      baseUrl,
			Middlewares:  ConvertToDiscount(middlewares),
			ErrorHandler: HandleError,
		})

	product.RegisterHandlersWithOptions(router,
		product.NewProductHandler(handler.productService),
		product.GinServerOptions{
			BaseURL:      baseUrl,
			Middlewares:  ConvertToProduct(middlewares),
			ErrorHandler: HandleError,
		})

	payment.RegisterHandlersWithOptions(router,
		payment.NewPaymentHandler(handler.paymentService),
		payment.GinServerOptions{
			BaseURL:      baseUrl,
			Middlewares:  ConvertToPayment(middlewares),
			ErrorHandler: HandleError,
		})

	purchase.RegisterHandlersWithOptions(router,
		purchase.NewPurchaseRequestHandler(handler.purchaseRequestService),
		purchase.GinServerOptions{
			BaseURL:      baseUrl,
			Middlewares:  ConvertToPurchaseRequest(middlewares),
			ErrorHandler: HandleError,
		})

	payout.RegisterHandlersWithOptions(router,
		payout.NewPayoutRequestHandler(handler.payoutRequestService),
		payout.GinServerOptions{
			BaseURL:      baseUrl,
			Middlewares:  ConvertToPayoutRequest(middlewares),
			ErrorHandler: HandleError,
		})
}

func getVersion() string {
	return fmt.Sprintf("v%s", strings.Split(version, ".")[0])
}
