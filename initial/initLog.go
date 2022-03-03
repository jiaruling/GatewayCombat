package initial

import (
	"GatewayCombat/global"
	"GatewayCombat/utils"
	"fmt"
	"log"
	"os"
)

// 初始化自带的日志库
func InitLog(LogPath []string) {
	// 删除历史日志文件
	deleteLogFile(LogPath)
	for i, item := range LogPath {
		logFile, err := os.OpenFile(item, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("open log file failed, err:", err)
			return
		}
		//defer logFile.Close()
		switch i {
		case 0:
			log.SetOutput(logFile)
			log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
			log.Println("*****************************************************************************************")
			log.Println("通用日志记录")
			log.Println("*****************************************************************************************")
		case 1:
			global.AccessLog = log.New(logFile, "<Access>: ", log.Lshortfile|log.Ldate|log.Ltime)
			global.AccessLog.Println("*****************************************************************************************")
			global.AccessLog.Println("访问日志记录")
			global.AccessLog.Println("*****************************************************************************************")
		case 2:
			global.SqlLog = log.New(logFile, "<SQL>: ", log.Lshortfile|log.Ldate|log.Ltime)
			global.SqlLog.Println("*****************************************************************************************")
			global.SqlLog.Println("SQL日志记录")
			global.SqlLog.Println("*****************************************************************************************")
		case 3:
			global.TaskLog = log.New(logFile, "<Task>: ", log.Lshortfile|log.Ldate|log.Ltime)
			global.TaskLog.Println("*****************************************************************************************")
			global.TaskLog.Println("后台任务日志记录")
			global.TaskLog.Println("*****************************************************************************************")
		}
	}
}

func deleteLogFile(LogList []string) {
	for _, item := range LogList {
		if utils.Exists(item) {
			_ = os.Remove(item)
		}
	}
}
