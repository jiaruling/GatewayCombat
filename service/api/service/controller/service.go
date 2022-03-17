package controller

import (
	"GatewayCombat/global"
	"GatewayCombat/global/errInfo"
	"GatewayCombat/service/api/service/dao"
	"GatewayCombat/service/api/service/dto"
	"GatewayCombat/service/grf"
	"GatewayCombat/service/public"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/10 14:42
*/

type ServiceController struct{}

func ServiceRegister(group *gin.RouterGroup) {
	service := &ServiceController{}
	group.GET("/list", service.ServiceList)
	group.DELETE("/delete", service.ServiceDelete)
	group.GET("/detail", service.ServiceDetail)
	group.GET("/stat", service.ServiceStat)
	add := group.Group("/add")
	{
		add.POST("/http", service.ServiceAddHTTP)
		add.POST("/tcp", service.ServiceAddTcp)
		add.POST("/grpc", service.ServiceAddGrpc)
	}
	update := group.Group("/update")
	{
		update.PUT("/http", service.ServiceUpdateHTTP)
		update.PUT("/tcp", service.ServiceUpdateTcp)
		update.PUT("/grpc", service.ServiceUpdateGrpc)
	}
}

func (sc *ServiceController) ServiceList(c *gin.Context) {
	// 请求参数表单验证
	params := &dto.ServiceListInput{}
	if err := c.ShouldBind(params); err != nil {
		public.FormsVerifyFailed(c, err)
		return
	}
	serviceInfo := &dao.ServiceInfo{}
	list, total, err := serviceInfo.PageList(global.RDB, params)
	if err != nil {
		grf.Handler500(c, err.Error(), err)
		return
	}
	//格式化输出信息
	var outList []dto.ServiceListItemOutput
	for _, listItem := range list {
		serviceDetail, err := serviceInfo.ServiceDetail(global.RDB, &listItem)
		if err != nil {
			grf.Handler500(c, err.Error(), err)
			return
		}
		//1、http后缀接入 clusterIP+clusterPort+path
		//2、http域名接入 domain
		//3、tcp、grpc接入 clusterIP+servicePort
		serviceAddr := "unknow"
		clusterIP := global.Config.Cluster.Ip
		clusterPort := global.Config.Cluster.Port
		clusterSSLPort := global.Config.Cluster.SSLPort
		if serviceDetail.Info.LoadType == global.LoadTypeHTTP &&
			serviceDetail.HTTPRule.RuleType == global.HTTPRuleTypePrefixURL &&
			serviceDetail.HTTPRule.NeedHttps == 1 {
			serviceAddr = fmt.Sprintf("%s:%d%s", clusterIP, clusterSSLPort, serviceDetail.HTTPRule.Rule)
		}
		if serviceDetail.Info.LoadType == global.LoadTypeHTTP &&
			serviceDetail.HTTPRule.RuleType == global.HTTPRuleTypePrefixURL &&
			serviceDetail.HTTPRule.NeedHttps == 0 {
			serviceAddr = fmt.Sprintf("%s:%d%s", clusterIP, clusterPort, serviceDetail.HTTPRule.Rule)
		}
		if serviceDetail.Info.LoadType == global.LoadTypeHTTP &&
			serviceDetail.HTTPRule.RuleType == global.HTTPRuleTypeDomain {
			serviceAddr = serviceDetail.HTTPRule.Rule
		}
		if serviceDetail.Info.LoadType == global.LoadTypeTCP {
			serviceAddr = fmt.Sprintf("%s:%d", clusterIP, serviceDetail.TCPRule.Port)
		}
		if serviceDetail.Info.LoadType == global.LoadTypeGRPC {
			serviceAddr = fmt.Sprintf("%s:%d", clusterIP, serviceDetail.GRPCRule.Port)
		}
		loadBalance := &dao.LoadBalance{}
		ipList := loadBalance.GetIPListByModel()
		//counter, err := public.FlowCounterHandler.GetCounter(public.FlowServicePrefix + listItem.ServiceName)
		//if err != nil {
		//	grf.Handler500(c, err.Error(), err)
		//	return
		//}
		outItem := dto.ServiceListItemOutput{
			ID:          listItem.ID,
			LoadType:    listItem.LoadType,
			ServiceName: listItem.ServiceName,
			ServiceDesc: listItem.ServiceDesc,
			ServiceAddr: serviceAddr,
			Qps:         0,
			Qpd:         0,
			TotalNode:   len(ipList),
		}
		outList = append(outList, outItem)
	}
	out := &dto.ServiceListOutput{
		Total: total,
		List:  outList,
	}
	grf.Handler200(c, out)
	return
}

func (sc *ServiceController) ServiceDelete(c *gin.Context) {
	s := dao.Service
	s.M = new(dao.ServiceInfo)
	s.DeleteViewAPI(c)
	return
}

func (sc *ServiceController) ServiceDetail(c *gin.Context) {
	// 请求参数表单验证
	params := &dto.ServiceSingleByIdInput{}
	if err := c.ShouldBind(params); err != nil {
		public.FormsVerifyFailed(c, err)
		return
	}
	// 读取基本信息
	serviceInfo := &dao.ServiceInfo{ID: params.ID}
	serviceInfo, err := serviceInfo.Find(global.RDB, serviceInfo)
	if err != nil {
		grf.Handler500(c, err.Error(), err)
		return
	}
	serviceDetail, err := serviceInfo.ServiceDetail(global.RDB, serviceInfo)
	if err != nil {
		grf.Handler500(c, err.Error(), err)
		return
	}
	grf.Handler200(c, serviceDetail)
	return
}

func (sc *ServiceController) ServiceStat(c *gin.Context) {
	// 请求参数表单验证
	params := &dto.ServiceSingleByIdInput{}
	if err := c.ShouldBind(params); err != nil {
		public.FormsVerifyFailed(c, err)
		return
	}

	var todayList []int64
	currentTime := time.Now()
	for i := 0; i <= currentTime.Hour(); i++ {
		todayList = append(todayList, 0)
	}

	var yesterdayList []int64
	for i := 0; i <= 23; i++ {
		yesterdayList = append(yesterdayList, 0)
	}
	grf.Handler200(c, &dto.ServiceStatOutput{
		Today:     todayList,
		Yesterday: yesterdayList,
	})
}

func (sc *ServiceController) ServiceAddHTTP(c *gin.Context) {
	params := &dto.ServiceAddHTTPInput{}
	if err := c.ShouldBind(params); err != nil {
		public.FormsVerifyFailed(c, err)
		return
	}

	if len(strings.Split(params.IpList, ",")) != len(strings.Split(params.WeightList, ",")) {
		grf.Handler400(c, "IP列表与权重列表数量不一致", nil)
		return
	}

	tx := global.WDB
	tx = tx.Begin()
	serviceInfo := &dao.ServiceInfo{ServiceName: params.ServiceName}
	if _, err := serviceInfo.Find(tx, serviceInfo); err == nil {
		tx.Rollback()
		grf.Handler500(c, "服务已存在", nil)
		return
	}

	httpUrl := &dao.HttpRule{RuleType: params.RuleType, Rule: params.Rule}
	if _, err := httpUrl.Find(tx, httpUrl); err == nil {
		tx.Rollback()
		grf.Handler500(c, "服务接入前缀或域名已存在", nil)
		return
	}

	serviceModel := &dao.ServiceInfo{
		ServiceName: params.ServiceName,
		ServiceDesc: params.ServiceDesc,
	}
	if err := serviceModel.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10001 --> "+err.Error(), nil)
		return
	}
	//serviceModel.ID
	httpRule := &dao.HttpRule{
		ServiceID:      serviceModel.ID,
		RuleType:       params.RuleType,
		Rule:           params.Rule,
		NeedHttps:      params.NeedHttps,
		NeedStripUri:   params.NeedStripUri,
		NeedWebsocket:  params.NeedWebsocket,
		UrlRewrite:     params.UrlRewrite,
		HeaderTransfor: params.HeaderTransfor,
	}
	if err := httpRule.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10002 --> "+err.Error(), nil)
		return
	}

	accessControl := &dao.AccessControl{
		ServiceID:         serviceModel.ID,
		OpenAuth:          params.OpenAuth,
		BlackList:         params.BlackList,
		WhiteList:         params.WhiteList,
		ClientIPFlowLimit: params.ClientipFlowLimit,
		ServiceFlowLimit:  params.ServiceFlowLimit,
	}
	if err := accessControl.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10003 --> "+err.Error(), nil)
		return
	}

	loadBalance := &dao.LoadBalance{
		ServiceID:              serviceModel.ID,
		RoundType:              params.RoundType,
		IpList:                 params.IpList,
		WeightList:             params.WeightList,
		UpstreamConnectTimeout: params.UpstreamConnectTimeout,
		UpstreamHeaderTimeout:  params.UpstreamHeaderTimeout,
		UpstreamIdleTimeout:    params.UpstreamIdleTimeout,
		UpstreamMaxIdle:        params.UpstreamMaxIdle,
	}
	if err := loadBalance.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10004 --> "+err.Error(), nil)
		return
	}
	tx.Commit()
	grf.Handler200(c, errInfo.SUCCESS)
	return
}

func (sc *ServiceController) ServiceUpdateHTTP(c *gin.Context) {
	params := &dto.ServiceUpdateHTTPInput{}
	if err := c.ShouldBind(params); err != nil {
		public.FormsVerifyFailed(c, err)
		return
	}

	if len(strings.Split(params.IpList, ",")) != len(strings.Split(params.WeightList, ",")) {
		grf.Handler400(c, "IP列表与权重列表数量不一致", nil)
		return
	}

	tx := global.WDB
	tx = tx.Begin()
	serviceInfo := &dao.ServiceInfo{ServiceName: params.ServiceName}
	serviceInfo, err := serviceInfo.Find(tx, serviceInfo)
	if err != nil {
		tx.Rollback()
		grf.Handler500(c, "服务不存在", nil)
		return
	}
	serviceDetail, err := serviceInfo.ServiceDetail(tx, serviceInfo)
	if err != nil {
		tx.Rollback()
		grf.Handler500(c, "20001 --> "+err.Error(), nil)
		return
	}

	info := serviceDetail.Info
	info.ServiceName = params.ServiceName
	info.ServiceDesc = params.ServiceDesc
	if err := info.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "20002 --> "+err.Error(), nil)
		return
	}

	httpRule := serviceDetail.HTTPRule
	httpRule.NeedHttps = params.NeedHttps
	httpRule.NeedStripUri = params.NeedStripUri
	httpRule.NeedWebsocket = params.NeedWebsocket
	httpRule.UrlRewrite = params.UrlRewrite
	httpRule.HeaderTransfor = params.HeaderTransfor
	if err := httpRule.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "20003 --> "+err.Error(), nil)
		return
	}

	accessControl := serviceDetail.AccessControl
	accessControl.OpenAuth = params.OpenAuth
	accessControl.BlackList = params.BlackList
	accessControl.WhiteList = params.WhiteList
	accessControl.ClientIPFlowLimit = params.ClientipFlowLimit
	accessControl.ServiceFlowLimit = params.ServiceFlowLimit
	if err := accessControl.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "20004 --> "+err.Error(), nil)
		return
	}

	loadbalance := serviceDetail.LoadBalance
	loadbalance.RoundType = params.RoundType
	loadbalance.IpList = params.IpList
	loadbalance.WeightList = params.WeightList
	loadbalance.UpstreamConnectTimeout = params.UpstreamConnectTimeout
	loadbalance.UpstreamHeaderTimeout = params.UpstreamHeaderTimeout
	loadbalance.UpstreamIdleTimeout = params.UpstreamIdleTimeout
	loadbalance.UpstreamMaxIdle = params.UpstreamMaxIdle
	if err := loadbalance.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "20005 --> "+err.Error(), nil)
		return
	}
	tx.Commit()
	grf.Handler200(c, errInfo.SUCCESS)
	return
}

func (sc *ServiceController) ServiceAddTcp(c *gin.Context) {
	params := &dto.ServiceAddTcpInput{}
	if err := c.ShouldBind(params); err != nil {
		public.FormsVerifyFailed(c, err)
		return
	}

	//验证 service_name 是否被占用
	infoSearch := &dao.ServiceInfo{
		ServiceName: params.ServiceName,
		DeleteAt:    nil,
	}
	if _, err := infoSearch.Find(global.RDB, infoSearch); err == nil {
		grf.Handler400(c, "服务名被占用，请重新输入", nil)
		return
	}

	//验证端口是否被占用?
	tcpRuleSearch := &dao.TcpRule{
		Port: params.Port,
	}
	if _, err := tcpRuleSearch.Find(global.RDB, tcpRuleSearch); err == nil {
		grf.Handler400(c, "服务端口被占用，请重新输入", nil)
		return
	}
	grpcRuleSearch := &dao.GrpcRule{
		Port: params.Port,
	}
	if _, err := grpcRuleSearch.Find(global.RDB, grpcRuleSearch); err == nil {
		grf.Handler400(c, "服务端口被占用，请重新输入", nil)
		return
	}

	//ip与权重数量一致
	if len(strings.Split(params.IpList, ",")) != len(strings.Split(params.WeightList, ",")) {
		grf.Handler400(c, "ip列表与权重设置不匹配", nil)
		return
	}

	tx := global.WDB.Begin()
	info := &dao.ServiceInfo{
		LoadType:    global.LoadTypeTCP,
		ServiceName: params.ServiceName,
		ServiceDesc: params.ServiceDesc,
	}
	if err := info.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10001 --> "+err.Error(), nil)
		return
	}
	loadBalance := &dao.LoadBalance{
		ServiceID:  info.ID,
		RoundType:  params.RoundType,
		IpList:     params.IpList,
		WeightList: params.WeightList,
		ForbidList: params.ForbidList,
	}
	if err := loadBalance.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10002 --> "+err.Error(), nil)
		return
	}

	httpRule := &dao.TcpRule{
		ServiceID: info.ID,
		Port:      params.Port,
	}
	if err := httpRule.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10003 --> "+err.Error(), nil)
		return
	}

	accessControl := &dao.AccessControl{
		ServiceID:         info.ID,
		OpenAuth:          params.OpenAuth,
		BlackList:         params.BlackList,
		WhiteList:         params.WhiteList,
		WhiteHostName:     params.WhiteHostName,
		ClientIPFlowLimit: params.ClientIPFlowLimit,
		ServiceFlowLimit:  params.ServiceFlowLimit,
	}
	if err := accessControl.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10004 --> "+err.Error(), nil)
		return
	}
	tx.Commit()
	grf.Handler200(c, errInfo.SUCCESS)
	return
}

func (sc *ServiceController) ServiceUpdateTcp(c *gin.Context) {
	params := &dto.ServiceUpdateTcpInput{}
	if err := c.ShouldBind(params); err != nil {
		public.FormsVerifyFailed(c, err)
		return
	}

	//ip与权重数量一致
	if len(strings.Split(params.IpList, ",")) != len(strings.Split(params.WeightList, ",")) {
		grf.Handler400(c, "IP列表与权重列表数量不一致", nil)
		return
	}

	service := &dao.ServiceInfo{
		ID: params.ID,
	}
	detail, err := service.ServiceDetail(global.RDB, service)
	if err != nil {
		grf.Handler500(c, "10001 --> "+err.Error(), nil)
		return
	}

	tx := global.WDB.Begin()
	info := detail.Info
	info.ServiceDesc = params.ServiceDesc
	if err := info.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10002 --> "+err.Error(), nil)
		return
	}

	loadBalance := &dao.LoadBalance{}
	if detail.LoadBalance != nil {
		loadBalance = detail.LoadBalance
	}
	loadBalance.ServiceID = info.ID
	loadBalance.RoundType = params.RoundType
	loadBalance.IpList = params.IpList
	loadBalance.WeightList = params.WeightList
	loadBalance.ForbidList = params.ForbidList
	if err := loadBalance.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10003 --> "+err.Error(), nil)
		return
	}

	tcpRule := &dao.TcpRule{}
	if detail.TCPRule != nil {
		tcpRule = detail.TCPRule
	}
	tcpRule.ServiceID = info.ID
	tcpRule.Port = params.Port
	if err := tcpRule.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10004 --> "+err.Error(), nil)
		return
	}

	accessControl := &dao.AccessControl{}
	if detail.AccessControl != nil {
		accessControl = detail.AccessControl
	}
	accessControl.ServiceID = info.ID
	accessControl.OpenAuth = params.OpenAuth
	accessControl.BlackList = params.BlackList
	accessControl.WhiteList = params.WhiteList
	accessControl.WhiteHostName = params.WhiteHostName
	accessControl.ClientIPFlowLimit = params.ClientIPFlowLimit
	accessControl.ServiceFlowLimit = params.ServiceFlowLimit
	if err := accessControl.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10005 --> "+err.Error(), nil)
		return
	}
	tx.Commit()
	grf.Handler200(c, errInfo.SUCCESS)
	return
}

func (sc *ServiceController) ServiceAddGrpc(c *gin.Context) {
	params := &dto.ServiceAddGrpcInput{}
	if err := c.ShouldBind(params); err != nil {
		public.FormsVerifyFailed(c, err)
		return
	}

	//验证 service_name 是否被占用
	infoSearch := &dao.ServiceInfo{
		ServiceName: params.ServiceName,
		DeleteAt:    nil,
	}
	if _, err := infoSearch.Find(global.RDB, infoSearch); err == nil {
		grf.Handler400(c, "服务名被占用，请重新输入", nil)
		return
	}

	//验证端口是否被占用?
	tcpRuleSearch := &dao.TcpRule{
		Port: params.Port,
	}
	if _, err := tcpRuleSearch.Find(global.RDB, tcpRuleSearch); err == nil {
		grf.Handler400(c, "服务端口被占用，请重新输入", nil)
		return
	}
	grpcRuleSearch := &dao.GrpcRule{
		Port: params.Port,
	}
	if _, err := grpcRuleSearch.Find(global.RDB, grpcRuleSearch); err == nil {
		grf.Handler400(c, "服务端口被占用，请重新输入", nil)
		return
	}

	//ip与权重数量一致
	if len(strings.Split(params.IpList, ",")) != len(strings.Split(params.WeightList, ",")) {
		grf.Handler400(c, "ip列表与权重设置不匹配", nil)
		return
	}

	tx := global.WDB.Begin()
	info := &dao.ServiceInfo{
		LoadType:    global.LoadTypeGRPC,
		ServiceName: params.ServiceName,
		ServiceDesc: params.ServiceDesc,
	}
	if err := info.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10001 --> "+err.Error(), nil)
		return
	}

	loadBalance := &dao.LoadBalance{
		ServiceID:  info.ID,
		RoundType:  params.RoundType,
		IpList:     params.IpList,
		WeightList: params.WeightList,
		ForbidList: params.ForbidList,
	}
	if err := loadBalance.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10002 --> "+err.Error(), nil)
		return
	}

	grpcRule := &dao.GrpcRule{
		ServiceID:      info.ID,
		Port:           params.Port,
		HeaderTransfor: params.HeaderTransfor,
	}
	if err := grpcRule.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10003 --> "+err.Error(), nil)
		return
	}

	accessControl := &dao.AccessControl{
		ServiceID:         info.ID,
		OpenAuth:          params.OpenAuth,
		BlackList:         params.BlackList,
		WhiteList:         params.WhiteList,
		WhiteHostName:     params.WhiteHostName,
		ClientIPFlowLimit: params.ClientIPFlowLimit,
		ServiceFlowLimit:  params.ServiceFlowLimit,
	}
	if err := accessControl.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10004 --> "+err.Error(), nil)
		return
	}
	tx.Commit()
	grf.Handler200(c, errInfo.SUCCESS)
	return
}

func (sc *ServiceController) ServiceUpdateGrpc(c *gin.Context) {
	params := &dto.ServiceUpdateGrpcInput{}
	if err := c.ShouldBind(params); err != nil {
		public.FormsVerifyFailed(c, err)
		return
	}

	//ip与权重数量一致
	if len(strings.Split(params.IpList, ",")) != len(strings.Split(params.WeightList, ",")) {
		grf.Handler400(c, "ip列表与权重设置不匹配", nil)
		return
	}

	service := &dao.ServiceInfo{
		ID: params.ID,
	}
	detail, err := service.ServiceDetail(global.RDB, service)
	if err != nil {
		grf.Handler500(c, "10001 --> "+err.Error(), nil)
		return
	}

	tx := global.WDB.Begin()
	info := detail.Info
	info.ServiceDesc = params.ServiceDesc
	if err := info.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10002 --> "+err.Error(), nil)
		return
	}

	loadBalance := &dao.LoadBalance{}
	if detail.LoadBalance != nil {
		loadBalance = detail.LoadBalance
	}
	loadBalance.ServiceID = info.ID
	loadBalance.RoundType = params.RoundType
	loadBalance.IpList = params.IpList
	loadBalance.WeightList = params.WeightList
	loadBalance.ForbidList = params.ForbidList
	if err := loadBalance.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10003 --> "+err.Error(), nil)
		return
	}

	grpcRule := &dao.GrpcRule{}
	if detail.GRPCRule != nil {
		grpcRule = detail.GRPCRule
	}
	grpcRule.ServiceID = info.ID
	//grpcRule.Port = params.Port
	grpcRule.HeaderTransfor = params.HeaderTransfor
	if err := grpcRule.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10004 --> "+err.Error(), nil)
		return
	}

	accessControl := &dao.AccessControl{}
	if detail.AccessControl != nil {
		accessControl = detail.AccessControl
	}
	accessControl.ServiceID = info.ID
	accessControl.OpenAuth = params.OpenAuth
	accessControl.BlackList = params.BlackList
	accessControl.WhiteList = params.WhiteList
	accessControl.WhiteHostName = params.WhiteHostName
	accessControl.ClientIPFlowLimit = params.ClientIPFlowLimit
	accessControl.ServiceFlowLimit = params.ServiceFlowLimit
	if err := accessControl.Save(tx); err != nil {
		tx.Rollback()
		grf.Handler500(c, "10005 --> "+err.Error(), nil)
		return
	}
	tx.Commit()
	grf.Handler200(c, errInfo.SUCCESS)
	return
}
