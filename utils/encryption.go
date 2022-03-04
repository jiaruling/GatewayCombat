package utils

import (
	"crypto/sha256"
	"fmt"
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
