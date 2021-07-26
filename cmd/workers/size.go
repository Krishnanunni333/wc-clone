package workers

import (
    "os"
    "fmt"
)

func PrintBytes(name string) int {
    fi, err := os.Stat(name)
    if err != nil {
        fmt.Println("Err ", err)
    }
    return int(fi.Size())
}
