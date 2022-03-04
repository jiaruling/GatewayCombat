package global

/*
   功能说明: 常量
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/18 10:52
*/

const (
	METHOD_POST    = "POST"
	METHOD_GET     = "GET"
	METHOD_PUT     = "PUT"
	METHOD_DELETE  = "DELETE"
	FILEPATH       = "./static"
	LogFilePath    = "./log"
	ConfigFilePath = "./config"
	MaxUploadSize  = 5 * 1024 * 2014 // 5 MB

	AdminSessionInfoKey = "AdminSessionInfoKey"
)
