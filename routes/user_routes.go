package routes

import (
	"github.com/elizielx/inventory-api-go/controllers"
	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	*controllers.UserController
}

func NewRouteUserController(userController *controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {
	router := rg.Group("users")
	router.GET("/", uc.GetAll)
}
