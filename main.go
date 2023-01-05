package main

import (
	"go-server/dao"
	"go-server/route"
)

func main() {
	dao.Init()
	r := route.SetupRouter()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
