package grab

import (
	"bytes"
	"emails/credentials"
	"fmt"
	"net/http"
)

func GetEmails() {
	fmt.Println("Building call . . . ")

	getUrl := "https://google.com"
	bodyData := []byte("")

	request, err := http.NewRequest("GET", getUrl, bytes.NewBuffer(bodyData))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Status:", response.Status)
	openAPIKey := credentials.GoDotEnvVariable("OPENAI_API_KEY")
	mailgunAPIKey := credentials.GoDotEnvVariable("MAILGUN_API_KEY")
	fmt.Println(openAPIKey, mailgunAPIKey)
}
