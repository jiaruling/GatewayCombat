package dao

import (
	"GatewayCombat/global"
	ddto "GatewayCombat/service/api/dashboard/dto"
	"GatewayCombat/service/api/service/dto"
	"GatewayCombat/service/grf"

	"gorm.io/gorm"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/10 15:43
*/

type ServiceInfo struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	LoadType    int    `json:"load_type" db:"load_type" gorm:"column:load_type"`
	ServiceName string `json:"service_name" db:"service_name" gorm:"column:service_name"`
	ServiceDesc string `json:"service_desc" db:"service_desc" gorm:"column:service_desc"`
	UpdatedAt   int64  `json:"created_at" db:"created_at" gorm:"column:created_at"`
	CreatedAt   int64  `json:"updated_at" db:"updated_at" gorm:"column:updated_at"`
	DeletedAt   int64  `json:"deleted_at" db:"deleted_at" gorm:"column:deleted_at"`
}

func (si *ServiceInfo) TableName() string {
	return "gateway_service_info"
}

// ---------------------------------------------------------------------------------------------------------------------
var Service = grf.Model{
	M:            nil, // M: new(Student) 传入模型的结构体指针
	Table:        "gateway_service_info",
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

// ---------------------------------------------------------------------------------------------------------------------
func (si *ServiceInfo) Find(tx *gorm.DB, search *ServiceInfo) (*ServiceInfo, error) {
	out := &ServiceInfo{}
	err := tx.Where(search).Find(out).Error
	return out, err
}

func (si *ServiceInfo) Save(tx *gorm.DB) error {
	return tx.Save(si).Error
}

// ---------------------------------------------------------------------------------------------------------------------
func (si *ServiceInfo) PageList(tx *gorm.DB, param *dto.ServiceListInput) ([]ServiceInfo, int64, error) {
	total := int64(0)
	list := make([]ServiceInfo, param.PageSize)
	offset := (param.PageNo - 1) * param.PageSize

	query := tx.Table(Service.Table).Where("deleted_at is null")
	if param.Info != "" {
		query = query.Where("(service_name like '%?%' or service_desc like '%?%')", param.Info, param.Info)
	}
	if err := query.Limit(param.PageSize).Offset(offset).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	query.Count(&total)
	return list, total, nil
}

func (si *ServiceInfo) ServiceDetail(tx *gorm.DB, search *ServiceInfo) (*ServiceDetail, error) {
	if search.ServiceName == "" {
		info, err := si.Find(tx, search)
		if err != nil {
			return nil, err
		}
		search = info
	}
	httpRule := &HttpRule{ServiceID: search.ID}
	httpRule, err := httpRule.Find(tx, httpRule)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	tcpRule := &TcpRule{ServiceID: search.ID}
	tcpRuleDao := TcpRule{}
	tcpRule, err = tcpRuleDao.Find(tx, tcpRule)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	grpcRule := &GrpcRule{ServiceID: search.ID}
	grpcRule, err = grpcRule.Find(tx, grpcRule)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	accessControl := &AccessControl{ServiceID: search.ID}
	accessControl, err = accessControl.Find(tx, accessControl)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	loadBalance := &LoadBalance{ServiceID: search.ID}
	loadBalanceDao := LoadBalance{}
	loadBalance, err = loadBalanceDao.Find(tx, loadBalance)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	detail := &ServiceDetail{
		Info:          search,
		HTTPRule:      httpRule,
		TCPRule:       tcpRule,
		GRPCRule:      grpcRule,
		LoadBalance:   loadBalance,
		AccessControl: accessControl,
	}
	return detail, nil
}

func (si *ServiceInfo) GroupByLoadType(tx *gorm.DB) ([]ddto.DashServiceStatItemOutput, error) {
	var list []ddto.DashServiceStatItemOutput
	if err := tx.Table(si.TableName()).Where("is_delete=0").Select("load_type, count(*) as value").Group("load_type").Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
