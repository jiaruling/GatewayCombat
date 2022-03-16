package global

import (
	. "GatewayCombat/global/config_struct"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
)

/*
   功能说明: 变量
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/18 10:52
*/

type Router struct {
	Router *gin.Engine
	V1     *gin.RouterGroup
}

var (
	Config         ServerConfig
	GinRouter      *Router
	Trans          ut.Translator
	RDB            *gorm.DB
	WDB            *gorm.DB
	AccessLog      *log.Logger
	SqlLog         *log.Logger
	TaskLog        *log.Logger
	LogPath        []string
	Validate       *validator.Validate
	Expires        time.Duration
	ETicker        *time.Ticker
	HttpSrvHandler *http.Server
	LoadTypeMap    map[int]string
	//Store        sessions.RedisStore

)

// 初始化全局变量
func init() {
	LogPath = []string{"./log/logs.log", "./log/access.log", "./log/sql.log", "./log/backend_task.log"}
	Validate = validator.New()
	Expires = 10 // 10s
	ETicker = time.NewTicker(Expires * time.Second)
	LoadTypeMap = map[int]string{
		LoadTypeHTTP: "HTTP",
		LoadTypeTCP:  "TCP",
		LoadTypeGRPC: "GRPC",
	}
}
