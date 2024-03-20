package http

import (
	"fmt"
	"strings"
	"time"
	"xor-go/services/finances/internal/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	httpPrefix = "api"
	version    = "1"
)

type Handler struct {
	cfg                 *config.Config
	driverHandler       *driverAPI.DriverHandler
	userServiceProvider adapters.DriverService
}

// HandleError is a sample error handler function
func HandleError(c *gin.Context, err error, statusCode int) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}

func InitHandler(
	router gin.IRouter,
	logger *zap.Logger,
	middlewares []generated.MiddlewareFunc,
	driverService adapters.DriverService,
	socketTimeout time.Duration,
) {
	driverHandler := driverAPI.NewDriverHandler(logger, driverService, socketTimeout)

	ginOpts := generated.GinServerOptions{
		BaseURL:      fmt.Sprintf("%s/%s", httpPrefix, getVersion()),
		Middlewares:  middlewares,
		ErrorHandler: HandleError,
	}
	generated.RegisterHandlersWithOptions(router, driverHandler, ginOpts)
}

func getVersion() string {
	return fmt.Sprintf("v%s", strings.Split(version, ".")[0])
}
