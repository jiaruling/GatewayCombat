package initial

import (
	"GatewayCombat/global"
	"GatewayCombat/utils"
	"fmt"
	"os"
)

/*
   功能说明: 创建静态目录
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/19 14:56
*/

func InitCreateDir() {
	if !utils.Exists(global.FILEPATH) {
		if err := os.MkdirAll(global.FILEPATH, os.ModePerm); err != nil {
			fmt.Println("创建目录"+global.FILEPATH+"失败: ", err.Error())
			os.Exit(1)
		}
	}
	if !utils.Exists(global.LogFilePath) {
		if err := os.MkdirAll(global.LogFilePath, os.ModePerm); err != nil {
			fmt.Println("创建目录"+global.LogFilePath+"失败: ", err.Error())
			os.Exit(1)
		}
	}
}
