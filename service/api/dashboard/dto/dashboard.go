package dto

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/10 17:04
*/
type PanelGroupDataOutput struct {
	ServiceNum      int64 `json:"serviceNum"`
	AppNum          int64 `json:"appNum"`
	CurrentQPS      int64 `json:"currentQps"`
	TodayRequestNum int64 `json:"todayRequestNum"`
}

type DashServiceStatItemOutput struct {
	Name     string `json:"name"`
	LoadType int    `json:"load_type"`
	Value    int64  `json:"value"`
}

type DashServiceStatOutput struct {
	Legend []string                    `json:"legend"`
	Data   []DashServiceStatItemOutput `json:"data"`
}

type ServiceStatOutput struct {
	Today     []int64 `json:"today" dto:"today" comment:"今日流量" example:"" validate:""`         //列表
	Yesterday []int64 `json:"yesterday" dto:"yesterday" comment:"昨日流量" example:"" validate:""` //列表
}
