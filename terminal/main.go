package terminal

import "os"

var argLength int = -1

func ArgCount() int {
	if argLength == -1 {
		argLength = len(os.Args) - 1
	}

	return argLength
}

func GetArg(index int) string {
	if ArgCount() > index {
		return os.Args[index+1]
	}

	return ""
}
