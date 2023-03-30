package routers

import (
	"net/http"

	// "backend-intern-homework/controllers"

	docs "backend-intern-homework/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(route *gin.Engine) {
	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})

	registerSwaggerRoutes(route)

	// baseRoute := route.Group("/api/v1")
}

func registerSwaggerRoutes(route *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}