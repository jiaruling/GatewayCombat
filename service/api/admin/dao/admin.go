package dao

import (
	"GatewayCombat/service/api/admin/dto"
	"GatewayCombat/utils"
	"errors"

	"gorm.io/gorm"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/4 15:17
*/
type Admin struct {
	Id        int    `json:"id" gorm:"primary_key" description:"自增主键"`
	UserName  string `json:"user_name" gorm:"column:user_name" description:"管理员用户名"`
	Salt      string `json:"salt" gorm:"column:salt" description:"盐"`
	Password  string `json:"password" gorm:"column:password" description:"密码"`
	UpdatedAt int64  `json:"updated_at" gorm:"column:updated_at" description:"更新时间"`
	CreatedAt int64  `json:"created_at" gorm:"column:created_at" description:"创建时间"`
	DeletedAt int64  `json:"deleted_at" gorm:"column:deleted_at" description:"是否删除"`
}

func (a *Admin) TableName() string {
	return "gateway_admin"
}

// ---------------------------------------------------------------------------------------------------------------------
func (a *Admin) Find(tx *gorm.DB, search *Admin) (*Admin, error) {
	admin := &Admin{}
	if err := tx.Where(search).Find(admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (a *Admin) Save(tx *gorm.DB, model *Admin) (err error) {
	return tx.Save(model).Error
}

// ---------------------------------------------------------------------------------------------------------------------
func (a *Admin) LoginCheck(tx *gorm.DB, req *dto.AdminLoginInput) (adminInfo *Admin, err error) {
	adminInfo, err = a.Find(tx, &Admin{UserName: req.UserName})
	if err != nil {
		return nil, errors.New("用户信息不存在")
	}
	saltPassword := utils.GenSaltPassword(adminInfo.Salt, req.Password)
	if adminInfo.Password != saltPassword {
		return nil, errors.New("密码错误，请重新输入")
	}
	return adminInfo, nil
}
