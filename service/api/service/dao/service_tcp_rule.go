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

type TcpRuleDao struct{}

func (t *TcpRuleDao) Find(tx *gorm.DB, search *model.TcpRule) (*model.TcpRule, error) {
	out := &model.TcpRule{}
	err := tx.Where(search).Find(out).Error
	return out, err
}

func (t *TcpRuleDao) Save(tx *gorm.DB, model *model.TcpRule) error {
	if err := tx.Save(model).Error; err != nil {
		return err
	}
	return nil
}
