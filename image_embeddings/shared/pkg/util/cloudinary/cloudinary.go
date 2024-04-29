package cloudinaryservice

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// BriefAssetResult represents the expected format of each asset in the response.
type BriefAssetResult struct {
	SecureURL string `json:"secure_url"`
}

// FetchFolder makes a direct HTTP request to Cloudinary to fetch asset folder contents.
func FetchFolder() ([]BriefAssetResult, error) {
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")
	cloudName := os.Getenv("CLOUD_NAME")
	folder := "fotos_Streetphotography"
	maxResults := 100 // Set the desired number of results per request

	// Set up the request URL and authentication
	url := fmt.Sprintf("https://api.cloudinary.com/v1_1/%s/resources/image/upload?prefix=%s&max_results=%d", cloudName, folder, maxResults)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("error creating request: %v", err)
		return nil, err
	}

	// Basic Auth Header
	req.SetBasicAuth(apiKey, apiSecret)

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("error sending request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("HTTP Error Response Body: %s", string(body))
		return nil, fmt.Errorf("received HTTP status %d", resp.StatusCode)
	}

	// Parse the JSON response
	var data struct {
		Resources []BriefAssetResult `json:"resources"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("error decoding response: %v", err)
		return nil, err
	}

	return data.Resources, nil
}
