package main

import (
	"fmt"
	"net/http"
	"webhook/handler"
	"webhook/webHook"
)

func main() {
	http.HandleFunc("/", handler.HandleRequest)
	go webHook.StartWorker()

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
