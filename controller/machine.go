package controller

import (
	"fmt"
	"go-server/service"

	"github.com/gin-gonic/gin"
)

type MachineController struct{}

func (MachineController) ListMachine(c *gin.Context) {
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

	data, total, err := service.Machine{}.ListMachine(p.PageNum, p.PageSize)

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

func (MachineController) CreateMachine(c *gin.Context) {
	type MachineCreateQuery struct {
		Ip      string `json:"ip"`
		Mip     string `json:"mip"`
		Vpnid   int    `json:"vpnid" form:"vpnid"`
		Windows int    `json:"windows" form:"windows"`
	}
	json := MachineCreateQuery{}
	c.ShouldBindJSON(&json)
	fmt.Printf("%v", &json)

	machine := new(service.Machine)

	machine.Mip = json.Mip
	machine.Vpnid = json.Vpnid
	machine.Windows = json.Windows

	id, err := machine.CreateMachine()
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	machine.Id = id

	c.JSON(200, gin.H{
		"code":    200,
		"message": "create success",
		"data":    machine,
	})
}

func (MachineController) DeleteMachine(c *gin.Context) {
	type deleteJSON struct {
		Id int `form:"id" json:"id"`
	}
	json := deleteJSON{}
	c.ShouldBindJSON(&json)
	fmt.Printf("%v", &json)

	service.Machine{}.DeleteMachineById(json.Id)

	c.JSON(200, gin.H{
		"code":    200,
		"message": "delete success",
	})
}

func (MachineController) UpdateMachine(c *gin.Context) {
	type MachineUpdateQuery struct {
		Id      int    `form:"id" json:"id" binding:"required"`
		Mip     string `json:"mip" binding:"required"`
		Vpnid   int    `json:"vpnid" form:"vpnid"`
		Windows int    `json:"windows" form:"windows"`
	}
	json := MachineUpdateQuery{}
	c.ShouldBindJSON(&json)
	fmt.Printf("%v", &json)

	if json.Id <= 0 {
		c.JSON(200, gin.H{
			"code":    500,
			"message": "id is required",
		})
		return
	}

	if json.Mip == "" {
		c.JSON(200, gin.H{
			"code":    500,
			"message": "mip is required",
		})
		return
	}

	machine := new(service.Machine)
	machine.Id = json.Id
	machine.Mip = json.Mip
	machine.Vpnid = json.Vpnid
	machine.Windows = json.Windows

	err := machine.UpdateMachine()
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
		"data":    machine,
	})
}
