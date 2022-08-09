// 创建简单的内存 map，存储临时数据，可以设置过期时间
package memdb

import (
	"time"
)

type Grain struct {
	Index   string
	Content string
	Expire  time.Time
}

// 全局变量作为存储容器
var Bucket map[string]*Grain

func init() {
	Bucket = map[string]*Grain{}
}

// 存储数据
func Store(grain *Grain, aliveTime time.Duration) {
	grain.Expire = grain.Expire.Add(aliveTime)
	Bucket[grain.Index] = grain
}

func SelectByIndex(Index string) string {
	g := Bucket[Index]
	if g.isAlive() {
		return g.Content
	}
	return ""
}

func (g *Grain) isAlive() bool {
	return time.Since(g.Expire) <= 0
}
