package route

import (
	"go-server/controller"

	"github.com/gin-gonic/gin"
)

func UserRegister(api *gin.RouterGroup) {

	userRoute := api.Group("/user")
	{
		userRoute.POST("/login", controller.UserController{}.UserLogin)
		userRoute.POST("/logout", controller.UserController{}.UserLogout)
		userRoute.GET("/info", controller.UserController{}.UserInfo)
		userRoute.GET("/list", controller.UserController{}.UserList)
	}
}
