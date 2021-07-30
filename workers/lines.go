package workers

import (
	"bufio"
	"strings"
)

func CountLines(str string) int {
	lines := 0

	scanner := bufio.NewScanner(strings.NewReader(str))
	for scanner.Scan() {
		lines++
	}

	return lines
}
