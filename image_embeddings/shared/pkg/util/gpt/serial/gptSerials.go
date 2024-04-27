package serial

import "encoding/json"

// GPTRequest defines the structure for sending requests to the GPT API.
type GPTRequest struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

// Message defines the role and contents of a message in a chat.
type Message struct {
	Role     string    `json:"role"`
	Contents []Content `json:"content"`
}

// Content can be text or image_url.
type Content struct {
	Type     string       `json:"type"`
	Text     string       `json:"text,omitempty"`
	ImageURL *ImageDetail `json:"image_url,omitempty"`
}

// ImageDetail defines the URL and detail level for an image content type.
type ImageDetail struct {
	URL    string `json:"url"`
	Detail string `json:"detail"`
}

type GPTResponse struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	SystemFingerprint string   `json:"system_fingerprint"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
}

type Choice struct {
	Index        int              `json:"index"`
	Message      ResponseMessage  `json:"message"`
	LogProbs     *json.RawMessage `json:"logprobs"`
	FinishReason string           `json:"finish_reason"`
}

type ResponseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// Utility function to help with creating JSON payloads that include dynamic types.
func (c Content) MarshalJSON() ([]byte, error) {
	if c.Type == "text" {
		return json.Marshal(struct {
			Type string `json:"type"`
			Text string `json:"text"`
		}{
			Type: c.Type,
			Text: c.Text,
		})
	} else if c.Type == "image_url" {
		return json.Marshal(struct {
			Type     string       `json:"type"`
			ImageURL *ImageDetail `json:"image_url"`
		}{
			Type:     c.Type,
			ImageURL: c.ImageURL,
		})
	}
	return nil, nil // Consider proper error handling or default behavior
}
