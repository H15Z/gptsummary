package chatgpt

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/H15Z/gptsummary/domain/models"
	"github.com/spf13/viper"
)

const apiUrl = "https://api.openai.com/v1/chat/completions"

type GPT struct {
}

func NewGPT() *GPT {
	return &GPT{}
}

func (g GPT) QueryGPT(prompt string) models.Response {
	var apiKey = viper.GetString("chatgptkey")

	// TODO handle errors better
	requestData, _ := createQueryPalyoad(prompt)

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestData))
	if err != nil {
		log.Println("Error creating the request:", err)
		return models.Response{}
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 120 * time.Second,
	}

	response, err := client.Do(req)
	if err != nil {
		log.Println("Error sending the request:", err)
		return models.Response{}
	}

	defer response.Body.Close()

	response_body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading the response:", err)
		return models.Response{}
	}

	// log.Println("models.Response:", string(response_body)) //DEBUG

	return processResponse(response_body)
}

func processResponse(response_body []byte) models.Response {

	var response models.Response

	err := json.Unmarshal([]byte(response_body), &response)
	if err != nil {
		log.Println("Error:", err)
	}

	return response
}

func createQueryPalyoad(prompt string) ([]byte, error) {
	requestData := models.RequestData{
		Model: "gpt-3.5-turbo",
		Messages: []models.Message{
			{
				Role:    "system",
				Content: "You are a helpful assistant.",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	jsonData, err := json.Marshal(requestData)

	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return []byte{}, err
	}

	return jsonData, nil
}
