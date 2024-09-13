package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func LogWatchHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("./log.txt")
	if err != nil {
		log.Fatal("File not found: ", err)
	}
	defer file.Close()

	offset, err := readLastNlines(file, 10)
	if err != nil {
		fmt.Println("Error in getting last N line")
		return
	}

	for {
		time.Sleep(1 * time.Second)
		newoffset, err := file.Seek(0, io.SeekEnd)
		if err != nil {
			log.Fatal("Error in Reading File: ", err)
		}
		if newoffset > offset { // to check if file has grow
			_, err := file.Seek(offset, io.SeekStart)
			msg := make([]byte, newoffset-offset)
			_, err = file.Read(msg)
			if err != nil {
				if err != nil {
					log.Fatal("Error in Reading File: ", err)
				}
				return
			}

			if err = conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				return
			}
			offset = newoffset
		}

	}
}

func readLastNlines(file *os.File, n int) (int64, error) {
	var lines int
	n = n + 1
	offset, err := file.Seek(0, io.SeekEnd)

	if err != nil {
		return 0, err
	}

	for {
		if offset == 0 {
			return 0, nil
		}
		offset--

		_, err := file.Seek(offset, io.SeekStart)

		if err != nil {
			return 0, err
		}
		buf := make([]byte, 1)
		_, err = file.Read(buf)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}

		if buf[0] == '\n' {
			lines++
		}
		if lines == n {
			offset++
			return offset, nil
		}

	}

}
