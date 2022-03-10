package controller

import (
	"GatewayCombat/service/api/service/model"

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

}

func (service *ServiceController) ServiceDelete(c *gin.Context) {
	m := new(model.ServiceInfo)
	s := model.Service
	s.Table = m.TableName()
	s.M = m
	s.DeleteViewAPI(c)
	return
}

func (service *ServiceController) ServiceDetail(c *gin.Context) {

}

func (service *ServiceController) ServiceStat(c *gin.Context) {

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
