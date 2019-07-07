package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nyaruka/phonenumbers"
)

var Version = "dev"

type errorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type successResponse struct {
	NationalNumber         uint64 `json:"national_number"`
	CountryCode            int32  `json:"country_code"`
	IsPossible             bool   `json:"is_possible"`
	IsValid                bool   `json:"is_valid"`
	InternationalFormatted string `json:"international_formatted"`
	NationalFormatted      string `json:"national_formatted"`
	Version                string `json:"version"`
}

func writeResponse(status int, body interface{}) (events.APIGatewayProxyResponse, error) {
	js, err := json.MarshalIndent(body, "", "    ")
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(js),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func parse(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	phone := request.QueryStringParameters["phone"]

	// required phone number
	if phone == "" {
		return writeResponse(http.StatusBadRequest, errorResponse{"missing body", "missing 'phone' parameter"})
	}

	// optional country code
	country := request.QueryStringParameters["country"]

	metadata, err := phonenumbers.Parse(phone, country)
	if err != nil {
		return writeResponse(http.StatusBadRequest, errorResponse{"error parsing phone", err.Error()})
	}

	return writeResponse(http.StatusOK, successResponse{
		NationalNumber:         *metadata.NationalNumber,
		CountryCode:            *metadata.CountryCode,
		IsPossible:             phonenumbers.IsPossibleNumber(metadata),
		IsValid:                phonenumbers.IsValidNumber(metadata),
		NationalFormatted:      phonenumbers.Format(metadata, phonenumbers.NATIONAL),
		InternationalFormatted: phonenumbers.Format(metadata, phonenumbers.INTERNATIONAL),
		Version:                Version,
	})
}

func main() {
	lambda.Start(parse)
}
