package model

import (
	"GatewayCombat/global"
	"GatewayCombat/service/grf"
)

/*
   功能说明: 网关基本信息模型
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/10 14:25
*/

type ServiceInfo struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	LoadType    int    `json:"load_type" db:"load_type" gorm:"column:load_type"`
	ServiceName string `json:"service_name" db:"service_name" gorm:"column:service_name"`
	ServiceDesc string `json:"service_desc" db:"service_desc" gorm:"column:service_desc"`
	UpdatedAt   int64  `json:"created_at" db:"created_at" gorm:"column:created_at"`
	CreatedAt   int64  `json:"updated_at" db:"updated_at" gorm:"column:updated_at"`
	DeleteAt    int64  `json:"deleted_at" db:"deleted_at" gorm:"column:deleted_at"`
}

func (t *ServiceInfo) TableName() string {
	return "gateway_service_info"
}

var Service = grf.Model{
	M:            nil, // M: new(Student) 传入模型的结构体指针
	Table:        "",
	AllowMethods: []string{global.METHOD_GET, global.METHOD_DELETE},
	CreateField: grf.CreateField{
		CreatedFields:        nil,
		CreatedIgnoreFields:  []string{"deleted_at"},
		CreatedSetTimeFields: []string{"created_at", "updated_at"},
	},
	SoftDeleteField: grf.SoftDeleteField{
		DeletedFields: "deleted_at",
	},
	UpdateField: grf.UpdateField{
		UpdateFields:        nil,
		UpdateIgnoreFields:  []string{"created_at", "deleted_at"},
		UpdateSetTimeFields: []string{"updated_at"},
	},
	SelectField: grf.SelectField{
		SelectFields:       nil,
		SelectIgnoreFields: []string{"created_at", "updated_at", "deleted_at"},
	},
	SelectFieldList: grf.SelectFieldList{
		Search:  []string{"name", "age"},
		Filter:  nil,
		Sort:    []string{"id"},
		PageMax: 100,
		PageMin: 10,
	},
}
