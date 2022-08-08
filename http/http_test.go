package http

import (
	"fmt"
	"testing"
)

func TestHttpBasicAuth(T *testing.T) {
	url := "https://epurs.com/health"
	username := "123"
	password := "123"
	res, err := HttpBasicAuth(url, username, password)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("res: ", string(res))
}

func TestHttpGet(t *testing.T) {
	url := "http://baidu.com"
	res, err := HttpGet(url)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
