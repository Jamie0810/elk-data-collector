package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda/messages"
	lc "github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/djhworld/go-lambda-invoke/golambdainvoke"
)

type Input struct {
	Port          int
	Payload       interface{}
	ClientContext *lc.ClientContext
	Deadline      *messages.InvokeRequest_Timestamp
}

type AwsLogs struct {
	Data string `json:"data"`
}
type Event struct {
	AwsLogs AwsLogs `json:"awslogs"`
}
type Response struct {
	Message string `json:"message"`
}

func main() {
	awsLogs := AwsLogs{Data: "H4sIAAAAAAAAAO1S22rbQBD9FbMU8mLFe7/ozVAnFNIGbD81DmKlHScCXdyV5MQN+fesV016DYS+Ffq0MHPmnJmz5wHV0HX2BtaHHaAUvZ+v59nHxWo1P1+gKWrvGvChbAiVTHFGFTWhXLU3574ddqEzs3fdrLJ17uwMCucTB/ukaOu6be6gqpKd7Uto+qQD64vbZESOFKveg60DB8WUzLCcUTW7encxXy9W62tOMCswl0aA5kay3PB8mwNRlG+dtSpQdEPeFb7c9WXbnJVVD75D6RVytrdhg6qCom89uo5ai31Y4th+QKULkkxSxrGgRkpBCaZaKU0Il0JLSqXSRnIiNNVEUGyMElRrI7g+Ht+XwbLe1uF6IilXhhCqDebTZysD/cMGVbCHaoPSDfrw6exyg6ah1Bb2uGys3trGVeBTSmPv22xsXUSTJnAPxXDET4Kc78FF4It8hB6tS7BMqFoTnAqaYn0advkcoR34fVmMnH/8h4gKVrksKgRg7wcIte3QFEflrLH1OP/mv42cL/M11K0/ZF35daQJVv0MsH60I7xpiFI6kqRDl4Dt+oSkP0YvfR5L/3IdD1+G4F5WuigqqXYsdyZRuVAJ56RIdM5dghmXIDAmW8CRIB+6sgk/lLm2tuW48XflCLn39pD13hbwzE4SiZ0WOlArpayWOJygHNgQG0UKVjC5QY+bBj1Of02lUpxrZpihghkuJTOSYhbyKQkhWhlKNGGCC0ONYQy/ksow/3oqF8vl5fL1WDL8WyyXwbzSg5ucbG1dVocYjpNJ2U2atp9s26Fxp2+KqDkNi/2P6D8V0evHJ9SBw5ErBgAA"}
	event := Event{
		AwsLogs: awsLogs,
	}

	resp, err := golambdainvoke.Run(Input{
		Port:    8080,
		Payload: event,
	})

	if err != nil {
		fmt.Println("Failed to invoke the lambda: ", err)
	}

	var response Response
	err = json.Unmarshal(resp, &response)
	fmt.Println("Message: ", response.Message)
}
