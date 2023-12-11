package builtins

import (
	"fmt"
	"os"
)

var (
	String string
)

func LSCommand(args ...string) error {
	files, err := os.ReadDir(String)
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	return nil
}
