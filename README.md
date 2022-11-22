# POC - Basic Event-router-api to Pubsub

**Installation**
- Enable Pubsub **local** docker: 
```python
cd third_parties 
docker-compose up -d
cd ..
```
- Set values for `configs/.env`

```python
PUBSUB_HOST = "127.0.0.1:8262"
PROJECT_ID = "my-project-id"
TOPIC_NAME = "topic-name"
PORT = "3000"
```

- Execute the applicaton: `go run cmd/main.go`

Using previous configuration values you can check with this curl:

```shell
curl --location --request POST 'http://localhost:3000' \
--header 'Content-Type: application/json' \
--data-raw '{
"op":"hitchhicker.s.guide",
"m":"marvin@paranoid.android",
"id":42,
"e":"d.adams",
"ms":"thanks.for.all.the.fish"
}'
```

The response will be coded with 200 and will be the msgId in the Pubsub.

```json
{
    "MessageID": "20"
}
```


### Notice a 500 response means something went wrong.