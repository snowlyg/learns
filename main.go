package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/snowlyg/learns/windows"
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
	s, err := windows.NewWindowsService(&program{}, "myservice")
	if err != nil {
		fmt.Printf("new service get error %v \n", err)
	}
	if len(os.Args) < 2 {
		usage("no command specified")
	}
	srvName := strings.ToLower(os.Args[2])
	cmd := strings.ToLower(os.Args[1])
	switch cmd {
	case "start":
		err := windows.ServiceStart(srvName)
		if err != nil {
			fmt.Printf("%v \n", err)
		}
		println("start success")
	case "install":

		if len(os.Args) != 7 {
			usage("no command specified")
		}
		err := windows.ServiceInstall(srvName, os.Args[3], os.Args[4], os.Args[5], os.Args[6])
		if err != nil {
			fmt.Printf("%v \n", err)
		}
		println("install success")
	case "stop":
		err := windows.ServiceStop(srvName)
		if err != nil {
			fmt.Printf("%v \n", err)
		}
		println("stop success")
	case "remove":
		err := windows.ServiceUninstall(srvName)
		if err != nil {
			fmt.Printf("%v \n", err)
		}
		println("remove success")
	case "status":
		println(windows.ServiceStatus(srvName))
	default:
		println("invaild command")
	}
	s.Run(false)
}
