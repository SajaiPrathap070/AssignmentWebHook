package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"webhook/model"
	"webhook/webHook"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
	}

	var inputData map[string]interface{}
	err = json.Unmarshal(body, &inputData)
	if err != nil {
		http.Error(w, "Json unmarshal error", http.StatusBadRequest)
	}

	// Close the body to prevent resource leaks
	defer r.Body.Close()

	processRequest(inputData)

	w.WriteHeader(http.StatusOK)
}

func processRequest(inputData map[string]interface{}) {
	outputData := ConvertInputToRequiredOutput(inputData)

	// Send the data to the worker
	webHook.WorkerChannel <- outputData
}

func ConvertInputToRequiredOutput(inputData map[string]interface{}) model.OutputData {

	var outputData = model.OutputData{}

	outputData.AppId = inputData["id"].(string)
	outputData.UserId = inputData["uid"].(string)
	outputData.MessageId = inputData["mid"].(string)
	outputData.Event = inputData["ev"].(string)
	outputData.EventType = inputData["et"].(string)
	outputData.PageTitle = inputData["t"].(string)
	outputData.PageUrl = inputData["p"].(string)
	outputData.BrowserLang = inputData["l"].(string)
	outputData.ScreenSize = inputData["sc"].(string)

	i := 1
	att := "atrk"
	attributes := make(map[string]model.AttrData)
	for {
		if v, ok := inputData[fmt.Sprintf("%v%v", att, i)]; ok {
			attributes[v.(string)] = model.AttrData{
				Type:  inputData[fmt.Sprintf("atrt%v", i)].(string),
				Value: inputData[fmt.Sprintf("atrv%v", i)].(string),
			}

		} else {
			break
		}
		i++
	}

	j := 1
	uatrk := "uatrk"
	traits := make(map[string]model.AttrData)
	for {
		if v, ok := inputData[fmt.Sprintf("%v%v", uatrk, j)]; ok {
			traits[v.(string)] = model.AttrData{
				Type:  inputData[fmt.Sprintf("uatrk%v", j)].(string),
				Value: inputData[fmt.Sprintf("uatrv%v", j)].(string),
			}

		} else {
			break
		}
		j++
	}

	outputData.Attributes = attributes
	outputData.Traits = traits

	return outputData
}
