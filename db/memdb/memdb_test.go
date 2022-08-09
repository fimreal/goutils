package memdb

import (
	"testing"
	"time"
)

func TestStoreAndSelect(t *testing.T) {
	sth1 := &Grain{
		Index:   "first",
		Content: "sth1",
		Expire:  time.Now(),
	}
	Store(sth1, 2*time.Second) // 60s 过期
	t.Log(SelectByIndex("first"))
	time.Sleep(2e9)
	t.Log(SelectByIndex("first")) // return ""
	if SelectByIndex("first") != "" {
		t.Error()
	}
}
