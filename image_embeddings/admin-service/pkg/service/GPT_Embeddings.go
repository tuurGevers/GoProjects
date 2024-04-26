package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

	log.Println("Request body: ", string(body))

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

	log.Println("Response body: ", string(respBody))

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

	log.Println("Embedding data found: ", result.Data[0].Embedding)
	return result.Data[0].Embedding, nil
}

// embedding request sruct
type EmbedRequest struct {
	Input string `json:"input"`
	Model string `json:"model"`
}
