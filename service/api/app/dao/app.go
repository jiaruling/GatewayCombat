package dao

import (
	"GatewayCombat/service/api/app/dto"

	"gorm.io/gorm"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/10 17:03
*/

type App struct {
	ID        int64  `json:"id" gorm:"primary_key"`
	AppID     string `json:"app_id" db:"app_id" gorm:"column:app_id"`
	Name      string `json:"name" db:"name" gorm:"column:name"`
	Secret    string `json:"secret" db:"secret" gorm:"column:secret"`
	WhiteIPS  string `json:"white_ips" db:"white_ips" gorm:"column:white_ips"`
	Qpd       int64  `json:"qpd" db:"qpd" gorm:"column:qpd"`
	Qps       int64  `json:"qps" db:"qps" gorm:"column:qps"`
	CreatedAt int64  `json:"created_at" db:"created_at" gorm:"column:created_at"`
	UpdatedAt int64  `json:"updated_at" db:"updated_at" gorm:"column:updated_at"`
	DeletedAt int64  `json:"deleted_at" db:"deleted_at" gorm:"column:deleted_at"`
}

func (a *App) TableName() string {
	return "gateway_app"
}

// ---------------------------------------------------------------------------------------------------------------------
func (a *App) Find(tx *gorm.DB, search *App) (*App, error) {
	model := &App{}
	err := tx.Where(search).Find(model).Error
	return model, err
}

func (a *App) Save(tx *gorm.DB) error {
	if err := tx.Save(a).Error; err != nil {
		return err
	}
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------
func (a *App) APPList(tx *gorm.DB, params *dto.APPListInput) ([]App, int64, error) {
	var list []App
	var count int64
	pageNo := params.PageNo
	pageSize := params.PageSize

	//limit offset,pagesize
	offset := (pageNo - 1) * pageSize
	query := tx.Table(a.TableName()).Select("*")
	query = query.Where("is_delete=?", 0)
	if params.Info != "" {
		query = query.Where(" (name like '%?%' or app_id like '%?%')", params.Info, params.Info)
	}
	err := query.Limit(pageSize).Offset(offset).Order("id desc").Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	errCount := query.Count(&count).Error
	if errCount != nil {
		return nil, 0, err
	}
	return list, count, nil
}
