package gpt

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"shared/pkg/util/gpt/serial"
	"strings"
	"time"
)

func GptVisionRequest(imageUrl string, detailed bool) serial.GPTRequest {
	detailLevel := "low"
	if detailed {
		detailLevel = "high"
	}

	// Create the GPT request with system and user messages
	return serial.GPTRequest{
		Model: "gpt-4-turbo",
		Messages: []serial.Message{
			{
				Role: "user",
				Contents: []serial.Content{
					{
						Type: "text",
						Text: "What’s in this image?",
					},
					{
						Type: "image_url",
						ImageURL: &serial.ImageDetail{
							URL:    imageUrl,
							Detail: detailLevel,
						},
					},
				},
			},
		},
		MaxTokens: 300,
	}
}

func MakeGPTRequest(model, apiKey, endpoint string, request serial.GPTRequest) (*serial.GPTResponse, error) {
	// Marshal the request data into JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request: %v", err)
	}

	// Create a new HTTP client with a timeout
	client := &http.Client{Timeout: time.Second * 30}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(string(jsonData)))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	log.Println("apiKey: ", apiKey)
	// Set necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Perform the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Log the raw response body for debugging purposes
	log.Printf("Raw GPT Response Body: %s", string(body))

	// Unmarshal the JSON data into the GPTResponse struct
	var gptResponse serial.GPTResponse
	err = json.Unmarshal(body, &gptResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return &gptResponse, nil
}
