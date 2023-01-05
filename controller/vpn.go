package controller

import (
	"fmt"
	"go-server/service"

	"github.com/gin-gonic/gin"
)

type VpnController struct{}

func (VpnController) ListVpn(c *gin.Context) {
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

	data, total, err := service.Vpn{}.ListVpn(p.PageNum, p.PageSize)

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

func (VpnController) CreateVpn(c *gin.Context) {
	type VpnCreateQuery struct {
		Ip       string `json:"ip"`
		Port     string `json:"port"`
		Protocol string `json:"protocol"`
		Conf     string `json:"conf"`
	}
	json := VpnCreateQuery{}
	c.ShouldBindJSON(&json)
	fmt.Printf("%v", &json)

	if json.Ip == "" || json.Conf == "" {
		c.JSON(200, gin.H{
			"code":    500,
			"message": "ip and config is required",
		})
		return
	}

	vpn := new(service.Vpn)

	vpn.Ip = json.Ip
	vpn.Port = json.Port
	vpn.Protocol = json.Protocol
	vpn.Conf = json.Conf

	id, err := vpn.CreateVpn()
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	vpn.Id = id

	c.JSON(200, gin.H{
		"code":    200,
		"message": "create success",
		"data":    vpn,
	})
}

func (VpnController) DeleteVpn(c *gin.Context) {
	type deleteJSON struct {
		Id int `form:"id" json:"id"`
	}
	json := deleteJSON{}
	c.ShouldBindJSON(&json)
	fmt.Printf("%v", &json)

	service.Vpn{}.DeleteVpnById(json.Id)

	c.JSON(200, gin.H{
		"code":    200,
		"message": "delete success",
	})
}

func (VpnController) UpdateVpn(c *gin.Context) {
	type VpnUpdateQuery struct {
		Id       int    `form:"id" json:"id" binding:"required"`
		Ip       string `json:"ip" binding:"required"`
		Port     string `json:"port"`
		Protocol string `json:"protocol"`
		Conf     string `json:"conf"`
	}
	json := VpnUpdateQuery{}
	c.ShouldBindJSON(&json)
	fmt.Printf("%v", &json)

	if json.Id <= 0 {
		c.JSON(200, gin.H{
			"code":    500,
			"message": "id is required",
		})
		return
	}

	if json.Ip == "" || json.Conf == "" {
		c.JSON(200, gin.H{
			"code":    500,
			"message": "ip and config is required",
		})
		return
	}

	vpn := new(service.Vpn)
	vpn.Id = json.Id

	if json.Ip != "" {
		vpn.Ip = json.Ip
	}

	if json.Port != "" {
		vpn.Port = json.Port
	}

	if json.Protocol != "" {
		vpn.Protocol = json.Protocol
	}

	if json.Conf != "" {
		vpn.Conf = json.Conf
	}

	err := vpn.UpdateVpn()
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
		"data":    vpn,
	})
}
