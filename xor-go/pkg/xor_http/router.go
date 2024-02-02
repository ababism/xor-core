package xor_http

import (
	"github.com/gin-gonic/gin"
	"xor-go/pkg/xor_http/middleware"
)

type Router struct {
	router *gin.Engine
}

func (r *Router) Router() *gin.Engine {
	return r.router
}

func NewRouter() *Router {
	return &Router{
		router: gin.Default(),
	}
}

func NewRouterWithSystemHandlers() *Router {
	r := Router{
		router: gin.Default(),
	}
	r.RegisterSystemHandlers()
	return &r
}

func (r *Router) RegisterSystemHandlers() {
	s := r.Router().Group("/system")
	middleware.Ping(s)
	middleware.Prometheus(s)
}
