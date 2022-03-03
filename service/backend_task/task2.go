package backend_task

import (
	"GatewayCombat/global"
	"time"
)

/*
   功能说明: 固定时间执行
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/18 17:43
*/
func task2() {
	// 启动时执行一次
	executeTask2()
	// 设置定时器, 每一分钟执行一次
	ticker := time.NewTicker(time.Duration(60) * time.Second)
	for {
		<-ticker.C
		executeTask2()
	}
}

func executeTask2() {
	global.TaskLog.Println("task2", time.Now().String())
}
