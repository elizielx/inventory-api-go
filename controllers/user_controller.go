package controllers

import (
	"github.com/elizielx/inventory-api-go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
	database *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{database: db}
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	var users []models.User

	if err := uc.database.Find(&users).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
