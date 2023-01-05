package service

import (
	"fmt"
	"go-server/dao"
)

type Webproxy struct {
	Id     int    `json:"id"`
	Ip     string `json:"ip"`
	Port   string `json:"port"`
	Region string `json:"region"`
}

// 新建WebProxy
func (webproxy *Webproxy) CreateWebProxy() (id int, err error) {
	err = dao.DB.Create(&webproxy).Error
	return webproxy.Id, err
}

// 删除WebProxy
func (Webproxy) DeleteWebProxyById(id int) (err error) {

	err = dao.DB.Delete(&Webproxy{}, id).Error
	return
}

// 分页获取所有的端口列表
func (Webproxy) ListWebProxy(pageNum int, pageSize int) (webProxy []Webproxy, total int64, err error) {
	fmt.Println("params", pageNum, pageSize, (pageNum-1)*pageSize)
	queryError := dao.DB.Debug().Model(&Webproxy{}).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&webProxy).Error
	if queryError != nil {
		err = queryError
	}

	countError := dao.DB.Debug().Model(&Webproxy{}).Count(&total).Error
	if countError != nil {
		err = countError
	}
	return
}

// 修改WebProxy
func (webproxy *Webproxy) UpdateWebProxy() error {
	var proxy Webproxy = Webproxy{Id: webproxy.Id}

	proxy.Ip = webproxy.Ip
	proxy.Port = webproxy.Port
	proxy.Region = webproxy.Region
	err := dao.DB.Debug().Save(&proxy).Error
	return err
}
