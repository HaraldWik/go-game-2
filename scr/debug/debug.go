package debug

import "log"

const (
	ERROR = "\033[31m" + "[ERROR]" + "\033[0m"
)

func HandleError(text string, err error) {
	if err != nil {
		log.Fatalf("%s: %s: %v", ERROR, text, err)
	}
}
