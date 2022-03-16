package public

import (
	"GatewayCombat/global"
	"GatewayCombat/service/grf"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 表单验证失败
func FormsVerifyFailed(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if ok {
		// validators.ValidationErrors类型错误则进行翻译
		// 并使用removeTopStruct函数去除字段名中的结构体名称标识
		grf.Handler400(c, "参数验证失败", RemoveTopStruct(errs.Translate(global.Trans)))
	} else {
		// 非validator.ValidationErrors类型错误直接返回
		grf.Handler400(c, err.Error(), nil)
	}
}

/*
	去除字符串前面的结构体名称:
		{User.Name: "张三"} --> {Name: "张三"}
*/
func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
