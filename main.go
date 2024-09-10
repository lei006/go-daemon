package main

import (
	"fmt"
	"time"

	"github.com/lei006/go-daemon/daemontool"
	"github.com/lei006/zlog"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {

	daemonTool := daemontool.DefDaemonTool

	ok, _ := daemontool.RunAtBuild()
	if ok {
		daemonTool.Run("test_app12", "desc 111测试333", Run)
	} else {
		Run()
	}
}

func Run() {

	zlog.SetSaveFile("logs.log", true)
	zlog.LogMaxDurationDate = 3
	zlog.ForceConsoleColor()

	for true {

		zlog.Debug("Run")

		path, err := daemontool.GetWordPath()
		if err != nil {
			fmt.Println("GetWordPath error:", err)
			return
		}
		zlog.Warn("path", path)

		time.Sleep(time.Second)
	}

}
