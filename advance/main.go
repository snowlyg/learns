package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/snowlyg/learns/advance/windows"
)

func run() {
	// 每隔 5 秒打印当前时间
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	// for 循环阻塞程序主进程
	// ticker.C 通道每隔五秒会发送一个值
	for {
		<-ticker.C
		// 格式化时间
		now := time.Now().Format(time.RFC3339)
		fmt.Printf("当前时间：%s \n", now)
	}
}

func (p *program) Start() error {
	go run()
	return nil
}

func (p *program) Stop() error {
	//stop
	return nil
}

type program struct{}

// usage 获取终端输入参数
func usage(errmsg string) {
	fmt.Fprintf(os.Stderr,
		"%s\n\n"+
			"usage: %s <command> <servicename>\n"+
			"       where <command> is one of\n"+
			"       install, remove, start, stop, status .\n"+
			"       and use install : .\n"+
			"       install <service name> <exec path> <display name> <system name> <password>  \n",
		errmsg, os.Args[0])
	os.Exit(2)
}

func main() {
	// new 一个服务
	s, err := windows.NewWindowsService(&program{}, "myservice")
	if err != nil {
		fmt.Printf("new service get error %v \n", err)
	}
	if len(os.Args) >= 2 {

		srvName := strings.ToLower(os.Args[2])
		cmd := strings.ToLower(os.Args[1])
		switch cmd {
		case "start": // 启动
			err := windows.ServiceStart(srvName)
			if err != nil {
				fmt.Printf("%v \n", err)
			}
			println("start success")
			return
		case "install": //安装
			if len(os.Args) == 7 {
				err := windows.ServiceInstall(srvName, os.Args[3], os.Args[4], os.Args[5], os.Args[6])
				if err != nil {
					fmt.Printf("%v \n", err)
				}
			} else if len(os.Args) == 5 {
				err := windows.ServiceInstall(srvName, os.Args[3], os.Args[4])
				if err != nil {
					fmt.Printf("%v \n", err)
				}
			} else {
				usage("error command specified")
			}

			println("install success")
			return
		case "stop": // 停止
			err := windows.ServiceStop(srvName)
			if err != nil {
				fmt.Printf("%v \n", err)
			}
			println("stop success")
			return
		case "remove": // 卸载
			err := windows.ServiceUninstall(srvName)
			if err != nil {
				fmt.Printf("%v \n", err)
			}
			println("remove success")
			return
		case "status": // 查询服务状态
			status, _ := windows.ServiceStatus(srvName)
			if status == windows.StatusRunning {
				println("运行中")
			} else if status == windows.StatusStopped {
				println("已停止")
			} else if status == windows.StatusUninstall {
				println("未安装")
			} else if status == windows.StatusUninstall {
				println("未安装")
			}
			return
		default:
			println("invaild command")
		}
		switch cmd {
		case "start": // 启动
			err := windows.ServiceStart(srvName)
			if err != nil {
				fmt.Printf("%v \n", err)
			}
			println("start success")
			return
		case "install": //安装
			if len(os.Args) == 7 {
				err := windows.ServiceInstall(srvName, os.Args[3], os.Args[4], os.Args[5], os.Args[6])
				if err != nil {
					fmt.Printf("%v \n", err)
				}
			} else if len(os.Args) == 5 {
				err := windows.ServiceInstall(srvName, os.Args[3], os.Args[4])
				if err != nil {
					fmt.Printf("%v \n", err)
				}
			} else {
				usage("error command specified")
			}

			println("install success")
			return
		case "stop": // 停止
			err := windows.ServiceStop(srvName)
			if err != nil {
				fmt.Printf("%v \n", err)
			}
			println("stop success")
			return
		case "remove": // 卸载
			err := windows.ServiceUninstall(srvName)
			if err != nil {
				fmt.Printf("%v \n", err)
			}
			println("remove success")
			return
		case "status": // 查询服务状态
			status, _ := windows.ServiceStatus(srvName)
			if status == windows.StatusRunning {
				println("运行中")
			} else if status == windows.StatusStopped {
				println("已停止")
			} else if status == windows.StatusUninstall {
				println("未安装")
			} else if status == windows.StatusUnknown {
				println("未知状态")
			}
			return
		default:
			println("invaild command")
		}
	}
	s.Run(false)
}
