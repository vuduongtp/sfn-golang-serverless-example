package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/aws/aws-lambda-go/lambda"
)

// Event defines your lambda input/output data structure,
type Event struct {
	Payload string `json:"payload"`
	Status  int    `json:"status"`
}

// HandleRequest handles the incomming StepFunction request
func HandleRequest(e Event) (Event, error) {
	prettyPrint(e)
	return Event{
		Payload: fmt.Sprintf("%s is handled by Lambda function", e.Payload),
		Status:  e.Status,
	}, nil
}

func prettyPrint(i interface{}) {
	fmt.Printf("Type: %s\n", reflect.TypeOf(i))
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
	fmt.Println("==============")
}

func main() {
	lambda.Start(HandleRequest)
}
