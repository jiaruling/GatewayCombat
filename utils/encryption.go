package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
)

/*
   功能说明: 加密
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/4 15:37
*/

// sha256加盐加密
func GenSaltPassword(salt, password string) string {
	s1 := sha256.New()
	s1.Write([]byte(password))
	str1 := fmt.Sprintf("%x", s1.Sum(nil))
	s2 := sha256.New()
	s2.Write([]byte(str1 + salt))
	return fmt.Sprintf("%x", s2.Sum(nil))
}

//MD5 md5加密
func MD5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
