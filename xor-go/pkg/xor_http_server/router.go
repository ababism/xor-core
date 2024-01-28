package xor_http_server

import (
	"github.com/gin-gonic/gin"
	"xor-go/pkg/xor_http_server/middleware"
)

type Router struct {
	router *gin.Engine
}

func (r *Router) GetRouter() *gin.Engine {
	return r.router
}

func NewRouter() *Router {
	r := Router{
		router: gin.Default(),
	}
	r.RegisterSystemHandlers()
	return &r
}

func (r *Router) RegisterSystemHandlers() {
	s := r.GetRouter().Group("/system")
	middleware.Ping(s)
	middleware.Prometheus(s)
}
