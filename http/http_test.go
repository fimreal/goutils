package http

import "fmt"
import "testing"

func TestHttpBasicAuth(T *testing.T) {
	url := "https://hproxy.epurs.com/allowme"
	username := "i"
	password := "1"
	res, err := HttpBasicAuth(url, username, password)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("res: ", string(res))
}
