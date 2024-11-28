package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data)) // convert data into a byte slice
	tempStr := h.Sum(nil) // nil indicates no additional dat should be appended to the hash output
	return hex.EncodeToString(tempStr)
}

func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// 加密操作
func MakePassword(plainwd, salt string) string {
	return Md5Encode(plainwd + salt)
}

// 解密操作
func ValidPassword(plainwd, salt, password string) bool {
	return Md5Encode(plainwd+salt) == password
}
