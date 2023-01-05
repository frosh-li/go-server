package service

import (
	"go-server/dao"
)

type SysUserService struct{}

type SysUser struct {
	Id        int    `json:"id"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	LastLogin string `json:"last_login"`
}

// 查询用户是否存在
func (SysUserService) QuerySysUser(entity *SysUser) (sysUser *SysUser) {
	dao.DB.Where(entity).Limit(1).Find(&sysUser)
	return
}

// 创建用户
func (SysUserService) CreateSysUser(sysUser *SysUser) (err error) {
	err = dao.DB.Create(&sysUser).Error
	return err
}

// 根据ID删除用户
func (SysUserService) DeleteSysUser(id int) (err error) {
	err = dao.DB.Delete(&SysUser{}, id).Error
	return
}

func (SysUserService) ListSysUser() (sysUsers []SysUser) {
	err := dao.DB.Debug().Find(&sysUsers).Error
	if err != nil {
		panic(err)
	}
	return
}
