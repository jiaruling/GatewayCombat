package backend_task

import (
	"GatewayCombat/global"
	"math/rand"
	"time"
)

/*
   功能说明: 不定时间执行
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/18 17:43
*/
func task3() {
	// 启动时执行一次
	executeTask3()
	for {
		<-global.ETicker.C
		executeTask3()
		global.ETicker.Reset(global.Expires * time.Second)
	}
}

func executeTask3() {
	global.TaskLog.Println("task3", time.Now().String())
	global.Expires = time.Duration(rand.Intn(10))
}
