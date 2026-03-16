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
	"time"
)

// Message represents a chat message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OllamaClient implements the Client interface for Ollama
type OllamaClient struct {
	baseURL string
	model   string
}

// OllamaRequest represents the request structure for Ollama API
type OllamaRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream,omitempty"`
}

// OllamaResponse represents the response structure from Ollama API
type OllamaResponse struct {
	Model         string  `json:"model"`
	CreatedAt     string  `json:"created_at"`
	Message       Message `json:"message"`
	Done          bool    `json:"done"`
	DoneReason    string  `json:"done_reason,omitempty"`
	TotalDuration int64   `json:"total_duration,omitempty"`
	Error         string  `json:"error,omitempty"`
}

// NewOllamaClient creates a new Ollama client
func NewOllamaClient(baseURL, model string) *OllamaClient {
	// Ensure the URL doesn't end with a slash
	baseURL = strings.TrimSuffix(baseURL, "/")

	return &OllamaClient{
		baseURL: baseURL,
		model:   model,
	}
}

// GenerateText generates text using the Ollama API
func (c *OllamaClient) GenerateText(prompt string, systemPrompt string) (string, error) {
	// Create messages array with system prompt and user prompt
	messages := []Message{
		{
			Role:    "system",
			Content: systemPrompt,
		},
		{
			Role:    "user",
			Content: prompt,
		},
	}

	// Create request body
	requestBody := OllamaRequest{
		Model:    c.model,
		Messages: messages,
		Stream:   false, // Explicitly set to false to get a complete response
	}

	// Convert request to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", c.baseURL+"/api/chat", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{
		Timeout: 60 * time.Second, // Set a longer timeout for model responses
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to Ollama: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	// Check for error status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Ollama API error (status %d): %s", resp.StatusCode, string(body))
	}

	// Handle streaming response
	lines := strings.Split(string(body), "\n")
	var fullContent strings.Builder

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		var resp OllamaResponse
		if err := json.Unmarshal([]byte(line), &resp); err == nil {
			fullContent.WriteString(resp.Message.Content)
		}
	}

	result := fullContent.String()

	// Clean up markdown formatting if present
	result = strings.ReplaceAll(result, "```bash", "")
	result = strings.ReplaceAll(result, "```", "")
	result = strings.TrimSpace(result)

	return result, nil
}

// GenerateChat generates a chat response using the Ollama API
func (c *OllamaClient) GenerateChat(messages []Message, systemPrompt string) (string, error) {
	// Prepend system message if provided
	if systemPrompt != "" {
		sysMsg := Message{
			Role:    "system",
			Content: systemPrompt,
		}
		messages = append([]Message{sysMsg}, messages...)
	}

	// Create request body
	requestBody := OllamaRequest{
		Model:    c.model,
		Messages: messages,
		Stream:   false, // Explicitly set to false to get a complete response
	}

	// Convert request to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", c.baseURL+"/api/chat", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{
		Timeout: 60 * time.Second, // Set a longer timeout for model responses
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to Ollama: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	// Check for error status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Ollama API error (status %d): %s", resp.StatusCode, string(body))
	}

	// Handle streaming response
	lines := strings.Split(string(body), "\n")
	var fullContent strings.Builder

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		var resp OllamaResponse
		if err := json.Unmarshal([]byte(line), &resp); err == nil {
			fullContent.WriteString(resp.Message.Content)
		}
	}

	result := fullContent.String()

	// Clean up markdown formatting if present
	result = strings.ReplaceAll(result, "```bash", "")
	result = strings.ReplaceAll(result, "```", "")
	result = strings.TrimSpace(result)

	return result, nil
}

// Query sends a query to the Ollama API and returns the response
func (c *OllamaClient) Query(query string) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		pwd = "unknown"
	}
	systemPrompt := fmt.Sprintf("%s\n\n%s\n\nCurrent Working Directory: %s",
		SystemInstructions, EnvContext(), pwd)
	return c.GenerateText(query, systemPrompt)
}

// GetCompletion sends a prompt to the Ollama API and returns the completion
func (c *OllamaClient) GetCompletion(ctx context.Context, prompt string) (string, error) {
	systemPrompt := fmt.Sprintf("%s\n\n%s", AgentInstructions, EnvContext())
	return c.GenerateText(prompt, systemPrompt)
}

// ListModels returns a list of available models from Ollama
func (c *OllamaClient) ListModels() ([]string, error) {
	// Create HTTP request
	req, err := http.NewRequest("GET", c.baseURL+"/api/tags", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request to Ollama: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	// Check for error status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Ollama API error (status %d): %s", resp.StatusCode, string(body))
	}

	// Parse response
	var response struct {
		Models []struct {
			Name string `json:"name"`
		} `json:"models"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	// Extract model names
	var models []string
	for _, model := range response.Models {
		models = append(models, model.Name)
	}

	return models, nil
}
