package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"shared/pkg/db"
	"shared/pkg/util"
	"shared/pkg/util/gpt"
	"shared/pkg/util/gpt/serial"
)

// Embed generates text embeddings using OpenAI's API.
func Embed(s string) ([]float64, error) {
	log.Printf("Generating embeddings for text: %s", s)

	payload := EmbedRequest{
		Input: s,
		Model: "text-embedding-3-small",
	}
	body, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling payload: %v", err)
		return nil, fmt.Errorf("failed to marshal payload: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/embeddings", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var result struct {
		Data []struct {
			Embedding []float64 `json:"embedding"`
		} `json:"data"`
	}
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		log.Printf("Error unmarshaling response: %v", err)
		return nil, fmt.Errorf("failed to unmarshal response: %v, response body: %s", err, string(respBody))
	}

	if len(result.Data) == 0 {
		log.Println("No embeddings returned: ", string(respBody))
		return nil, fmt.Errorf("no embeddings returned, response body: %s", string(respBody))
	}

	return result.Data[0].Embedding, nil
}

// embedding request sruct
type EmbedRequest struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

func InsertBasicEmbeddingURL(url string) error {
	// Call GptVisionRequest function to generate a description with low detail
	request := gpt.GptVisionRequest(url, false)
	apiKey := os.Getenv("OPENAI_API_KEY")
	endpoint := "https://api.openai.com/v1/chat/completions"

	var response *serial.GPTResponse // Ensure that ResponseType can handle the response structure
	retry := 0
	maxRetries := 5
	var err error

	for {
		response, err = gpt.MakeGPTRequest("gpt-4-turbo", apiKey, endpoint, request)
		if err == nil {
			break // Exit loop if request is successful
		}
		if retry >= maxRetries {
			return fmt.Errorf("max retries reached: %v", err) // Return after max retries
		}
		retry++
		util.ExponentialBackoff(retry) // Wait before next retry
		log.Printf("Retrying after error: %v", err)
	}

	// Check for non-empty response
	if len(response.Choices) == 0 || response.Choices[0].Message.Content == "" {
		return fmt.Errorf("no content received from GPT response")
	}

	// Use the description to generate vectors
	vectors, err := Embed(response.Choices[0].Message.Content)
	if err != nil {
		return err
	}

	_, err = db.InsertData(vectors, url)
	if err != nil {
		return err
	}

	return nil
}
