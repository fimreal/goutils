package http

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
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
		if err == io.EOF {
			break
		} else if err != nil {
			return result, err
		}
		// 拼接每次buf中读到的数据，到result中，返回
		result += string(buf[:n])
	}
	return
}

/* SimplePost 简单 post 调用，内部做了 header 处理和返回错误判断
用法举例：
url := "http://example.com"
headers := make(map[string]string)
headers["Content-Type"] = "application/json;charset=utf-8"

resp, err := http.SimplePost(url, []byte("这里是 post 的数据"), headers)
*/
func SimplePost(url string, data []byte, headers map[string]string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// 省去判断 err ？
	return body, err
}

func HttpBasicAuth(url string, username string, password string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, password)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}
