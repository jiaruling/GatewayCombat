package dto

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/10 17:02
*/
type APPListInput struct {
	Info     string `json:"info" form:"info" comment:"查找信息" binding:""`
	PageSize int    `json:"page_size" form:"page_size" comment:"页数" binding:"required,min=1,max=999"`
	PageNo   int    `json:"page_no" form:"page_no" comment:"页码" binding:"required,min=1,max=999"`
}

type APPSingleByIdInput struct {
	ID int64 `json:"id" form:"id" comment:"租户ID" binding:"required"`
}

type APPAddHttpInput struct {
	AppID    string `json:"app_id" form:"app_id" comment:"租户id" binding:"required"`
	Name     string `json:"name" form:"name" comment:"租户名称" binding:"required"`
	Secret   string `json:"secret" form:"secret" comment:"密钥" binding:""`
	WhiteIPS string `json:"white_ips" form:"white_ips" comment:"ip白名单，支持前缀匹配"`
	Qpd      int64  `json:"qpd" form:"qpd" comment:"日请求量限制" binding:""`
	Qps      int64  `json:"qps" form:"qps" comment:"每秒请求量限制" binding:""`
}

type APPUpdateHttpInput struct {
	APPSingleByIdInput
	APPAddHttpInput
}

// ---------------------------------------------------------------------------------------------------------------------
type APPListItemOutput struct {
	ID        int64  `json:"id" gorm:"primary_key"`
	AppID     string `json:"app_id" gorm:"column:app_id" description:"租户id	"`
	Name      string `json:"name" gorm:"column:name" description:"租户名称	"`
	Secret    string `json:"secret" gorm:"column:secret" description:"密钥"`
	WhiteIPS  string `json:"white_ips" gorm:"column:white_ips" description:"ip白名单，支持前缀匹配		"`
	Qpd       int64  `json:"qpd" gorm:"column:qpd" description:"日请求量限制"`
	Qps       int64  `json:"qps" gorm:"column:qps" description:"每秒请求量限制"`
	RealQpd   int64  `json:"real_qpd" description:"日请求量限制"`
	RealQps   int64  `json:"real_qps" description:"每秒请求量限制"`
	UpdatedAt int64  `json:"create_at" gorm:"column:create_at" description:"添加时间	"`
	CreatedAt int64  `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	DeletedAt int64  `json:"is_delete" gorm:"column:is_delete" description:"是否已删除；0：否；1：是"`
}

type APPListOutput struct {
	List     []APPListItemOutput `json:"list" form:"list" comment:"租户列表"`
	Total    int64               `json:"total" form:"total" comment:"租户总数"`
	Page     int64               `json:"page" form:"page" comment:"页码"`
	PageSize int64               `json:"page_size" form:"page_size" comment:"每一页大小"`
}

type StatisticsOutput struct {
	Today     []int64 `json:"today" form:"today" comment:"今日统计" validate:"required"`
	Yesterday []int64 `json:"yesterday" form:"yesterday" comment:"昨日统计" validate:"required"`
}
