package main

import (
	"flag"
	"go-to-cloud/internal/builder"
	"go-to-cloud/internal/repositories/migrations"
	"go-to-cloud/internal/routers"
	"os"
	"strings"
)

var aPort = flag.String("port", "", "端口")

// runMode 获取运行模式
// @string: 端口
func runMode() string {
	// 优先读取命令行参数，其次使用go env，最后使用默认值
	flag.Parse()

	if len(*aPort) == 0 {
		*aPort = os.Getenv("port")
	}

	if len(*aPort) == 0 {
		*aPort = ":80"
	}

	if !strings.HasPrefix(*aPort, ":") {
		*aPort = ":" + *aPort
	}

	return *aPort
}

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @BasePath /api
func main() {
	port := runMode()
	// 迁移数据库
	migrations.AutoMigrate()
	// 启动流水线监测
	builder.PipelinesWatcher()
	// server模式运行
	_ = routers.SetRouters().Run(port)
}
