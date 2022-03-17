package controller

import (
	"GatewayCombat/global"
	appDao "GatewayCombat/service/api/app/dao"
	appDto "GatewayCombat/service/api/app/dto"
	"GatewayCombat/service/api/dashboard/dto"
	serviceDao "GatewayCombat/service/api/service/dao"
	serviceDto "GatewayCombat/service/api/service/dto"
	"GatewayCombat/service/grf"
	"time"

	"github.com/gin-gonic/gin"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/10 17:03
*/

type DashboardController struct{}

func DashboardRegister(group *gin.RouterGroup) {
	dashboard := &DashboardController{}
	group.GET("/panel_group_data", dashboard.PanelGroupData)
	group.GET("/flow_stat", dashboard.FlowStat)
	group.GET("/service_stat", dashboard.ServiceStat)
}

func (dc *DashboardController) PanelGroupData(c *gin.Context) {
	serviceInfo := &serviceDao.ServiceInfo{}
	_, serviceNum, err := serviceInfo.PageList(global.RDB, &serviceDto.ServiceListInput{PageSize: 1, PageNo: 1})
	if err != nil {
		grf.Handler500(c, "10001 -->"+err.Error(), nil)
		return
	}
	app := &appDao.App{}
	_, appNum, err := app.APPList(global.RDB, &appDto.APPListInput{PageNo: 1, PageSize: 1})
	if err != nil {
		grf.Handler500(c, "10002 -->"+err.Error(), nil)
		return
	}
	//counter, err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
	//if err != nil {
	//	middleware.ResponseError(c, 2003, err)
	//	return
	//}
	out := &dto.PanelGroupDataOutput{
		ServiceNum:      serviceNum,
		AppNum:          appNum,
		TodayRequestNum: 0,
		CurrentQPS:      0,
	}
	grf.Handler200(c, out)
}

func (dc *DashboardController) FlowStat(c *gin.Context) {
	//counter, err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
	//if err != nil {
	//	middleware.ResponseError(c, 2001, err)
	//	return
	//}
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

func (dc *DashboardController) ServiceStat(c *gin.Context) {
	serviceInfo := &serviceDao.ServiceInfo{}
	list, err := serviceInfo.GroupByLoadType(global.RDB)
	if err != nil {
		grf.Handler500(c, err.Error(), nil)
		return
	}
	var legend []string
	for index, item := range list {
		name, ok := global.LoadTypeMap[item.LoadType]
		if !ok {
			grf.Handler500(c, "load_type not found", nil)
			return
		}
		list[index].Name = name
		legend = append(legend, name)
	}
	out := &dto.DashServiceStatOutput{
		Legend: legend,
		Data:   list,
	}
	grf.Handler200(c, out)
}
