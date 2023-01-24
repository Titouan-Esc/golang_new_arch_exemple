package server

import (
	"fmt"
	"net/http"
	"time"
)

var message chan string

func SseHandler(res http.ResponseWriter, req *http.Request) {
	flusher, ok := res.(http.Flusher)
	if !ok {
		http.Error(res, "Error", http.StatusInternalServerError)
		return
	}
	EnableCORS(&res)

	go SendMessage()

	defer func() {
		close(message)
		message = nil
	}()

	for {
		select {
		case word := <-message:
			fmt.Fprintf(res, "time%s\n\n", word)
			flusher.Flush()
		case <-req.Context().Done():
			return
		}
	}
}

func SendMessage() {
	for {
		ouai := time.Now().Location().String()
		message <- ouai
		time.Sleep(3 * time.Second)
	}
}

func EnableCORS(res *http.ResponseWriter) {
	(*res).Header().Set("Access-Control-Allow-Origin", "*")
	(*res).Header().Set("Content-Type", "text/event-stream")
	(*res).Header().Set("Cache-Control", "no-cache")
	(*res).Header().Set("Connection", "keep-alive")
}
