package service

import (
	"fmt"
	"go-server/dao"
)

type Vpn struct {
	Id       int    `json:"id"`
	Ip       string `json:"ip"`
	Port     string `json:"port" gorm:"default:1194"`
	Protocol string `json:"protocol" gorm:"default:udp"`
	Conf     string `json:"conf"`
}

// 新建
func (vpn *Vpn) CreateVpn() (id int, err error) {
	err = dao.DB.Create(&vpn).Error
	return vpn.Id, err
}

// 删除
func (Vpn) DeleteVpnById(id int) (err error) {

	err = dao.DB.Delete(&Vpn{}, id).Error
	return
}

// 分页获取所有的端口列表
func (Vpn) ListVpn(pageNum int, pageSize int) (vpn []Vpn, total int64, err error) {
	fmt.Println("params", pageNum, pageSize, (pageNum-1)*pageSize)
	queryError := dao.DB.Debug().Model(&Vpn{}).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&vpn).Error
	if queryError != nil {
		err = queryError
	}

	countError := dao.DB.Debug().Model(&Vpn{}).Count(&total).Error
	if countError != nil {
		err = countError
	}
	return
}

// 修改
func (vpn *Vpn) UpdateVpn() error {
	var iVpn Vpn = Vpn{Id: vpn.Id}

	iVpn.Ip = vpn.Ip
	iVpn.Port = vpn.Port
	iVpn.Protocol = vpn.Protocol
	iVpn.Conf = vpn.Conf
	err := dao.DB.Debug().Save(&iVpn).Error
	return err
}
