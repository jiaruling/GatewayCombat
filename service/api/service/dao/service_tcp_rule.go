package dao

import (
	"gorm.io/gorm"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/10 15:43
*/

type TcpRule struct {
	ID        int64 `json:"id" gorm:"primary_key"`
	ServiceID int64 `json:"service_id" gorm:"column:service_id" description:"服务id	"`
	Port      int   `json:"port" gorm:"column:port" description:"端口	"`
}

func (tr *TcpRule) TableName() string {
	return "gateway_service_tcp_rule"
}

// ---------------------------------------------------------------------------------------------------------------------
func (tr *TcpRule) Find(tx *gorm.DB, search *TcpRule) (*TcpRule, error) {
	out := &TcpRule{}
	err := tx.Where(search).Find(out).Error
	return out, err
}

func (tr *TcpRule) Save(tx *gorm.DB) error {
	if err := tx.Save(tr).Error; err != nil {
		return err
	}
	return nil
}
