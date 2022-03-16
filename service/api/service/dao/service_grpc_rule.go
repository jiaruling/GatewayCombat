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

type GrpcRule struct {
	ID             int64  `json:"id" gorm:"primary_key"`
	ServiceID      int64  `json:"service_id" gorm:"column:service_id" description:"服务id	"`
	Port           int    `json:"port" gorm:"column:port" description:"端口	"`
	HeaderTransfor string `json:"header_transfor" gorm:"column:header_transfor" description:"header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue"`
}

func (gr *GrpcRule) TableName() string {
	return "gateway_service_grpc_rule"
}

// ---------------------------------------------------------------------------------------------------------------------
func (gr *GrpcRule) Find(tx *gorm.DB, search *GrpcRule) (*GrpcRule, error) {
	out := &GrpcRule{}
	err := tx.Where(search).Find(out).Error
	return out, err
}

func (gr *GrpcRule) Save(tx *gorm.DB) error {
	if err := tx.Save(gr).Error; err != nil {
		return err
	}
	return nil
}
