package daemontool

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"

	"github.com/lei006/go-daemon/daemontool/zcli"
)

type DaemonTool struct {
}

var DefDaemonTool *DaemonTool

func init() {
	DefDaemonTool = &DaemonTool{}
}

func GetWordPath() (string, error) {

	// 取得工作目录
	workPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// 获取当前执行程序的文件信息
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	bret := strings.Contains(exePath, "go-build")
	if bret {
		return strings.Replace(workPath, "\\", "/", -1), nil
	}

	workPath, err = filepath.Abs(filepath.Dir(exePath)) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}

	return strings.Replace(workPath, "\\", "/", -1), nil
}

func (daemonTool *DaemonTool) Run(name string, desc string, fn func()) {
	zcli.LaunchServiceRun(name, desc, fn)
}

// 判断程序是否运行在build之后的程序
func RunAtBuild() (bool, error) {

	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return false, errors.New("ReadBuildInfo Run Error")
	}

	for _, item := range buildInfo.Settings {
		if item.Key == "vcs.time" {
			return true, nil
		}
	}

	return false, nil
}
