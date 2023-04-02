package page

import (
	"backend-intern-homework/models"
	articleRepository "backend-intern-homework/repository/article"
	pageRepository "backend-intern-homework/repository/page"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// GetPages godoc
//
// @Summary Get all pages
// @Description Get all pages
// @Tags pages
// @Accept json
// @Produce json
// @Success 200 {array} models.Page
// @Failure 500 {object} models.ErrorResponse
// @Router /pages [get]
func GetPages(ctx *gin.Context) {
	var pages []models.Page
	var err error

	//get data form database
	if pages, err = pageRepository.GetPages(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &pages)
}

// CreatePage godoc
//
// @Summary Create page
// @Description Create a page
// @Tags pages
// @Accept json
// @Produce json
// @Param request body models.CreatePage true "Create page request"
// @Success 201 {object} models.Page
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /pages [post]
func CreatePage(ctx *gin.Context) {
	var input models.CreatePage

	_ = ctx.ShouldBindJSON(&input)

	if err := validate.Struct(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nextPageKey, err := uuid.Parse(input.NextPageKey)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "next page key is wrong"})
		return
	}

	var articles []models.Article

	if len(input.ArticleIds) > 0 {
		articles, err = articleRepository.GetArticles(input.ArticleIds)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	page := models.Page{NextPageKey: nextPageKey, Articles: articles}

	if err = pageRepository.CreatePage(&page); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, &page)
}

// GetPage godoc
//
// @Summary Get page
// @Description Get a page
// @Tags pages
// @Accept json
// @Produce json
// @Param key path string true "Page Key"
// @Success 200 {object} models.Page
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /pages/{key} [get]
func GetPage(ctx *gin.Context) {
	key, err := uuid.Parse(ctx.Param("key"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "page key is wrong"})
		return
	}

	//get data form database
	page, err := pageRepository.GetPage(key)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
		return
	}

	ctx.JSON(http.StatusOK, &page)
}

// UpdatePage godoc
//
// @Summary Update page
// @Description Update a page
// @Tags pages
// @Accept json
// @Produce json
// @Param key path string true "Page Key"
// @Param request body models.UpdatePage true "Update page request"
// @Success 200 {object} models.Page
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /pages/{key} [patch]
func UpdatePage(ctx *gin.Context) {
	var input models.UpdatePage
	_ = ctx.ShouldBindJSON(&input)

	if err := validate.Struct(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key, err := uuid.Parse(ctx.Param("key"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "page key is wrong"})
		return
	}

	nextPageKey, err := uuid.Parse(input.NextPageKey)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "next page key is wrong"})
		return
	}

	var articles []models.Article

	if len(input.ArticleIds) > 0 {
		articles, err = articleRepository.GetArticles(input.ArticleIds)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	page := models.Page{Key: key, NextPageKey: nextPageKey, Articles: articles}

	if err := pageRepository.UpdatePage(&page); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "page not found", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &page)
}

// DeletePage godoc
//
// @Summary Delete page
// @Description Delete a page
// @Tags pages
// @Accept json
// @Produce json
// @Param key path string true "Page Key"
// @Success 204
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /pages/{key} [delete]
func DeletePage(ctx *gin.Context) {
	key, err := uuid.Parse(ctx.Param("key"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "page key is wrong"})
		return
	}

	page := models.Page{Key: key}

	if err := pageRepository.DeletePage(&page); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
