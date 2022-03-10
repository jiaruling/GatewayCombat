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

type GrpcRuleDao struct{}

func (t *GrpcRuleDao) Find(tx *gorm.DB, search *model.GrpcRule) (*model.GrpcRule, error) {
	out := &model.GrpcRule{}
	err := tx.Where(search).Find(out).Error
	return out, err
}

func (t *GrpcRuleDao) Save(tx *gorm.DB, model *model.GrpcRule) error {
	if err := tx.Save(model).Error; err != nil {
		return err
	}
	return nil
}
