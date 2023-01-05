package controller

import (
	"crypto/md5"
	"fmt"
	"go-server/service"

	"github.com/gin-gonic/gin"
)

type AccountController struct{}

func (AccountController) ListAccount(c *gin.Context) {
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

	data, total, err := service.Account{}.ListAccount(p.PageNum, p.PageSize)

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

func (AccountController) CreateAccount(c *gin.Context) {
	type AccountCreateQuery struct {
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	json := AccountCreateQuery{}
	c.ShouldBindJSON(&json)
	fmt.Printf("%v", &json)

	account := new(service.Account)

	account.Phone = json.Phone
	account.Password = fmt.Sprintf("%x", md5.Sum([]byte(json.Password)))

	id, err := account.CreateAccount()
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	account.Id = id

	c.JSON(200, gin.H{
		"code":    200,
		"message": "create success",
		"data":    account,
	})
}

func (AccountController) DeleteAccount(c *gin.Context) {
	type deleteJSON struct {
		Id int `form:"id" json:"id"`
	}
	json := deleteJSON{}
	c.ShouldBindJSON(&json)
	fmt.Printf("%v", &json)

	service.Account{}.DeleteAccountById(json.Id)

	c.JSON(200, gin.H{
		"code":    200,
		"message": "delete success",
	})
}

func (AccountController) UpdateAccount(c *gin.Context) {
	type AccountUpdateQuery struct {
		Id       int    `form:"id" json:"id" binding:"required"`
		Password string `json:"password"`
		Status   int    `json:"status" form:"status"`
	}
	json := AccountUpdateQuery{}
	c.ShouldBindJSON(&json)
	fmt.Printf("%v", &json)

	if json.Id <= 0 {
		c.JSON(200, gin.H{
			"code":    500,
			"message": "id is required",
		})
		return
	}

	account := new(service.Account)
	account.Id = json.Id

	if json.Password != "" {
		account.Password = fmt.Sprintf("%x", md5.Sum([]byte(json.Password)))
	}

	if json.Status >= 0 && json.Status <= 4 {
		account.Status = json.Status
	}

	err := account.UpdateAccount()
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
		"data":    account,
	})
}
