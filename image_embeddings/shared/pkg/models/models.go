package models

type Embedding struct {
	Description string `json:"description"`
	Url         string `json:"url"`
}

// Response represents the JSON response structure from the database.
type EmbeddingResponse struct {
	Code int    `json:"code"`
	Data []Item `json:"data"`
}

// Item represents a single item within the "data" array in the JSON response.
type Item struct {
	AutoID int64     `json:"Auto_id"`
	URL    string    `json:"url"`
	Vector []float64 `json:"vector"`
}

// searchtring stuct
type SearchResponse struct {
	Code int            `json:"code"`
	Data []SearItemItem `json:"data"`
}

// search item struct
type SearItemItem struct {
	AutoID   int64   `json:"Auto_id"`
	Distance float64 `json:"distance"`
}
