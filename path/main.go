package main

import (
	"fmt"

	"github.com/snowlyg/learns/path/dir"
)

func main() {
	fmt.Printf("cwd path is %s\n", dir.GetCWD())
	fmt.Printf("exec path is %s\n", dir.GetExec())
	path, line := dir.GetCaller()
	fmt.Printf("caller filepath is %s\ncaller line is %d\n", path, line)
}
