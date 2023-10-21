package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// 小写
func Sha256Encode(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	tempstr := h.Sum(nil)
	return hex.EncodeToString(tempstr)
}

// 大写
func SHA256Encode(data string) string {
	return strings.ToUpper(Sha256Encode(data))
}

// 加密
func MakePasswrod(password string, random string) string {
	return Sha256Encode(password + random)
}

// 解密
func VaildPassword(password string, random string, Hashpassword string) bool {
	return Sha256Encode(password+random) == Hashpassword
}
