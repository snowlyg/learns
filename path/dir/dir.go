package dir

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// 返回相对当前目录之一的根路径
func GetCWD() string {
	cwd, _ := os.Getwd()
	return cwd
}

// 返回当前进程的可执行文件绝对路径
func GetExec() string {
	exePath, _ := os.Executable()
	return filepath.Dir(exePath)
}

// 返回存放临时文件的默认路径
func GetTempDir() string {
	return os.TempDir()
}

// 返回有关调用 goroutine 堆栈上的函数调用的文件和行号信息
func GetCaller() (string, int) {
	_, filename, line, ok := runtime.Caller(0)
	if ok {
		return path.Dir(filename), line
	}
	return "", 0
}
