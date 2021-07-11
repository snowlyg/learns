package dir

import (
	"testing"

	"fmt"
)

func TestPath(t *testing.T) {
	t.Run("cwd path", func(t *testing.T) {
		fmt.Printf("cwd path is %s\n", GetCWD())
		fmt.Printf("exec path is %s\n", GetExec())
		path, line := GetCaller()
		fmt.Printf("caller filepath is %s\ncaller line is %d\n", path, line)
	})
}
