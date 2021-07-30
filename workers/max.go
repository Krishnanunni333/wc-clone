package workers

import (
	"bufio"
	"strings"
)

func MaxLine(str string) int {
	max := 0

	scanner := bufio.NewScanner(strings.NewReader(str))
	for scanner.Scan() {
        temp := CountChars(scanner.Text())
        if temp > max{
            max = temp
        }
	}
	return max
}
