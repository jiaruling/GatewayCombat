package controller

import (
	"GatewayCombat/global"
	"GatewayCombat/service/api/service/dao"
	"GatewayCombat/service/api/service/dto"
	"GatewayCombat/service/grf"
	"GatewayCombat/service/public"
	"fmt"

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

func (service *ServiceController) ServiceList(c *gin.Context) {
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

func (service *ServiceController) ServiceDelete(c *gin.Context) {
	s := dao.Service
	s.M = new(dao.ServiceInfo)
	s.DeleteViewAPI(c)
	return
}

func (service *ServiceController) ServiceDetail(c *gin.Context) {
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

func (service *ServiceController) ServiceStat(c *gin.Context) {
	// 请求参数表单验证
	params := &dto.ServiceSingleByIdInput{}
	if err := c.ShouldBind(params); err != nil {
		public.FormsVerifyFailed(c, err)
		return
	}
}

func (service *ServiceController) ServiceAddHTTP(c *gin.Context) {

}

func (service *ServiceController) ServiceUpdateHTTP(c *gin.Context) {

}

func (service *ServiceController) ServiceAddTcp(c *gin.Context) {

}

func (service *ServiceController) ServiceUpdateTcp(c *gin.Context) {

}

func (service *ServiceController) ServiceAddGrpc(c *gin.Context) {

}

func (service *ServiceController) ServiceUpdateGrpc(c *gin.Context) {

}
