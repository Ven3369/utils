package utils

import (
	"io"
	"io/ioutil"
	"log"
	"os/exec"
)

// path 工作目录 name 命令名 args 命令携带的参数
func RunCommand(path string, name string, args ...string) ([]byte, error) {
	var stdout io.ReadCloser
	var err error
	cmd := exec.Command(name, args...)
	cmd.Dir = path
	if stdout, err = cmd.StdoutPipe(); err != nil { //获取输出对象，可以从该对象中读取输出结果
		log.Fatal(err)
	}
	defer stdout.Close()                // 保证关闭输出流
	if err := cmd.Start(); err != nil { // 运行命令
		log.Fatal("RunCommand cmd start error:", err)
		return nil, err
	}
	return ioutil.ReadAll(stdout)
}
