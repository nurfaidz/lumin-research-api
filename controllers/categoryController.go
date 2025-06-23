package controllers

import (
	"net/http"
	"project-research-gin/database"
	"project-research-gin/helpers"
	"project-research-gin/models"
	"project-research-gin/structs"

	"github.com/gin-gonic/gin"
)

func FindCategory(c *gin.Context) {
	var categories []models.Category

	database.DB.Find(&categories)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "List data category",
		Data:    categories,
	})
}

func CreateCategory(c *gin.Context) {
	var req = structs.CategoryCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation error",
			Errors:  helpers.TranslateErrorMessage(err),
		})

		return
	}

	category := models.Category{
		Name: req.Name,
	}

	if err := database.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Internal server error",
			Errors:  helpers.TranslateErrorMessage(err),
		})

		return
	}

	c.JSON(http.StatusCreated, structs.CategoryResponse{
		Id:        category.Id,
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: category.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}

func FindCategoryById(c *gin.Context) {
	id := c.Param("id")

	var category models.Category

	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Record not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})

		return
	}

	c.JSON(http.StatusOK, structs.CategoryResponse{
		Id:        category.Id,
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: category.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	var category models.Category

	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Category not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})

		return
	}

	var req = structs.CategoryUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation error",
			Errors:  helpers.TranslateErrorMessage(err),
		})

		return
	}

	category.Name = req.Name

	if err := database.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Internal server error",
			Errors:  helpers.TranslateErrorMessage(err),
		})

		return
	}

	c.JSON(http.StatusOK, structs.CategoryResponse{
		Id:        category.Id,
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: category.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	var category models.Category

	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Category not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})

		return
	}

	if err := database.DB.Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Internal server error",
			Errors:  helpers.TranslateErrorMessage(err),
		})

		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "category deleted successfully",
	})
}
