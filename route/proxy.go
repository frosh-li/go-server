package route

import (
	"go-server/controller"

	"github.com/gin-gonic/gin"
)

func ProxyRegister(api *gin.RouterGroup) {

	proxyRoute := api.Group("/webproxy")
	{
		proxyRoute.POST("/create", controller.WebProxyController{}.CreateWebProxy)
		proxyRoute.POST("/update", controller.WebProxyController{}.UpdateWebProxy)
		proxyRoute.POST("/delete", controller.WebProxyController{}.DeleteWebProxy)
		proxyRoute.GET("/list", controller.WebProxyController{}.ListWebProxy)
	}
}
