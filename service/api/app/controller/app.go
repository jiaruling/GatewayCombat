package controller

import "github.com/gin-gonic/gin"

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/10 16:51
*/

type AppController struct{}

func AppRegister(router *gin.RouterGroup) {
	app := &AppController{}
	router.Any("/*id", app.App)
}

func (app *AppController) App(c *gin.Context) {

}
