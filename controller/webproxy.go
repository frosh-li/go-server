package controller

import (
	"fmt"
	"go-server/service"

	"github.com/gin-gonic/gin"
)

type WebProxyController struct{}

func (WebProxyController) ListWebProxy(c *gin.Context) {
	var p PageQuery
	if c.ShouldBindQuery(&p) != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": "参数错误",
		})
		return
	}
	fmt.Println("pagequery", p)
	if p.PageNum <= 0 {
		p.PageNum = 1
	}

	if p.PageSize <= 0 {
		p.PageSize = 20
	}

	data, total, err := service.Webproxy{}.ListWebProxy(p.PageNum, p.PageSize)

	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "",
		"data": gin.H{
			"current":  p.PageNum,
			"pageSize": p.PageSize,
			"total":    total,
			"list":     data,
		},
	})
}

func (WebProxyController) CreateWebProxy(c *gin.Context) {
	type WebProxyCreateQuery struct {
		Ip     string `json:"ip"`
		Port   string `json:"port"`
		Region string `json:"region"`
	}
	json := WebProxyCreateQuery{}
	c.ShouldBindJSON(&json)
	fmt.Printf("%v", &json)

	webproxy := new(service.Webproxy)

	webproxy.Ip = json.Ip
	webproxy.Port = json.Port
	webproxy.Region = json.Region

	id, err := webproxy.CreateWebProxy()
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	webproxy.Id = id

	c.JSON(200, gin.H{
		"code":    200,
		"message": "create success",
		"data":    webproxy,
	})
}

func (WebProxyController) DeleteWebProxy(c *gin.Context) {
	type deleteJSON struct {
		Id int `form:"id" json:"id"`
	}
	json := deleteJSON{}
	c.ShouldBindJSON(&json)
	fmt.Printf("%v", &json)

	service.Webproxy{}.DeleteWebProxyById(json.Id)

	c.JSON(200, gin.H{
		"code":    200,
		"message": "delete success",
	})
}

func (WebProxyController) UpdateWebProxy(c *gin.Context) {
	type WebProxyUpdateQuery struct {
		Id     int    `form:"id" json:"id" binding:"required"`
		Ip     string `json:"ip" binding:"required"`
		Port   string `json:"port" binding:"required"`
		Region string `json:"region" binding:"required"`
	}
	json := WebProxyUpdateQuery{}
	c.ShouldBindJSON(&json)
	fmt.Printf("%v", &json)

	if json.Id <= 0 {
		c.JSON(200, gin.H{
			"code":    500,
			"message": "id is required",
		})
		return
	}

	if json.Ip == "" || json.Port == "" || json.Region == "" {
		c.JSON(200, gin.H{
			"code":    500,
			"message": "ip, port and region is required",
		})
		return
	}

	webproxy := new(service.Webproxy)
	webproxy.Id = json.Id
	webproxy.Ip = json.Ip
	webproxy.Port = json.Port
	webproxy.Region = json.Region

	err := webproxy.UpdateWebProxy()
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "update success",
		"data":    webproxy,
	})
}
