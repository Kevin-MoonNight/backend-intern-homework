package routers

import (
	"net/http"

	articleController "backend-intern-homework/controllers/article"
	listController "backend-intern-homework/controllers/list"
	pageController "backend-intern-homework/controllers/page"

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
	registerPageRoutes(baseRoute)
	registerListRoutes(baseRoute)
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

func registerPageRoutes(route *gin.RouterGroup) {
	pageRoute := route.Group("/pages")
	{
		pageRoute.GET("/", pageController.GetPages)
		pageRoute.POST("/", pageController.CreatePage)
		pageRoute.GET("/:key", pageController.GetPage)
		pageRoute.PATCH("/:key", pageController.UpdatePage)
		pageRoute.DELETE("/:key", pageController.DeletePage)
	}
}

func registerListRoutes(route *gin.RouterGroup) {
	listRoute := route.Group("/lists")
	{
		listRoute.GET("/", listController.GetLists)
		listRoute.POST("/", listController.CreateList)
		listRoute.GET("/:key", listController.GetList)
		listRoute.PATCH("/:key", listController.UpdateList)
		listRoute.DELETE("/:key", listController.DeleteList)
	}
}
