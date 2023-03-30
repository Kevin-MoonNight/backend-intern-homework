package routers

import (
	"net/http"

	articleController "backend-intern-homework/controllers/article"

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

	baseRoute := route.Group("/api/v1")
	registerArticleRoutes(baseRoute)
}

func registerSwaggerRoutes(route *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func registerArticleRoutes(route *gin.RouterGroup) {
	articleRoute := route.Group("/articles")
	{
		articleRoute.GET("/", articleController.GetArticles)
		articleRoute.POST("/", articleController.CreateArticle)
		articleRoute.GET("/:id", articleController.FindArticle)
		articleRoute.PATCH("/:id", articleController.UpdateArticle)
		articleRoute.DELETE("/:id", articleController.DeleteArticle)
	}
}
