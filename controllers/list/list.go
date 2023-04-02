package list

import (
	"backend-intern-homework/models"
	listRepository "backend-intern-homework/repository/list"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// GetLists godoc
//
// @Summary Get all lists
// @Description Get all lists
// @Tags lists
// @Accept json
// @Produce json
// @Success 200 {array} models.List
// @Failure 500 {object} models.ErrorResponse
// @Router /lists [get]
func GetLists(ctx *gin.Context) {
	var lists []models.List
	var err error

	//get data form database
	if lists, err = listRepository.GetLists(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &lists)
}

// CreateList godoc
//
// @Summary Create list
// @Description Create a list
// @Tags lists
// @Accept json
// @Produce json
// @Param request body models.CreateList true "Create list request"
// @Success 201 {object} models.List
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /lists [post]
func CreateList(ctx *gin.Context) {
	var input models.CreateList

	_ = ctx.ShouldBindJSON(&input)

	err := validate.Struct(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var nextPageKey uuid.UUID = uuid.Nil

	if input.NextPageKey != "" {
		if nextPageKey, err = uuid.Parse(input.NextPageKey); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "next page key is wrong"})
			return
		}
	}

	list := models.List{Name: input.Name, NextPageKey: &nextPageKey}

	if err = listRepository.CreateList(&list); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, &list)
}

// GetList godoc
//
// @Summary Get list
// @Description Get a list
// @Tags lists
// @Accept json
// @Produce json
// @Param key path string true "List Key"
// @Success 200 {object} models.List
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /lists/{key} [get]
func GetList(ctx *gin.Context) {
	key, err := uuid.Parse(ctx.Param("key"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "list key is wrong"})
		return
	}

	//get data form database
	list, err := listRepository.GetList(key)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "list not found"})
		return
	}

	ctx.JSON(http.StatusOK, &list)
}

// UpdateList godoc
//
// @Summary Update list
// @Description Update a list
// @Tags lists
// @Accept json
// @Produce json
// @Param key path string true "List Key"
// @Param request body models.UpdateList true "Update list request"
// @Success 200 {object} models.List
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /lists/{key} [patch]
func UpdateList(ctx *gin.Context) {
	var input models.UpdateList
	_ = ctx.ShouldBindJSON(&input)

	if err := validate.Struct(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key, err := uuid.Parse(ctx.Param("key"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "list key is wrong"})
		return
	}

	var nextPageKey uuid.UUID = uuid.Nil

	if input.NextPageKey != "" {
		if nextPageKey, err = uuid.Parse(input.NextPageKey); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "next page key is wrong"})
			return
		}
	}

	list := models.List{Key: key, Name: input.Name, NextPageKey: &nextPageKey}

	if err := listRepository.UpdateList(&list); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "list not found"})
		return
	}

	ctx.JSON(http.StatusOK, &list)
}

// DeleteList godoc
//
// @Summary Delete list
// @Description Delete a list
// @Tags lists
// @Accept json
// @Produce json
// @Param key path string true "List Key"
// @Success 204
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /lists/{key} [delete]
func DeleteList(ctx *gin.Context) {
	key, err := uuid.Parse(ctx.Param("key"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "list key is wrong"})
		return
	}

	list := models.List{Key: key}

	if err := listRepository.DeleteList(&list); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "list not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
