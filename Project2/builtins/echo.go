package builtins

import (
	"fmt"
	"strings"
)

func EchoCommand(args ...string) {
	fmt.Println(strings.Join(args, " "))
}
