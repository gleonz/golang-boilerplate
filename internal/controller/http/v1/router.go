package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gleonz/golang-boilerplate/docs"
	"github.com/gleonz/golang-boilerplate/pkg/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter configures the router API.
// Swagger Specification:
// @title Go Clean Template API
// @description Using a translation service as an example
// @version 1.0
// @host localhost:8080
// @BasePath /v1
func NewRouter(handler *gin.Engine, l logger.Interface) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routes
	// Here you can configure your specific routes.
	// For example:
	// h := handler.Group("/v1")
	// {
	// h.GET("/my-path", myHandlingFunction)
	// }
}
