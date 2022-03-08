package sorryf

import "fmt"

func Sorryf(format string, a ...interface{}) string {
	return fmt.Sprintf("☹️ %s", fmt.Sprintf(format, a...))
}
