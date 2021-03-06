package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"sort"
)

type Data struct {
	Quantity int `json:"quantity"`
	Packs []int `json:"packs"`
}

type Response struct {
	Packs []int
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	data := &Data{}
	headers := map[string]string{
		"Access-Control-Allow-Origin": "*",
		"Access-Control-Allow-Headers" : "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		"Access-Control-Allow-Methods": "GET, POST, OPTIONS",
	}

	_ = json.Unmarshal([]byte(request.Body), data)

	// default quantity is 0
	if data.Quantity < 1 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body: "Quantity must be set or more than 0",
			Headers: headers,
		}, nil
	} else if data.Packs == nil || len(data.Packs) == 0 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body: "No packs available!",
			Headers: headers,
		}, nil
	}

	// sorting available packs as its needed for the calculation later on
	sort.Slice(data.Packs, func(i, j int) bool {
		return data.Packs[i] > data.Packs[j]
	})

	var res []int
	result := calculatePacksNeeded(data.Quantity, data.Packs, res)

	response := Response{Packs: result}
	var jsonData []byte

	jsonData, _ = json.Marshal(response)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body: string(jsonData),
		Headers: headers,
	}, nil
}

func main() {
	lambda.Start(Handler)
}