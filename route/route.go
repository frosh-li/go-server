package route

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
)

// 将其他路由统一注册到主程序中

func SetupRouter() *gin.Engine {
	route := gin.Default()
	store, err := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("secret"))
	if err != nil {
		panic("redis session连接失败")
	}
	fmt.Println(&store)
	route.Use(sessions.Sessions("mysession", store))

	api := route.Group("/api")
	UserRegister(api)
	ProxyRegister(api)
	VpnRegister(api)
	MachineRegister(api)
	return route
}
