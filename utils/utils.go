package utils

import (
	"log"
	"os"
	"os/exec"
)

// 验证路径是否存在
func CheckPath(path string) bool {
	if path == "" {
		return false
	}
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// 执行命令
//
// 标准输出内容
func ExecCmdWithStdOut(cmd string, argv ...string) error {
	logStr := cmd
	if len(argv) > 0 {
		for _, s := range argv {
			logStr = logStr + " " + s
		}
	}
	log.Println("exec cmd:[" + logStr + "]")
	command := exec.Command(cmd, argv...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}
