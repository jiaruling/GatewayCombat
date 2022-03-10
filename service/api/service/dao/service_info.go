package dao

import (
	"GatewayCombat/service/api/service/model"

	"gorm.io/gorm"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/10 15:43
*/

type ServiceInfoDao struct{}

func (t *ServiceInfoDao) Find(tx *gorm.DB, search *model.ServiceInfo) (*model.ServiceInfo, error) {
	out := &model.ServiceInfo{}
	err := tx.Where(search).Find(out).Error
	return out, err
}

func (t *ServiceInfoDao) Save(tx *gorm.DB, model *model.ServiceInfo) error {
	return tx.Save(model).Error
}
