package http

import (
	"fmt"
	"strings"
	"xor-go/services/finances/internal/config"
	bank_account "xor-go/services/finances/internal/handler/generated/bank-account"
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

type Handler struct {
	cfg                    *config.Config
	bankAccountHandler     *bank_account_api.BankAccountHandler
	discountHandler        *discount_api.DiscountHandler
	paymentHandler         *payment_api.PaymentHandler
	productHandler         *product_api.ProductHandler
	purchaseRequestHandler *purchase_request_api.PurchaseRequestHandler
	payoutRequestHandler   *payout_request_api.PayoutRequestHandler
	bankAccountService     adapters.BankAccountService
	discountService        adapters.DiscountService
	paymentService         adapters.PaymentService
	productService         adapters.ProductService
	purchaseRequestService adapters.PurchaseRequestService
	payoutRequestService   adapters.PayoutRequestService
}

// HandleError is a sample error handler function
func HandleError(c *gin.Context, err error, statusCode int) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}

func RegHandlerRoutes() {

}

func InitHandler(
	router gin.IRouter,
	middlewares []generated.MiddlewareFunc,
	bankAccountService adapters.BankAccountService,
	discountService adapters.DiscountService,
	paymentService adapters.PaymentService,
	productService adapters.ProductService,
	purchaseRequestService adapters.PurchaseRequestService,
	payoutRequestService adapters.PayoutRequestService,
) {
	bankAccountHandler := bank_account_api.NewBankAccountHandler(bankAccountService)
	ginOpts := bank_account.GinServerOptions{
		BaseURL:      fmt.Sprintf("%s/%s", httpPrefix, getVersion()),
		Middlewares:  middlewares,
		ErrorHandler: HandleError,
	}
	bank_account.RegisterHandlersWithOptions(router, bankAccountHandler, ginOpts)
}

func getVersion() string {
	return fmt.Sprintf("v%s", strings.Split(version, ".")[0])
}
