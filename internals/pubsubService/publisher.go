package pubsubService

import (
	"context"
	"encoding/json"
	"fmt"
	"poc-core-event-router-api/internals/models"
	"strings"

	"cloud.google.com/go/pubsub"
	"github.com/fatih/structs"
)

var (
	topicObj *pubsub.Topic
	ctx      context.Context
)

func Publish(event models.MessageModel) (*models.Response, error) {

	//compose attributes
	attributes := getAttributes(event.Attributes)

	pubsubMsg := models.PubSubMessageModel{
		Payload:         event.Payload,
		SpecificPayload: event.SpecificPayload,
	}

	eventStr, err := json.Marshal(pubsubMsg)
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

func getAttributes(eventAttributes models.MessageAttributes) map[string]string {
	attributes := make(map[string]string)

	keys := structs.Names(eventAttributes)
	values := structs.Values(eventAttributes)

	for i, key := range keys {
		attributes[strings.ToLower(key)] = fmt.Sprintf("%v", values[i])
	}

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
