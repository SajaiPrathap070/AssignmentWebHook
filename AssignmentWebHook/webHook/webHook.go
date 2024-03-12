package webHook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webhook/model"
)

// Worker

var WorkerChannel = make(chan model.OutputData)

func StartWorker() {
	for {
		select {
		case data := <-WorkerChannel:
			sendToWebhook(data)
		}
	}
}

func sendToWebhook(data model.OutputData) {
	webhookURL := "https://webhook.site/22ac439e-3a59-45fa-9fba-1a736307b685" // Replace with your webhook URL
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending data to webhook:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Data sent to webhook successfully")
}
