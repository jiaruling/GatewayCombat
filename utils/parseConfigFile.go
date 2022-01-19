package utils

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

/*
   功能说明: 加载配置文件并进行监听
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/18 10:41
*/

func ParseConfig(path string, obj interface{}) (err error) {
	// 读取配置文件，映射到结构体
	// 实例化viper对象
	v := viper.New()
	v.SetConfigFile(path)
	// 读取配置文件
	if err = v.ReadInConfig(); err != nil {
		return
	}
	// 反序列化为 struct 对象
	if err = v.Unmarshal(obj); err != nil {
		return
	}
	//fmt.Println(serverConfig)
	// viper的功能 -- 动态监控变化
	log.Println("开启配置 " + path + " 文件监听")
	go func() {
		// 开启监听功能
		v.WatchConfig()
		// 文件监听
		v.OnConfigChange(func(e fsnotify.Event) {
			// 打印变换的文件名
			fmt.Println("配置文件发生变化:", e.Name)
			_ = v.ReadInConfig() // 重新读取配置数据
			_ = v.Unmarshal(obj) // 将文件内容映射到结构体
			//fmt.Println(constant.INTERFACE)
		})
	}()
	return
}
