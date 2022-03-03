package main

import (
	"GatewayCombat/service/grf"

	"fmt"

	"github.com/gin-gonic/gin"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 15:08
*/

func main() {
	r := gin.Default()
	r.GET("/ping/*id", func(c *gin.Context) {
		id := c.Param("id")[1:]
		c.JSON(200, gin.H{
			"message": "pong",
			"data":    id,
		})
	})
	r.Any("/stu/*id", func(c *gin.Context) {
		// 通过全局变量赋值给局部变量，可以实现并发
		s := stu
		s.M = new(Student)
		grf.Dispatcher(s, c)
		fmt.Printf("stu: %p -- s:%p \n", &stu, &s)
		fmt.Printf("stu.M: %p -- s.M:%p \n", &stu.M, &s.M)
		fmt.Printf("stu.T: %p -- s.T:%p \n", &stu.Table, &s.Table)
		fmt.Printf("stu.M: %v -- s.M:%v \n", stu.M, s.M)
		fmt.Printf("-----------------------------------------------------\n")
		return
	})
	r.Run(":8002") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
