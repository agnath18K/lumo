package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// OpenAIClient implements the Client interface for OpenAI's API and any
// OpenAI-compatible endpoint (xAI, DeepSeek, Mistral, Groq, OpenRouter, ...).
type OpenAIClient struct {
	apiKey  string
	model   string
	baseURL string
	client  *http.Client
}

// OpenAIRequest represents a request to the OpenAI API
type OpenAIRequest struct {
	Model       string          `json:"model"`
	Messages    []OpenAIMessage `json:"messages"`
	Temperature float64         `json:"temperature"`
}

// OpenAIMessage represents a message in an OpenAI request
type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAIResponse represents a response from the OpenAI API
type OpenAIResponse struct {
	Choices []OpenAIChoice `json:"choices"`
	Error   *OpenAIError   `json:"error,omitempty"`
}

// OpenAIChoice represents a choice in an OpenAI response
type OpenAIChoice struct {
	Message OpenAIMessage `json:"message"`
}

// OpenAIError represents an error from the OpenAI API
type OpenAIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// NewOpenAIClient creates a new OpenAI client
func NewOpenAIClient(apiKey string, model string) *OpenAIClient {
	// If model is empty, use a default model
	if model == "" {
		model = "gpt-4o-mini"
	}

	return &OpenAIClient{
		apiKey:  apiKey,
		model:   model,
		baseURL: "https://api.openai.com/v1",
		client:  &http.Client{},
	}
}

// NewOpenAICompatibleClient creates a client for any OpenAI-compatible API by
// pointing at a custom base URL (e.g. https://api.x.ai/v1). The request and
// response shapes are identical to OpenAI's /chat/completions endpoint.
func NewOpenAICompatibleClient(apiKey string, model string, baseURL string) *OpenAIClient {
	if baseURL == "" {
		baseURL = "https://api.openai.com/v1"
	}
	// Normalize: callers may include a trailing slash.
	baseURL = strings.TrimRight(baseURL, "/")

	return &OpenAIClient{
		apiKey:  apiKey,
		model:   model,
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

// chatCompletionsURL returns the endpoint for chat completion requests.
func (c *OpenAIClient) chatCompletionsURL() string {
	base := c.baseURL
	if base == "" {
		base = "https://api.openai.com/v1"
	}
	return strings.TrimRight(base, "/") + "/chat/completions"
}

// Query sends a query to the OpenAI API and returns the response
func (c *OpenAIClient) Query(query string) (string, error) {
	// Get current working directory for better context
	pwd, err := os.Getwd()
	if err != nil {
		pwd = "unknown" // Fallback if we can't get the current directory
	}

	// Create request body with enhanced system instructions including pwd
	reqBody := OpenAIRequest{
		Model: c.model,
		Messages: []OpenAIMessage{
			{
				Role: "system",
				Content: fmt.Sprintf("You are Lumo, an AI assistant in the terminal. Be concise and helpful.\n\n%s\n\n%s\n\nCurrent Working Directory: %s",
					SystemInstructions, EnvContext(), pwd),
			},
			{
				Role:    "user",
				Content: query,
			},
		},
		Temperature: 0.7,
	}

	// Marshal request to JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", c.chatCompletionsURL(), bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	// Send request
	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	// Parse response
	var openaiResp OpenAIResponse
	if err := json.Unmarshal(body, &openaiResp); err != nil {
		return "", fmt.Errorf("error parsing response: %w", err)
	}

	// Check for API error
	if openaiResp.Error != nil {
		return "", fmt.Errorf("API error: %s", openaiResp.Error.Message)
	}

	// Check for empty response
	if len(openaiResp.Choices) == 0 {
		return "", fmt.Errorf("empty response from API")
	}

	// Return the content from the first choice
	return openaiResp.Choices[0].Message.Content, nil
}

// QueryChat sends a chat query to the OpenAI API with conversation history
func (c *OpenAIClient) QueryChat(messages []OpenAIMessage) (string, error) {
	// Create request body
	reqBody := OpenAIRequest{
		Model:       c.model,
		Messages:    messages,
		Temperature: 0.7,
	}

	// Marshal request to JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", c.chatCompletionsURL(), bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	// Send request
	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	// Parse response
	var openaiResp OpenAIResponse
	if err := json.Unmarshal(body, &openaiResp); err != nil {
		return "", fmt.Errorf("error parsing response: %w", err)
	}

	// Check for API error
	if openaiResp.Error != nil {
		return "", fmt.Errorf("API error: %s", openaiResp.Error.Message)
	}

	// Check for empty response
	if len(openaiResp.Choices) == 0 {
		return "", fmt.Errorf("empty response from API")
	}

	// Return the content from the first choice
	return openaiResp.Choices[0].Message.Content, nil
}

// GetCompletion sends a prompt to the OpenAI API and returns the completion
func (c *OpenAIClient) GetCompletion(ctx context.Context, prompt string) (string, error) {
	// Create request body
	reqBody := OpenAIRequest{
		Model: c.model,
		Messages: []OpenAIMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Temperature: 0.7,
	}

	// Marshal request to JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, "POST", c.chatCompletionsURL(), bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	// Send request
	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	// Parse response
	var openaiResp OpenAIResponse
	if err := json.Unmarshal(body, &openaiResp); err != nil {
		return "", fmt.Errorf("error parsing response: %w", err)
	}

	// Check for API error
	if openaiResp.Error != nil {
		return "", fmt.Errorf("API error: %s", openaiResp.Error.Message)
	}

	// Check for empty response
	if len(openaiResp.Choices) == 0 {
		return "", fmt.Errorf("empty response from API")
	}

	// Return the content from the first choice
	return openaiResp.Choices[0].Message.Content, nil
}

// ProcessChatMessage processes a chat message with conversation history
// and returns the AI response
func (c *OpenAIClient) ProcessChatMessage(ctx context.Context, conversation string) (string, error) {
	// Parse the conversation string into OpenAI messages
	var messages []OpenAIMessage

	// Split the conversation by role
	lines := strings.Split(conversation, "\n")
	var currentRole string
	var currentContent strings.Builder

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Check if this is a role marker
		if strings.HasPrefix(line, "system:") {
			// Save the previous message if any
			if currentRole != "" && currentContent.Len() > 0 {
				messages = append(messages, OpenAIMessage{
					Role:    currentRole,
					Content: currentContent.String(),
				})
				currentContent.Reset()
			}
			currentRole = "system"
			currentContent.WriteString(strings.TrimSpace(line[7:]))
		} else if strings.HasPrefix(line, "user:") {
			// Save the previous message if any
			if currentRole != "" && currentContent.Len() > 0 {
				messages = append(messages, OpenAIMessage{
					Role:    currentRole,
					Content: currentContent.String(),
				})
				currentContent.Reset()
			}
			currentRole = "user"
			currentContent.WriteString(strings.TrimSpace(line[5:]))
		} else if strings.HasPrefix(line, "assistant:") {
			// Save the previous message if any
			if currentRole != "" && currentContent.Len() > 0 {
				messages = append(messages, OpenAIMessage{
					Role:    currentRole,
					Content: currentContent.String(),
				})
				currentContent.Reset()
			}
			currentRole = "assistant"
			currentContent.WriteString(strings.TrimSpace(line[10:]))
		} else if currentRole != "" {
			// Continue the current message
			currentContent.WriteString(" " + line)
		}
	}

	// Add the last message if any
	if currentRole != "" && currentContent.Len() > 0 {
		messages = append(messages, OpenAIMessage{
			Role:    currentRole,
			Content: currentContent.String(),
		})
	}

	// Create request body
	reqBody := OpenAIRequest{
		Model:       c.model,
		Messages:    messages,
		Temperature: 0.7,
	}

	// Marshal request to JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, "POST", c.chatCompletionsURL(), bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	// Send request
	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	// Parse response
	var openaiResp OpenAIResponse
	if err := json.Unmarshal(body, &openaiResp); err != nil {
		return "", fmt.Errorf("error parsing response: %w", err)
	}

	// Check for API error
	if openaiResp.Error != nil {
		return "", fmt.Errorf("API error: %s", openaiResp.Error.Message)
	}

	// Check for empty response
	if len(openaiResp.Choices) == 0 {
		return "", fmt.Errorf("empty response from API")
	}

	// Return the content from the first choice
	return openaiResp.Choices[0].Message.Content, nil
}
