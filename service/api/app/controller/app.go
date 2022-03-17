package controller

import (
	"GatewayCombat/global"
	"GatewayCombat/global/errInfo"
	"GatewayCombat/service/api/app/dao"
	"GatewayCombat/service/api/app/dto"
	"GatewayCombat/service/grf"
	"GatewayCombat/service/public"
	"GatewayCombat/utils"
	"time"

	"github.com/gin-gonic/gin"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/10 16:51
*/

type AppController struct{}

func AppRegister(router *gin.RouterGroup) {
	app := &AppController{}
	router.GET("/list", app.AppList)
	router.GET("/detail", app.AppDetail)
	router.GET("/delete", app.AppDelete)
	router.POST("/add", app.AppAdd)
	router.POST("/update", app.AppUpdate)
	router.GET("/stat", app.AppStatistics)
}

func (app *AppController) AppList(c *gin.Context) {
	params := &dto.APPListInput{}
	if err := c.ShouldBind(params); err != nil {
		public.FormsVerifyFailed(c, err)
		return
	}
	info := &dao.App{}
	list, total, err := info.APPList(global.RDB, params)
	if err != nil {
		grf.Handler500(c, "10001 --> "+err.Error(), nil)
		return
	}

	outputList := []dto.APPListItemOutput{}
	for _, item := range list {
		//appCounter, err := public.FlowCounterHandler.GetCounter(public.FlowAppPrefix + item.AppID)
		//if err != nil {
		//	grf.Handler500(c, "10002 --> "+err.Error(), nil)
		//	c.Abort()
		//	return
		//}
		outputList = append(outputList, dto.APPListItemOutput{
			ID:       item.ID,
			AppID:    item.AppID,
			Name:     item.Name,
			Secret:   item.Secret,
			WhiteIPS: item.WhiteIPS,
			Qpd:      item.Qpd,
			Qps:      item.Qps,
			//RealQpd:  appCounter.TotalCount,
			//RealQps:  appCounter.QPS,
		})
	}
	output := dto.APPListOutput{
		List:     outputList,
		Total:    total,
		Page:     int64(params.PageNo),
		PageSize: int64(params.PageSize),
	}
	grf.Handler200(c, output)
	return
}

func (app *AppController) AppDetail(c *gin.Context) {
	params := &dto.APPSingleByIdInput{}
	if err := c.ShouldBind(params); err != nil {
		public.FormsVerifyFailed(c, err)
		return
	}
	search := &dao.App{
		ID: params.ID,
	}
	detail, err := search.Find(global.RDB, search)
	if err != nil {
		grf.Handler500(c, "10001 --> "+err.Error(), nil)
		return
	}
	grf.Handler200(c, detail)
	return
}

func (app *AppController) AppDelete(c *gin.Context) {
	params := &dto.APPSingleByIdInput{}
	if err := c.ShouldBind(params); err != nil {
		public.FormsVerifyFailed(c, err)
		return
	}
	search := &dao.App{
		ID: params.ID,
	}
	info, err := search.Find(global.RDB, search)
	if err != nil {
		grf.Handler500(c, "10001 --> "+err.Error(), nil)
		return
	}
	info.DeletedAt = time.Now().Unix()
	if err := info.Save(global.WDB); err != nil {
		grf.Handler500(c, "10002 --> "+err.Error(), nil)
		return
	}
	grf.Handler204(c)
	return
}

func (app *AppController) AppAdd(c *gin.Context) {
	params := &dto.APPAddHttpInput{}
	if err := c.ShouldBind(params); err != nil {
		public.FormsVerifyFailed(c, err)
		return
	}

	//验证app_id是否被占用
	search := &dao.App{
		AppID: params.AppID,
	}
	if _, err := search.Find(global.RDB, search); err == nil {
		grf.Handler400(c, "租户ID被占用，请重新输入", nil)
		return
	}
	if params.Secret == "" {
		params.Secret = utils.MD5(params.AppID)
	}
	tx := global.WDB
	info := &dao.App{
		AppID:    params.AppID,
		Name:     params.Name,
		Secret:   params.Secret,
		WhiteIPS: params.WhiteIPS,
		Qps:      params.Qps,
		Qpd:      params.Qpd,
	}
	if err := info.Save(tx); err != nil {
		grf.Handler500(c, err.Error(), nil)
		return
	}
	grf.Handler200(c, errInfo.SUCCESS)
	return
}

func (app *AppController) AppUpdate(c *gin.Context) {
	params := &dto.APPUpdateHttpInput{}
	if err := c.ShouldBind(params); err != nil {
		public.FormsVerifyFailed(c, err)
		return
	}
	search := &dao.App{
		ID: params.ID,
	}
	info, err := search.Find(global.RDB, search)
	if err != nil {
		grf.Handler500(c, "10001 --> "+err.Error(), nil)
		return
	}
	if params.Secret == "" {
		params.Secret = utils.MD5(params.AppID)
	}
	info.Name = params.Name
	info.Secret = params.Secret
	info.WhiteIPS = params.WhiteIPS
	info.Qps = params.Qps
	info.Qpd = params.Qpd
	if err := info.Save(global.WDB); err != nil {
		grf.Handler500(c, "10002 --> "+err.Error(), nil)
		return
	}
	grf.Handler200(c, errInfo.SUCCESS)
	return
}

func (app *AppController) AppStatistics(c *gin.Context) {
	params := &dto.APPSingleByIdInput{}
	if err := c.ShouldBind(params); err != nil {
		public.FormsVerifyFailed(c, err)
		return
	}

	//今日流量全天小时级访问统计
	var todayStat []int64
	for i := 0; i <= time.Now().In(global.TimeLocation).Hour(); i++ {
		todayStat = append(todayStat, 0)
	}

	//昨日流量全天小时级访问统计
	var yesterdayStat []int64
	for i := 0; i <= 23; i++ {
		yesterdayStat = append(yesterdayStat, 0)
	}
	stat := dto.StatisticsOutput{
		Today:     todayStat,
		Yesterday: yesterdayStat,
	}
	grf.Handler200(c, stat)
	return
}
