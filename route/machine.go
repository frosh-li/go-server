package route

import (
	"go-server/controller"

	"github.com/gin-gonic/gin"
)

func MachineRegister(api *gin.RouterGroup) {

	machineRoute := api.Group("/machine")
	{
		machineRoute.POST("/create", controller.MachineController{}.CreateMachine)
		machineRoute.POST("/update", controller.MachineController{}.UpdateMachine)
		machineRoute.POST("/delete", controller.MachineController{}.DeleteMachine)
		machineRoute.GET("/list", controller.MachineController{}.ListMachine)
	}
}
