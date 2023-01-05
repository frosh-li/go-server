package route

import (
	"go-server/controller"

	"github.com/gin-gonic/gin"
)

func AccountRegister(api *gin.RouterGroup) {

	accountRoute := api.Group("/account")
	{
		accountRoute.POST("/create", controller.AccountController{}.CreateAccount)
		accountRoute.POST("/update", controller.AccountController{}.UpdateAccount)
		accountRoute.POST("/delete", controller.AccountController{}.DeleteAccount)
		accountRoute.GET("/list", controller.AccountController{}.ListAccount)
	}
}
