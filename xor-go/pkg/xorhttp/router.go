package xorhttp

import (
	"github.com/gin-gonic/gin"
	"xor-go/pkg/xorhttp/middleware"
)

type Router struct {
	router *gin.Engine
}

func (r *Router) Router() *gin.Engine {
	return r.router
}

func NewRouter() *Router {
	r := &Router{router: gin.Default()}
	r.RegisterSystemHandlers()
	return r
}

func (r *Router) RegisterSystemHandlers() {
	s := r.Router().Group("/system")
	middleware.Ping(s)
	middleware.Prometheus(s)
}
