package service

import (
	"fmt"
	"go-server/dao"
)

type Machine struct {
	Id      int    `json:"id"`
	Mip     string `json:"mip"`
	Vpnid   int    `json:"vpnid"`
	Created string `json:"created"`
	Windows int    `json:"windows"`
}

// 新建机器，同时创建对应窗口
func (machine *Machine) CreateMachine() (id int, err error) {
	err = dao.DB.Create(&machine).Error
	Window{}.CreateWindows(machine.Windows, machine.Id)
	return machine.Id, err
}

// 删除
func (Machine) DeleteMachineById(id int) (err error) {

	err = dao.DB.Delete(&Machine{}, id).Error
	Window{}.DeleteWindowsByMachineId(id)
	return
}

// 分页获取所有的端口列表
func (Machine) ListMachine(pageNum int, pageSize int) (vpn []Machine, total int64, err error) {
	fmt.Println("params", pageNum, pageSize, (pageNum-1)*pageSize)
	queryError := dao.DB.Debug().Model(&Machine{}).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&vpn).Error
	if queryError != nil {
		err = queryError
	}

	countError := dao.DB.Debug().Model(&Machine{}).Count(&total).Error
	if countError != nil {
		err = countError
	}
	return
}

// 修改
func (machine *Machine) UpdateMachine() error {
	var iMachine Machine = Machine{Id: machine.Id}

	iMachine.Id = machine.Id
	iMachine.Mip = machine.Mip
	iMachine.Vpnid = machine.Vpnid
	err := dao.DB.Debug().Save(&iMachine).Error
	return err
}
