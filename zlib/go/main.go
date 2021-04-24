package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/chindeo/czlib"
)

func main() {
	// 提取文件内容
	f, err := os.Open("./data.json")
	if err != nil {
		fmt.Printf("os open %v\n", err)
		return
	}

	// 压缩内容
	var b bytes.Buffer
	w, err := czlib.NewWriterLevel(&b, -1)
	if err != nil {
		fmt.Printf("flate new writer %v\n", err)
		return
	}
	defer w.Close()
	_, err = io.Copy(w, f)
	if err != nil {
		fmt.Printf("zlib new reader %v", err)
		return
	}
	w.Flush()

	// 将压缩后内容写入文件
	_, err = writeBytes("./data.zip", b.Bytes())
	if err != nil {
		fmt.Printf("WriteBytes %v\n", err)
		return
	}

	// 解压内容
	r, err := czlib.NewReader(&b)
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

func writeBytes(filePath string, b []byte) (int, error) {
	os.MkdirAll(path.Dir(filePath), os.ModePerm)
	fw, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer fw.Close()
	return fw.Write(b)
}
