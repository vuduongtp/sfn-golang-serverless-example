package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/google/uuid"
)

// Event defines your lambda input/output data structure,
type Event struct {
	Payload string `json:"payload"`
	Status  int    `json:"status"`
}

var (
	stateMachineArn = "your-state-machine-arn"
	defaultRegion   = "ap-southeast-1"
)

func main() {
	sfnSvc := sfn.New(session.New(&aws.Config{Region: aws.String(defaultRegion)}))

	inputStruct := Event{
		Payload: fmt.Sprintf("Start at: %v", time.Now()),
		Status:  123,
	}
	input, _ := json.Marshal(inputStruct)

	result, err := sfnSvc.StartExecution(&sfn.StartExecutionInput{
		Input:           aws.String(string(input)),
		StateMachineArn: aws.String(stateMachineArn),
		Name:            aws.String(fmt.Sprintf("%s-vuduongtp", uuid.New().String())),
		TraceHeader:     aws.String(fmt.Sprintf("StartExecution:%s", stateMachineArn)),
	})

	prettyPrint(result)
	prettyPrint(err)
}

func prettyPrint(i interface{}) {
	fmt.Printf("Type: %s\n", reflect.TypeOf(i))
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
	fmt.Println("==============")
}
