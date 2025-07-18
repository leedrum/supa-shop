package routes

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/leedrum/supa-shop/api-gateway/middlewares"

	"github.com/gin-gonic/gin"
)

// ReverseProxy returns a handler forwarding requests to targetHost
func ReverseProxy(targetHost string) gin.HandlerFunc {
	return func(c *gin.Context) {
		target, err := url.Parse(targetHost)
		if err != nil {
			c.String(http.StatusInternalServerError, "Bad target URL")
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(target)
		c.Request.URL.Path = c.Param("proxyPath")

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

// SetupRoutes configures all routes and applies middleware
func SetupRoutes(r *gin.Engine) {
	r.Use(middlewares.RateLimitMiddleware())
	r.Use(middlewares.Logger())

	userGroup := r.Group("/user")
	userGroup.Use(middlewares.JWTAuthMiddleware())
	userGroup.Any("/*proxyPath", ReverseProxy("http://localhost:9000"))

	authGroup := r.Group("/auth")
	authGroup.Any("/*proxyPath", ReverseProxy("http://localhost:9001"))

	orderGroup := r.Group("/order")
	orderGroup.Any("/*proxyPath", ReverseProxy("http://localhost:9002"))

	productGroup := r.Group("/product")
	productGroup.Any("/*proxyPath", ReverseProxy("http://localhost:9003"))

	// Add more routes for other services here similarly
}
