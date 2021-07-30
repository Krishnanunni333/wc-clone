package workers

import (
	"log"
)

func ErrorCheck(er error){
	if er != nil {
        log.Fatal(er)
    }
}
