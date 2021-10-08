package net

import (
	"fmt"
	"net/http"
)

func HttpGet(url string) (result string, err error) {
	// 借助 http包的 Get()函数 获取网页数据
	resp, err := http.Get(url)
	if err != nil {
		// 将错误传出
		return
	}
	// 读取结束，关闭resp.Body
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	for {
		// 读取Body内容
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读完！err:", err)
			break
		}
		// 拼接每次buf中读到的数据，到result中，返回
		result += string(buf[:n])
	}
	return
}
