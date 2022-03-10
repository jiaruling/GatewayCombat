package model

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/10 16:59
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

func (t *App) TableName() string {
	return "gateway_app"
}
