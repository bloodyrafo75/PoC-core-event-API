package pubsubService

import (
	"context"
	"encoding/json"
	"poc-core-event-router-api/internals/models"

	"cloud.google.com/go/pubsub"
)

var (
	topicObj *pubsub.Topic
	ctx      context.Context
)

func Publish(event models.Event) (*models.Response, error) {
	//compose attributes
	attributes := getAttributes(event)
	//convert event struct in string
	eventStr, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	var rsp models.Response
	rsp.MessageID, err = send(string(eventStr), attributes)
	if err != nil {
		return nil, err
	}

	return &rsp, nil
}

// Create the connection w. the server, proyect and topic
func GetConnection(host string, projectId string, topicName string) {
	ctx = context.Background()
	_, topicObj = GetPubsubConnectionToTopic(ctx, host, projectId, topicName)
}

func getAttributes(event models.Event) map[string]string {
	attributes := make(map[string]string)
	attributes["op"] = event.Operation
	attributes["ms"] = event.Microservice

	return attributes
}

func send(payload string, attributes map[string]string) (serverID *string, err error) {
	res := topicObj.Publish(ctx, &pubsub.Message{Attributes: attributes, Data: []byte(payload)})
	messageId, err := res.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &messageId, nil
}
