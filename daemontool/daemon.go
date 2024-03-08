package daemontool

import (
	"errors"
	"fmt"
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

func (daemonTool *DaemonTool) GetWordPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

func (daemonTool *DaemonTool) Run(name string, desc string, fn func()) {
	fmt.Println("==============>")
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
