package initial

import (
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	"GatewayCombat/global"
	"GatewayCombat/service/validate"
)

// 注册自定义验证器并进行错误翻译
func InitValidator() (err error) {
	// 注册自定义验证器并进行错误翻译
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证器
		if err = v.RegisterValidation("mobile", validate.ValidatorMobile); err != nil {
			return
		}
		if err = v.RegisterValidation("email", validate.ValidatorEmail); err != nil {
			return
		}
		if err = v.RegisterValidation("db_string", validate.ValidatorDBString); err != nil {
			return
		}
		if err = v.RegisterValidation("db_ping", validate.ValidatorDBPing); err != nil {
			return
		}

		/*
			对自定义验证器的错误信息进行翻译
			格式参考: https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go#L105
		*/
		if err = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0}非法的手机号码!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		}); err != nil {
			return
		}

		if err = v.RegisterTranslation("email", global.Trans, func(ut ut.Translator) error {
			return ut.Add("email", "{0}非法的邮箱!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("email", fe.Field())
			return t
		}); err != nil {
			return
		}

		if err = v.RegisterTranslation("db_string", global.Trans, func(ut ut.Translator) error {
			return ut.Add("db_string", "{0}非法的数据库链接!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("db_string", fe.Field())
			return t
		}); err != nil {
			return
		}

		if err = v.RegisterTranslation("db_ping", global.Trans, func(ut ut.Translator) error {
			return ut.Add("db_ping", "{0}连接数据库失败，请检查链接是否正确!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("db_ping", fe.Field())
			return t
		}); err != nil {
			return
		}
	}
	return
}
