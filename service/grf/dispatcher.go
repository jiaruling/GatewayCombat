package grf

import (
	"GatewayCombat/global/errInfo"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
   功能说明: 分发器
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 15:30
*/

func Dispatcher(c *gin.Context, m ViewAPI) {
	if !m.GetModelIsInit() {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"code": http.StatusInternalServerError, "msg": "模型没有初始化, 不能进行操作", "data": ""},
		)
		return
	}
	if inFields(c.Request.Method, m.GetAllowMethod()) {
		switch c.Request.Method {
		case "GET":
			//fmt.Println(c.Param("id"))
			if len(c.Param("id")) == 1 {
				m.ListViewAPI(c)
			} else {
				m.RetrieveViewAPI(c)
			}
		case "POST":
			m.CreateViewAPI(c)
		case "PUT":
			m.UpdateViewAPI(c)
		case "DELETE":
			m.DeleteViewAPI(c)
		default:
			Handler400(c, errInfo.RequestNotAllow, "")
		}
	} else {
		Handler400(c, errInfo.RequestNotAllow, "")
	}
	return
}
