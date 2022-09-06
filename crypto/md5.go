package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// 返回一个32位md5加密后的字符串, 全长
func MD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
	// return fmt.Sprintf("%x", h.Sum(nil))
}

// 返回一个16位md5加密后的字符串, 半长
func MD5Encode16(data string) string {
	return MD5Encode(data)[8:24]
}

// 返回一个8位md5加密后的字符串
func MD5Encode8(data string) string {
	return MD5Encode(data)[8:16]
}
