package builtins

import (
	"fmt"
)

var (
	History []string
)

func DisplayHistory() {
	for i, cmd := range History {
		fmt.Printf("%d: %s\n", i+1, cmd)
	}
}

func ClearHistory() {
	History = nil
}

func ReverseHistory() {
	for i := len(History) - 1; i >= 0; i-- {
		fmt.Printf("%d: %s\n", i+1, History[i])
	}
}

func LastHistory(n int) {
	if n <= 0 {
		return
	}

	startIndex := len(History) - n
	if startIndex < 0 {
		startIndex = 0
	}

	for i := startIndex; i < len(History); i++ {
		fmt.Printf("%d: %s\n", i+1, History[i])
	}
}

func HistoryCommand(args ...string) error {
	if len(args) == 0 {
		DisplayHistory()
		return nil
	}

	/*switch len(args) {
	case "-c":
		ClearHistory()
	case "-h":
		DisplayHistory()
	case "-r":
		ReverseHistory()
	case "n":
		LastHistory()

	default:
		n, err := strconv.Atoi(args[0])
		if err != nil {
			return ErrInvalidArgCount
		}
		LastHistory(n)
	}
	*/
	return nil

}
