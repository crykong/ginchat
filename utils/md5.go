package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempstr := h.Sum(nil)
	return hex.EncodeToString(tempstr)
}

func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// 加密
func MakePassword(plainpwd, salt string) string {
	return MD5Encode(plainpwd + salt)
}

// 解密
func ValidPassword(plainpwd, salt string, password string) bool {
	return MD5Encode(plainpwd+salt) == password
}
