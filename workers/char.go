package workers

import (
	"unicode/utf8"
)

func CountChars(str string) int {
    return utf8.RuneCount([]byte(str))
}
