package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"

	"xor-go/services/courses/internal/config"
	"xor-go/services/courses/internal/handler/generated"
	coursesAPI "xor-go/services/courses/internal/handler/http/coursesapi"
	"xor-go/services/courses/internal/service/adapters"
)

const (
	httpPrefix = "api"
	version    = "1"
)

type Handler struct {
	logger              *zap.Logger
	cfg                 *config.Config
	coursesHandler      *coursesAPI.CoursesHandler
	userServiceProvider adapters.CourseService
}

// HandleError is a sample error handler function
func HandleError(c *gin.Context, err error, statusCode int) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}

func InitHandler(
	router gin.IRouter,
	logger *zap.Logger,
	middlewares []generated.MiddlewareFunc,
	coursesService adapters.CoursesService,
) {
	coursesHandler := coursesAPI.NewCoursesHandler(logger, coursesService)

	ginOpts := generated.GinServerOptions{
		BaseURL:      fmt.Sprintf("%s/%s", httpPrefix, getVersion()),
		Middlewares:  middlewares,
		ErrorHandler: HandleError,
	}
	generated.RegisterHandlersWithOptions(router, coursesHandler, ginOpts)
}

func getVersion() string {
	return fmt.Sprintf("v%s", strings.Split(version, ".")[0])
}
