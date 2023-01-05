package controller

import (
	"crypto/md5"
	"fmt"
	"go-server/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

type LoginParams struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// user login
func (UserController) UserLogin(c *gin.Context) {
	json := LoginParams{}
	c.ShouldBindJSON(&json)
	fmt.Printf("%v", &json)

	phone := json.Phone
	password := json.Password

	fmt.Println(phone, password)

	md5Password := fmt.Sprintf("%x", md5.Sum([]byte(password)))

	fmt.Println("md5password", md5Password)

	queryUserCond := new(service.SysUser)
	queryUserCond.Phone = phone
	queryUserCond.Password = md5Password

	result := service.SysUserService{}.QuerySysUser(queryUserCond)

	if result.Phone != "" {

		fmt.Println("查询到了", result)
		// 设置session相关
		session := sessions.Default(c)
		session.Set("user_id", result.Id)
		session.Set("user_phone", result.Phone)
		session.Set("last_login", result.LastLogin)
		session.Save()

		c.JSON(200, gin.H{
			"code":    200,
			"message": "login success",
			"data":    result,
		})
	} else {
		fmt.Println("没有登陆失败")
		c.JSON(200, gin.H{
			"code":    500,
			"message": "用户名或者密码错误",
		})
	}

}

// get current user info by session
func (UserController) UserInfo(c *gin.Context) {
	session := sessions.Default(c)
	user_id := session.Get("user_id")
	fmt.Println(&user_id)

	if user_id != nil {
		user := new(service.SysUser)
		user.Id = user_id.(int)
		user.Phone = session.Get("user_phone").(string)
		user.LastLogin = session.Get("last_login").(string)
		c.JSON(200, gin.H{
			"code":    200,
			"message": "user info success",
			"data":    user,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    401,
			"message": "user not login",
		})
	}
}

// user logout
func (UserController) UserLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(200, gin.H{
		"code":    200,
		"message": "logout success",
	})
}

func (UserController) UserList(c *gin.Context) {

	users := service.SysUserService{}.ListSysUser()

	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    users,
	})
}
