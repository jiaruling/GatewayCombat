package dao

import (
	"GatewayCombat/service/api/admin/dto"
	"GatewayCombat/service/api/admin/model"
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
type AdminDao struct{}

func (ad *AdminDao) LoginCheck(tx *gorm.DB, req *dto.AdminLoginInput) (adminInfo *model.Admin, err error) {
	adminInfo, err = ad.Find(tx, &model.Admin{UserName: req.UserName, IsDelete: 0})
	if err != nil {
		return nil, errors.New("用户信息不存在")
	}
	saltPassword := utils.GenSaltPassword(adminInfo.Salt, req.Password)
	if adminInfo.Password != saltPassword {
		return nil, errors.New("密码错误，请重新输入")
	}
	return adminInfo, nil
}

func (ad *AdminDao) Find(tx *gorm.DB, search *model.Admin) (*model.Admin, error) {
	admin := &model.Admin{}
	if err := tx.Where(search).Find(admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (ad *AdminDao) Save(tx *gorm.DB, model *model.Admin) (err error) {
	return tx.Save(model).Error
}
