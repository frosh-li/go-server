package service

import (
	"fmt"
	"go-server/dao"
)

type Window struct {
	Id      int    `json:"id"`
	Mid     int    `json:"mid"`
	Port    string `json:"port"`
	Proxyid string `json:"proxyid"`
}

// 批量创建窗口
func (Window) CreateWindows(total int, mid int) (err error) {
	var wins = []Window{}
	for i := 0; i < total; i++ {
		wins = append(wins, Window{
			Id:   mid + i,
			Mid:  mid,
			Port: fmt.Sprintf("%d", i+1),
		})
	}
	err = dao.DB.Create(&wins).Error
	return err
}

// 删除
func (Window) DeleteWindowsByMachineId(mid int) (err error) {
	err = dao.DB.Where("mid=?", mid).Delete(&Window{}).Error
	return
}

// 根据机器id获取所有的窗口列表
func (Window) ListWindowsByMid(mid int, pageNum int, pageSize int) (wins []Window, total int64, err error) {
	fmt.Println("params", pageNum, pageSize, (pageNum-1)*pageSize)
	queryError := dao.DB.Debug().Model(&Window{}).Where("mid=?", mid).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&wins).Error
	if queryError != nil {
		err = queryError
	}

	countError := dao.DB.Debug().Model(&Window{}).Where("mid=?", mid).Count(&total).Error
	if countError != nil {
		err = countError
	}
	return
}

// 设置或者说修改代理服务器地址
func (win *Window) UpdateProxy() error {
	var iWin Window = Window{Id: win.Id}

	iWin.Proxyid = win.Proxyid
	err := dao.DB.Debug().Save(&iWin).Error
	return err
}
