package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	client := twilio.NewRestClient()
	from := os.Getenv("TXT_FROM")
	to := os.Getenv("TXT_TO")

	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(strings.Join(os.Args[1:], " "))

	resp, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}
