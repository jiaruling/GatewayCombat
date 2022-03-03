package grf

import (
	"GatewayCombat/global"
	"GatewayCombat/global/errInfo"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

/*
   功能说明: 错误处理
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 15:56
*/

func Handler200(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": errInfo.SUCCESS, "data": data})
	return
}

func Handler201(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "msg": errInfo.SUCCESS, "data": data})
	return
}

func Handler204(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusNoContent, "msg": errInfo.SUCCESS, "data": nil})
	return
}

func Handler400(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": msg, "data": data})
	return
}

func Handler404(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "msg": errInfo.NotFound, "data": nil})
	return
}

func Handler500(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": msg, "data": data})
	return
}

// 表单验证失败
func FormsVerifyFailed(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if ok {
		// validators.ValidationErrors类型错误则进行翻译
		// 并使用removeTopStruct函数去除字段名中的结构体名称标识
		data, err := RemoveTopStruct(errs.Translate(global.Trans))
		if err != nil {
			Handler500(c, errInfo.TransError+err.Error(), nil)
			return
		}
		Handler400(c, string(data), nil)
		return
	} else {
		// 非validator.ValidationErrors类型错误直接返回
		Handler400(c, err.Error(), nil)
		return
	}

}
