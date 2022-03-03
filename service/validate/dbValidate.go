package validate

import (
	"GatewayCombat/global"
	"fmt"
	"regexp"

	"github.com/jmoiron/sqlx"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

// 自定义数据库连接字符串验证器
func ValidatorDBString(fl validator.FieldLevel) bool {
	// 获取参数数据
	dbString := fl.Field().String()
	// 使用正则表达式判断数据是否合法
	ok, _ := regexp.MatchString(`.*:.*@tcp\(.*:.*\)/.*\?$`, dbString)
	// 返回验证结果
	return ok
}

// 测试数据库是否能连接成功
func ValidatorDBPing(fl validator.FieldLevel) bool {
	// 获取参数数据
	dbString := fl.Field().String()
	// 测试连接数据库 生产环境使用
	if global.Config.RunMode != "debug" {
		database, err := sqlx.Open("mysql", fmt.Sprintf("%s%s", dbString, global.Config.MySQL.Parameter))
		if err != nil {
			return false
		}
		// 注意这行代码要写在上面err判断的下面
		_ = database.Close()
	}
	// 返回验证结果
	return true
}
