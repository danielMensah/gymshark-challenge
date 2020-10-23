package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"math"
	"net/http"
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

	var res []int
	result := order(data.Quantity, data.Packs, res)

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

func order (quantity int, packs []int, res []int) []int {
	if contains(quantity, packs) {
		res = append(res, quantity)
	} else {
		a := getNearest(quantity, packs)
		b := quantity - a

		res = append(res, a)

		if b > 0 {
			return order(b, packs, res)
		}
	}

	sumRes := sum(res)

	if contains(sumRes, packs) {
		res = []int{sumRes}
	}

	return res

}

func getNearest(quantity int, packs []int) int {
	var nearest = packs[0]

	for _, pack := range packs {
		a := math.Abs(float64(quantity - pack))
		b := math.Abs(float64(quantity - nearest))

		if a < b {
			nearest = pack
		}
	}

	return nearest
}

func contains(item int, arr []int) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}
	return false
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}