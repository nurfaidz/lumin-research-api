package controllers

import (
	"net/http"
	"project-research-gin/database"
	"project-research-gin/helpers"
	"project-research-gin/models"
	"project-research-gin/structs"
	"structs"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req = structs.UserCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation error",
			Errors:  helpers.TranslateErrorMessage(err),
		})

		return
	}

	user := models.User{
		Name:     req.Name,
		Username: req.Email,
		Email:    req.Email,
		Password: helpers.HashPassword(req.Password),
	}

	if err := database.DB.Create(&user).Error; err != nil {
	}
}
