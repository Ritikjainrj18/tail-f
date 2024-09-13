package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
		if _, err = f.WriteString(fmt.Sprintln("test", i)); err != nil {
			panic(err)
		}
	}

}
