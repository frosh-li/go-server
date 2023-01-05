package service

import (
	"fmt"
	"go-server/dao"
)

type Account struct {
	Id        int    `json:"id"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Status    int    `json:"status" gorm:"default:1"`
	LastLogin string `json:"last_login"`
}

/**
 *
 * status
 * 0 锁定
 * 1 正常
 * 2 试用中
 * 3 删除
 */

// 新建
func (account *Account) CreateAccount() (id int, err error) {
	err = dao.DB.Create(&account).Error
	return account.Id, err
}

// 删除用户 需要删除
func (Account) DeleteAccountById(id int) (err error) {
	var iAccount Account = Account{Id: id}
	iAccount.Status = 4
	err = dao.DB.Debug().Save(&iAccount).Error
	return
}

// 分页获取所有的账号列表
func (Account) ListAccount(pageNum int, pageSize int) (account []Account, total int64, err error) {
	fmt.Println("params", pageNum, pageSize, (pageNum-1)*pageSize)
	queryError := dao.DB.Debug().Model(&Account{}).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&account).Error
	if queryError != nil {
		err = queryError
	}

	countError := dao.DB.Debug().Model(&Account{}).Count(&total).Error
	if countError != nil {
		err = countError
	}
	return
}

// 修改用户密码或者修改用户状态
func (account *Account) UpdateAccount() error {
	var iAccount Account = Account{Id: account.Id}

	iAccount.Password = account.Password
	iAccount.Status = account.Status
	err := dao.DB.Debug().Save(&iAccount).Error
	return err
}
