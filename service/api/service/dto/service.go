package dto

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/10 17:04
*/

type ServiceListInput struct {
	Info     string `json:"info" dto:"info" comment:"关键词" example:"" binding:""`                      //关键词
	PageNo   int    `json:"page_no" dto:"page_no" comment:"页数" example:"1" binding:"required"`        //页数
	PageSize int    `json:"page_size" dto:"page_size" comment:"每页条数" example:"20" binding:"required"` //每页条数
}

type ServiceSingleByIdInput struct {
	ID int64 `json:"id" dto:"id" comment:"服务ID" example:"56" binding:"required"` //服务ID
}

type ServiceAddHTTPInput struct {
	ServiceName string `json:"service_name" dto:"service_name" comment:"服务名" example:"" binding:"required,valid_service_name"` //服务名
	ServiceDesc string `json:"service_desc" dto:"service_desc" comment:"服务描述" example:"" binding:"required,max=255,min=1"`     //服务描述

	RuleType       int    `json:"rule_type" dto:"rule_type" comment:"接入类型" example:"" binding:"max=1,min=0"`                           //接入类型
	Rule           string `json:"rule" dto:"rule" comment:"接入路径：域名或者前缀" example:"" binding:"required,valid_rule"`                      //域名或者前缀
	NeedHttps      int    `json:"need_https" dto:"need_https" comment:"支持https" example:"" binding:"max=1,min=0"`                      //支持https
	NeedStripUri   int    `json:"need_strip_uri" dto:"need_strip_uri" comment:"启用strip_uri" example:"" binding:"max=1,min=0"`          //启用strip_uri
	NeedWebsocket  int    `json:"need_websocket" dto:"need_websocket" comment:"是否支持websocket" example:"" binding:"max=1,min=0"`        //是否支持websocket
	UrlRewrite     string `json:"url_rewrite" dto:"url_rewrite" comment:"url重写功能" example:"" binding:"valid_url_rewrite"`              //url重写功能
	HeaderTransfor string `json:"header_transfor" dto:"header_transfor" comment:"header转换" example:"" binding:"valid_header_transfor"` //header转换

	OpenAuth          int    `json:"open_auth" dto:"open_auth" comment:"是否开启权限" example:"" binding:"max=1,min=0"`                  //关键词
	BlackList         string `json:"black_list" dto:"black_list" comment:"黑名单ip" example:"" binding:""`                            //黑名单ip
	WhiteList         string `json:"white_list" dto:"white_list" comment:"白名单ip" example:"" binding:""`                            //白名单ip
	ClientipFlowLimit int    `json:"clientip_flow_limit" dto:"clientip_flow_limit" comment:"客户端ip限流	" example:"" binding:"min=0"` //客户端ip限流
	ServiceFlowLimit  int    `json:"service_flow_limit" dto:"service_flow_limit" comment:"服务端限流" example:"" binding:"min=0"`       //服务端限流

	RoundType              int    `json:"round_type" dto:"round_type" comment:"轮询方式" example:"" binding:"max=3,min=0"`                                //轮询方式
	IpList                 string `json:"ip_list" dto:"ip_list" comment:"ip列表" example:"" binding:"required,valid_ipportlist"`                        //ip列表
	WeightList             string `json:"weight_list" dto:"weight_list" comment:"权重列表" example:"" binding:"required,valid_weightlist"`               //权重列表
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" dto:"upstream_connect_timeout" comment:"建立连接超时, 单位s" example:"" binding:"min=0"`   //建立连接超时, 单位s
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" dto:"upstream_header_timeout" comment:"获取header超时, 单位s" example:"" binding:"min=0"` //获取header超时, 单位s
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" dto:"upstream_idle_timeout" comment:"链接最大空闲时间, 单位s" example:"" binding:"min=0"`       //链接最大空闲时间, 单位s
	UpstreamMaxIdle        int    `json:"upstream_max_idle" dto:"upstream_max_idle" comment:"最大空闲链接数" example:"" binding:"min=0"`                     //最大空闲链接数
}

type ServiceUpdateHTTPInput struct {
	ServiceSingleByIdInput
	ServiceAddHTTPInput
}

type ServiceAddTcpInput struct {
	ServiceName       string `json:"service_name" dto:"service_name" comment:"服务名称" validate:"required,valid_service_name"`
	ServiceDesc       string `json:"service_desc" dto:"service_desc" comment:"服务描述" validate:"required"`
	Port              int    `json:"port" dto:"port" comment:"端口，需要设置8001-8999范围内" validate:"required,min=8001,max=8999"`
	HeaderTransfor    string `json:"header_transfor" dto:"header_transfor" comment:"header头转换" validate:""`
	OpenAuth          int    `json:"open_auth" dto:"open_auth" comment:"是否开启权限验证" validate:""`
	BlackList         string `json:"black_list" dto:"black_list" comment:"黑名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_iplist"`
	WhiteList         string `json:"white_list" dto:"white_list" comment:"白名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_iplist"`
	WhiteHostName     string `json:"white_host_name" dto:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"valid_iplist"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" dto:"clientip_flow_limit" comment:"客户端IP限流" validate:""`
	ServiceFlowLimit  int    `json:"service_flow_limit" dto:"service_flow_limit" comment:"服务端限流" validate:""`
	RoundType         int    `json:"round_type" dto:"round_type" comment:"轮询策略" validate:""`
	IpList            string `json:"ip_list" dto:"ip_list" comment:"IP列表" validate:"required,valid_ipportlist"`
	WeightList        string `json:"weight_list" dto:"weight_list" comment:"权重列表" validate:"required,valid_weightlist"`
	ForbidList        string `json:"forbid_list" dto:"forbid_list" comment:"禁用IP列表" validate:"valid_iplist"`
}

type ServiceUpdateTcpInput struct {
	ServiceSingleByIdInput
	ServiceAddTcpInput
}

type ServiceAddGrpcInput struct {
	ServiceAddTcpInput
}

type ServiceUpdateGrpcInput struct {
	ServiceUpdateTcpInput
}

// ---------------------------------------------------------------------------------------------------------------------
type ServiceListItemOutput struct {
	ID          int64  `json:"id" dto:"id"`                     //id
	ServiceName string `json:"service_name" dto:"service_name"` //服务名称
	ServiceDesc string `json:"service_desc" dto:"service_desc"` //服务描述
	LoadType    int    `json:"load_type" dto:"load_type"`       //类型
	ServiceAddr string `json:"service_addr" dto:"service_addr"` //服务地址
	Qps         int64  `json:"qps" dto:"qps"`                   //qps
	Qpd         int64  `json:"qpd" dto:"qpd"`                   //qpd
	TotalNode   int    `json:"total_node" dto:"total_node"`     //节点数
}

type ServiceListOutput struct {
	Total int64                   `json:"total" dto:"total" comment:"总数" example:"" binding:""` //总数
	List  []ServiceListItemOutput `json:"list" dto:"list" comment:"列表" example:"" binding:""`   //列表
}

type ServiceStatOutput struct {
	Today     []int64 `json:"today" dto:"today" comment:"今日流量" example:"" validate:""`         //列表
	Yesterday []int64 `json:"yesterday" dto:"yesterday" comment:"昨日流量" example:"" validate:""` //列表
}
