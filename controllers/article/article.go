package article

import (
	"backend-intern-homework/models"
	repository "backend-intern-homework/repository/article"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// GetArticles godoc
//
// @Summary Get all articles
// @Description Get all articles
// @Tags articles
// @Accept json
// @Produce json
// @Success 200 {array} models.Article
// @Failure 500 {object} models.ErrorResponse
// @Router /articles [get]
func GetArticles(ctx *gin.Context) {
	var articles []models.Article
	var err error

	//get data form database
	if articles, err = repository.GetArticles(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &articles)
}

// CreateArticle godoc
//
// @Summary Create article
// @Description Create a article
// @Tags articles
// @Accept json
// @Produce json
// @Param request body models.CreateArticle true "Create article request"
// @Success 201 {object} models.Article
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /articles [post]
func CreateArticle(ctx *gin.Context) {
	var input models.CreateArticle

	_ = ctx.ShouldBindJSON(&input)

	if err := validate.Struct(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article := models.Article{Title: input.Title, Content: input.Content}

	if err := repository.CreateArticle(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, &article)
}

// FindArticle godoc
//
// @Summary Find article
// @Description Find a article
// @Tags articles
// @Accept json
// @Produce json
// @Param id path uint true "Article ID"
// @Success 200 {object} models.Article
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /articles/{id} [get]
func FindArticle(ctx *gin.Context) {
	id, err := getArticleId(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}

	var article *models.Article

	//get data form database
	if article, err = repository.FindArticle(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
		return
	}

	ctx.JSON(http.StatusOK, &article)
}

// UpdateArticle godoc
//
// @Summary Update article
// @Description Update a article
// @Tags articles
// @Accept json
// @Produce json
// @Param id path uint true "Article ID"
// @Param request body models.UpdateArticle true "Update article request"
// @Success 200 {object} models.Article
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /articles/{id} [patch]
func UpdateArticle(ctx *gin.Context) {
	id, err := getArticleId(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}

	var input models.UpdateArticle
	_ = ctx.ShouldBindJSON(&input)

	if err = validate.Struct(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article := models.Article{ID: id, Title: input.Title, Content: input.Content}

	if err = repository.UpdateArticle(&article); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
		return
	}

	ctx.JSON(http.StatusOK, &article)
}

// DeleteArticle godoc
//
// @Summary Delete article
// @Description Delete a article
// @Tags articles
// @Accept json
// @Produce json
// @Param id path uint true "Article ID"
// @Success 204
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /articles/{id} [delete]
func DeleteArticle(ctx *gin.Context) {
	id, err := getArticleId(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}

	article := models.Article{ID: id}

	if err = repository.DeleteArticle(&article); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func getArticleId(ctx *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
