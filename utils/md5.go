package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
)

// 小写
func Md5Encode(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

// 大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// 加密
func EncryptionPassword(salt string, password string) string {
	return Md5Encode(salt + password)
}

// 解密
func DecryptPassword(salt string, password string, hashpassword string) bool {
	return Md5Encode(salt+password) == hashpassword
}

// 生成随机字节
func MakeRandomSalt(length int) string {
	randomByte := make([]byte, length)
	_, err := rand.Read(randomByte)
	if err != nil {
		fmt.Println("生成随机数失败：", err)
	}
	return hex.EncodeToString(randomByte)
}
