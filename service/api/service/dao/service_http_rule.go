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

type HttpRuleDao struct{}

func (t *HttpRuleDao) Find(tx *gorm.DB, search *model.HttpRule) (*model.HttpRule, error) {
	out := &model.HttpRule{}
	err := tx.Where(search).Find(out).Error
	return out, err
}

func (t *HttpRuleDao) Save(tx *gorm.DB, model *model.HttpRule) error {
	if err := tx.Save(model).Error; err != nil {
		return err
	}
	return nil
}
