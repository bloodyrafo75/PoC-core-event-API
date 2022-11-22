package eventRouter

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"poc-core-event-router-api/internals/models"
	"poc-core-event-router-api/internals/pubsubService"
)

func Start(host string, projectId string, topicName string, PORT string) {

	//Connect to Pubsub service
	//TODO stop app execution if connection is not possible
	pubsubService.GetConnection(host, projectId, topicName)

	//start api-server
	mux := http.NewServeMux()
	requestHandler := http.HandlerFunc(ProcessRequest)
	mux.Handle("/", requestHandler)

	err := http.ListenAndServe(":"+PORT, mux)
	log.Fatal(err)

}

// Main handler for requests
func ProcessRequest(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	event, err := GetPayload(r.Body)
	if err != nil {
		http.Error(w, "Payload error", http.StatusInternalServerError)
		return
	}

	rsp, err := pubsubService.Publish(*event)
	if err != nil {
		http.Error(w, "Publishing error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rsp)

}

// Get event struct from the request body
func GetPayload(body io.ReadCloser) (*models.Event, error) {
	var event models.Event
	bBody, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bBody, &event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
