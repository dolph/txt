package main

import (
	"fmt"
	"os"
	"strings"
	"time"

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

	retries := 0
	for {
		// Send it!
		_, err := client.ApiV2010.CreateMessage(params)
		if err != nil {
			// Log an error message and count the attempt
			fmt.Println(err.Error())
			if retries >= 3 {
				// Give up
				os.Exit(1)
			}
			retries += 1
			time.Sleep(time.Duration(retries) * time.Second)
		} else {
			// Required to break the retry loop
			os.Exit(0)
		}
	}
}
