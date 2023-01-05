package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() {
	username := "root"        //账号
	password := "123456"      //密码
	host := "127.0.0.1"       //数据库地址，可以是Ip或者域名
	port := 3306              //数据库端口
	Dbname := "multi_control" //数据库名
	timeout := "10s"          //连接超时，10秒

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	var err error
	var db *gorm.DB

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("数据库连接失败")
	}
	fmt.Print("数据库连接成功")
	DB = db

}
