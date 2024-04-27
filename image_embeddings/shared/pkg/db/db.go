package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// FetchCollection retrieves the specified collection items.
func FetchCollection() (*http.Response, error) {
	params := DBRequestParams{
		CollectionName: "image_testsv4",
		// Assuming "id" should be an array of integers you're interested in fetching.
		// If "id" should be something else, adjust the type accordingly.
		Data: map[string][]int{"id": {1, 2, 3}},
	}
	return makeDBRequest("/v1/vector/get", params)
}

// SearchData searches the collection using a vector.
func SearchData(vectors []float64) (*http.Response, error) {
	params := DBRequestParams{
		CollectionName: os.Getenv("COLLECTION_NAME"),
		Vector:         vectors,
	}
	return makeDBRequest("/v1/vector/search", params)
}

// InsertData sends a request to insert vector data into a collection.
func InsertData(vectors []float64, imageUrl string) (*http.Response, error) {
	params := DBRequestParams{
		CollectionName: os.Getenv("COLLECTION_NAME"),
		Data: []map[string]interface{}{
			{
				"vector": vectors,
				"url":    imageUrl,
			},
		},
	}
	return makeDBRequest("/v1/vector/insert", params)
}

// QueryData queries data by Auto_id.
func QueryData(id int) (*http.Response, error) {
	params := DBRequestParams{
		CollectionName: os.Getenv("COLLECTION_NAME"),
		Filter:         fmt.Sprintf("Auto_id in [%d]", id),
	}
	return makeDBRequest("/v1/vector/query", params)
}

// DBRequestParams struct for passing parameters to the helper function.
type DBRequestParams struct {
	CollectionName string      `json:"collectionName"`
	Filter         interface{} `json:"filter,omitempty"`
	Vector         interface{} `json:"vector,omitempty"`
	Data           interface{} `json:"data,omitempty"`
}

// Helper function for making a request to the database.
func makeDBRequest(endpoint string, params DBRequestParams) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", os.Getenv("COLLECTION_URL"), endpoint)

	// Marshal the params into a JSON payload.
	payloadBytes, err := json.Marshal(params)
	if err != nil {
		log.Printf("Error marshalling params: %v", err)
		return nil, fmt.Errorf("failed to marshal params: %v", err)
	}
	payload := bytes.NewReader(payloadBytes)

	// Create the request.
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Add necessary headers.
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("COLLECTION_TOKEN")))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Send the request.
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	// Handle non-200 status codes.
	if res.StatusCode >= 400 {
		body, err := io.ReadAll(res.Body)
		defer res.Body.Close() // Close the body regardless of ReadAll's success.
		if err != nil {
			log.Printf("Error reading response body: %v", err)
			return res, fmt.Errorf("failed to read response body: %v", err)
		}
		log.Printf("Error response from server: %s", body)
		return res, fmt.Errorf("server returned error status: %s, body: %s", res.Status, body)
	}

	return res, nil
}
