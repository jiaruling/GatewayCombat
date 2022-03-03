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

func task1() {
	// 开始时先执行一次
	executeTask1()
	// 设置定时器, 每半分钟执行一次
	ticker := time.NewTicker(time.Duration(30) * time.Second)
	for {
		<-ticker.C
		executeTask1()
	}
}

func executeTask1() {
	global.TaskLog.Println("task1", time.Now().String())
}
