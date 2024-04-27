package db

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func FetchCollection() (*http.Response, error) {
	url := fmt.Sprintf("%s/v1/vector/get", os.Getenv("COLLECTION_URL"))

	fmt.Println(url)

	payload := strings.NewReader("{\"collectionName\":\"image_testsv4\",\"id\":[1,2,3]}")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, fmt.Errorf("creating request failed: %v", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("COLLECTION_TOKEN")))

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request failed: %v", err)
	}

	return res, nil
}

func SearchData(vectors []float64) (*http.Response, error) {
	url := fmt.Sprintf("%s/v1/vector/search", os.Getenv("COLLECTION_URL"))
	payloadString := fmt.Sprintf("{\"collectionName\":\"%s\",\"vector\":[%v]}", os.Getenv("COLLECTION_NAME"), strings.Trim(strings.Join(strings.Fields(fmt.Sprint(vectors)), ","), "[]"))
	payload := strings.NewReader(payloadString)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("COLLECTION_TOKEN")))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	// Optional: Check for HTTP error responses (e.g., status codes 4xx or 5xx)
	if res.StatusCode >= 400 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			res.Body.Close() // Ensure resource release if reading body fails
			log.Printf("Error reading response body: %v", err)
			return res, fmt.Errorf("failed to read response body: %v", err)
		}
		res.Body.Close() // Close the response body after reading
		log.Printf("Error response from server: %s", body)
		return res, fmt.Errorf("server returned error status: %s, body: %s", res.Status, body)
	}

	return res, nil
}

// InsertData sends a request to insert vector data into a collection.
func InsertData(vectors []float64, imageUrl string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v1/vector/insert", os.Getenv("COLLECTION_URL"))
	log.Printf("Inserting data at URL: %s", url)

	// Create payload from vectors
	payloadString := fmt.Sprintf("{\"collectionName\":\"%s\",\"data\":[{\"vector\":[%v], \"url\":\"%s\"}]}", os.Getenv("COLLECTION_NAME"), strings.Trim(strings.Join(strings.Fields(fmt.Sprint(vectors)), ","), "[]"), imageUrl)
	payload := strings.NewReader(payloadString)
	log.Printf("Payload: %s", payloadString)

	// Create the request
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("COLLECTION_TOKEN")))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Send the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer res.Body.Close()

	// Read the response
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return res, fmt.Errorf("failed to read response body: %v", err)
	}
	log.Printf("Response status: %s, Body: %s", res.Status, string(respBody))

	if res.StatusCode >= 400 {
		return res, fmt.Errorf("server returned error status: %s, body: %s", res.Status, string(respBody))
	}

	return res, nil
}

// query data by id
func QueryData(id int) (*http.Response, error) {
	url := fmt.Sprintf("%s/v1/vector/query", os.Getenv("COLLECTION_URL"))
	payloadString := fmt.Sprintf("{\"collectionName\":\"%s\",\"filter\":\"Auto_id in [%d]\"}", os.Getenv("COLLECTION_NAME"), id) // Fix here
	payload := strings.NewReader(payloadString)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("COLLECTION_TOKEN")))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	// Optional: Check for HTTP error responses (e.g., status codes 4xx or 5xx)
	if res.StatusCode >= 400 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			res.Body.Close() // Ensure resource release if reading body fails
			log.Printf("Error reading response body: %v", err)
			return res, fmt.Errorf("failed to read response body: %v", err)
		}
		res.Body.Close() // Close the response body after reading
		log.Printf("Error response from server: %s", body)
		return res, fmt.Errorf("server returned error status: %s, body: %s", res.Status, body)
	}

	return res, nil
}
