package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

//返回一个32位md5加密后的字符串, 全长
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//返回一个16位md5加密后的字符串, 半长
func Get16MD5Encode(data string) string {
	return GetMD5Encode(data)[8:24]
}

//返回一个8位md5加密后的字符串
func Get8MD5Encode(data string) string {
	return GetMD5Encode(data)[8:16]
}
