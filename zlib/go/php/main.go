package main

import (
	"fmt"
	"io"
	"os"

	"github.com/chindeo/czlib"
)

func main() {

	// 提取文件内容
	f, err := os.Open("../php/data.zip")
	if err != nil {
		fmt.Printf("os open %v\n", err)
		return
	}

	// 解压内容
	r, err := czlib.NewReader(f)
	if err != nil {
		fmt.Printf("NewReader %v\n", err)
		return
	}
	defer r.Close()
	_, err = io.Copy(os.Stdout, r)
	if err != nil {
		fmt.Printf("io copy %v\n", err)
		return
	}
}
