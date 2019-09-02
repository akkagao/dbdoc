package bootstrap

import (
	"sync"
)

var once sync.Once

/**
* 启动服务入口 */
func Start() {
	once.Do(func() {
		// 启动数据库
		StartDb()
		// 启动web服务
		startWeb()
	})
}
