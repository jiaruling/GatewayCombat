package dto

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1/19 16:46
*/

type FileContent struct {
	FileName string `json:"filename" binding:"required"`
	Content  string `json:"content" binding:"required"`
}
