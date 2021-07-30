package workers

import (
	"os"
    "io/ioutil"
)

func OpenFile(fileName string) string {
	file, err := os.Open(fileName)
	ErrorCheck(err)

    str, err := ioutil.ReadFile(file.Name())
	ErrorCheck(err)

	return string(str)
}
