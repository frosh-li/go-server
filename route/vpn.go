package route

import (
	"go-server/controller"

	"github.com/gin-gonic/gin"
)

func VpnRegister(api *gin.RouterGroup) {

	vpnRoute := api.Group("/vpn")
	{
		vpnRoute.POST("/create", controller.VpnController{}.CreateVpn)
		vpnRoute.POST("/update", controller.VpnController{}.UpdateVpn)
		vpnRoute.POST("/delete", controller.VpnController{}.DeleteVpn)
		vpnRoute.GET("/list", controller.VpnController{}.ListVpn)
	}
}
