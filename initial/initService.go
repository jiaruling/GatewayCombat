package initial

import (
	"GatewayCombat/global"
	"context"
	"log"
	"net/http"
	"time"
)

/*
   功能说明: 启动服务
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/19 14:45
*/

func InitService() {
	global.HttpSrvHandler = &http.Server{
		Addr:           ":8880",
		Handler:        global.GinRouter.Router,
		ReadTimeout:    time.Duration(1) * time.Second,
		WriteTimeout:   time.Duration(1) * time.Second,
		MaxHeaderBytes: 1 << uint(20),
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n", ":8880")
		if err := global.HttpSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", ":8880", err)
		}
	}()
	//go func() {
	//	if err := global.GinRouter.Router.Run(fmt.Sprintf("0.0.0.0:%v", 8080)); err != nil {
	//		log.Fatalln("http服务启动失败: ", err.Error())
	//	}
	//}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := global.HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
