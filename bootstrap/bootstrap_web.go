package bootstrap

import (
	"fmt"
	_ "net/http/pprof"
	"os/exec"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"

	"dbdoc/conf"
	"dbdoc/route"
)

var commands = map[string]string{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

/**
* 启动web服务
 */
func startWeb() {
	engine := gin.New()

	// 设置图标
	// engine.Use(favicon.New("./favicon.ico"))

	// 设置中间件
	route.Middleware(engine)

	// 设置路由
	route.RouterInit(engine)
	port := conf.Config.GetString("sys.port")
	go startBrowser(conf.Config.GetString("sys.homeUrl"))
	engine.Run(":" + port)
}

func startBrowser(url string) {
	if url == "" {
		return
	}
	time.Sleep(time.Microsecond * 100)
	run, ok := commands[runtime.GOOS]
	if !ok {
		fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
		return
	}

	cmd := exec.Command(run, url)
	cmd.Start()
}
