package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"shared/pkg/models"
	"strings"
)

// SearchData searches the collection using a vector.
func SearchData(vectors []float64) (*http.Response, error) {
	params := models.DBRequestParams{
		CollectionName: os.Getenv("COLLECTION_NAME"),
		Vector:         vectors,
	}
	return makeDBRequest("/v1/vector/search", params)
}

// InsertData sends a request to insert vector data into a collection.
func InsertData(vectors []float64, imageUrl string) (*http.Response, error) {
	params := models.DBRequestParams{
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
	params := models.DBRequestParams{
		CollectionName: os.Getenv("COLLECTION_NAME"),
		Filter:         fmt.Sprintf("Auto_id in [%d]", id),
	}
	return makeDBRequest("/v1/vector/query", params)
}

// DeleteData deletes data by Auto_id.
func DeleteData(id int) (*http.Response, error) {
	params := models.DBRequestParams{
		CollectionName: os.Getenv("COLLECTION_NAME"),
		Id:             []int{id},
	}
	return makeDBRequest("/v1/vector/delete", params)
}

// DeleteMultiple deletes data by Auto_id.
func DeleteMultiple(ids []int) (*http.Response, error) {
	params := models.DBRequestParams{
		CollectionName: os.Getenv("COLLECTION_NAME"),
		Id:             ids,
	}
	return makeDBRequest("/v1/vector/delete", params)
}

// DeleteMultipleByUrl deletes data by URL.
func DeleteMultipleByUrl(urls []string) (*http.Response, error) {
	urlsJson := `["` + strings.Join(urls, `", "`) + `"]`
	params := models.DBRequestParams{
		CollectionName: os.Getenv("COLLECTION_NAME"),
		Filter:         "url in " + urlsJson,
	}
	res, err := makeDBRequest("/v1/vector/query", params)
	if err != nil {
		return nil, err
	}

	//fetch auto_id from response
	var data models.EmbeddingResponse
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	log.Println("Data: ", data)

	itemIds := make([]int, 0)
	//delete data by auto_id
	for _, item := range data.Data {
		itemIds = append(itemIds, int(item.AutoID))

	}

	res, err = DeleteMultiple(itemIds)
	if err != nil {
		return nil, err
	}

	return res, nil

}

// Helper function for making a request to the database.
func makeDBRequest(endpoint string, params models.DBRequestParams) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", os.Getenv("COLLECTION_URL"), endpoint)

	// Marshal the params into a JSON payload.
	payloadBytes, err := json.Marshal(params)
	if err != nil {
		log.Printf("Error marshalling params: %v", err)
		return nil, fmt.Errorf("failed to marshal params: %v", err)
	}
	payload := bytes.NewReader(payloadBytes)

	log.Printf("Sending request to %s with payload: %s", url, payloadBytes)

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
