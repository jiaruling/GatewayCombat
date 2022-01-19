package main

import (
	"GoGinServerBestPractice/service/core"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 15:47
*/

func init() {
	database, err := gorm.Open(mysql.Open("root:abc123456@tcp(127.0.0.1:3306)/imooc"), &gorm.Config{})
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	//defer database.Close()  // 注意这行代码要写在上面err判断的下面
	core.RDB = database
	core.WDB = database
	core.GlobalPageMax = 5
	core.GlobalPageMin = 1
}