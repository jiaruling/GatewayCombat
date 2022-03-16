package main

import (
	"GatewayCombat/service/grf"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 15:46
*/

type Student struct {
	Id        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty" db:"name" binding:""`
	Age       int64  `json:"age,omitempty" db:"age" binding:"required"`
	Class     string `json:"class,omitempty" db:"class" binding:"required"`
	CreatedAt int64  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt int64  `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt int64  `json:"deleted_at,omitempty" db:"deleted_at"`
}

var stu = grf.Model{
	M:            nil, // M: new(Student) 传入模型的结构体指针
	Table:        "student",
	AllowMethods: []string{"POST", "DELETE"},
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
		PageMin: 17,
	},
}
