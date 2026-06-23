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

// anthropicVersion is the required API version header value for the Messages API.
const anthropicVersion = "2023-06-01"

// claudeDefaultMaxTokens is the default output token budget for a request.
// max_tokens is required by the Anthropic Messages API.
const claudeDefaultMaxTokens = 4096

// ClaudeClient implements the Client interface for Anthropic's Messages API
type ClaudeClient struct {
	apiKey string
	model  string
	client *http.Client
}

// ClaudeMessage represents a single message in a Claude request.
// Roles are limited to "user" and "assistant"; the system prompt is a
// top-level field, not a message.
type ClaudeMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ClaudeRequest represents a request to the Anthropic Messages API
type ClaudeRequest struct {
	Model     string          `json:"model"`
	MaxTokens int             `json:"max_tokens"`
	System    string          `json:"system,omitempty"`
	Messages  []ClaudeMessage `json:"messages"`
}

// ClaudeContentBlock represents a single content block in a Claude response
type ClaudeContentBlock struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// ClaudeError represents an error returned by the Anthropic API
type ClaudeError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// ClaudeResponse represents a response from the Anthropic Messages API
type ClaudeResponse struct {
	Content []ClaudeContentBlock `json:"content"`
	Error   *ClaudeError         `json:"error,omitempty"`
}

// NewClaudeClient creates a new Anthropic Claude client
func NewClaudeClient(apiKey string, model string) *ClaudeClient {
	// If model is empty, use a sensible default (best speed/cost balance)
	if model == "" {
		model = "claude-sonnet-4-6"
	}

	return &ClaudeClient{
		apiKey: apiKey,
		model:  model,
		client: &http.Client{},
	}
}

// send issues a Messages API request and returns the first text block.
func (c *ClaudeClient) send(ctx context.Context, system string, messages []ClaudeMessage) (string, error) {
	reqBody := ClaudeRequest{
		Model:     c.model,
		MaxTokens: claudeDefaultMaxTokens,
		System:    system,
		Messages:  messages,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.apiKey)
	req.Header.Set("anthropic-version", anthropicVersion)

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	var claudeResp ClaudeResponse
	if err := json.Unmarshal(body, &claudeResp); err != nil {
		return "", fmt.Errorf("error parsing response: %w", err)
	}

	if claudeResp.Error != nil {
		return "", fmt.Errorf("API error: %s", claudeResp.Error.Message)
	}

	// Return the first text block from the response
	for _, block := range claudeResp.Content {
		if block.Type == "text" {
			return block.Text, nil
		}
	}

	return "", fmt.Errorf("empty response from API")
}

// systemPrompt builds the shared Lumo system prompt including environment context.
func systemPrompt() string {
	pwd, err := os.Getwd()
	if err != nil {
		pwd = "unknown"
	}
	return fmt.Sprintf("You are Lumo, an AI assistant in the terminal. Be concise and helpful.\n\n%s\n\n%s\n\nCurrent Working Directory: %s",
		SystemInstructions, EnvContext(), pwd)
}

// Query sends a query to the Anthropic API and returns the response
func (c *ClaudeClient) Query(query string) (string, error) {
	return c.send(context.Background(), systemPrompt(), []ClaudeMessage{
		{Role: "user", Content: query},
	})
}

// GetCompletion sends a prompt to the Anthropic API and returns the completion
func (c *ClaudeClient) GetCompletion(ctx context.Context, prompt string) (string, error) {
	return c.send(ctx, "", []ClaudeMessage{
		{Role: "user", Content: prompt},
	})
}

// ProcessChatMessage processes a chat message with conversation history
// and returns the AI response. The conversation is parsed the same way as the
// OpenAI client; "system:" lines are merged into the top-level system prompt
// since Claude does not accept a system role inside the messages array.
func (c *ClaudeClient) ProcessChatMessage(ctx context.Context, conversation string) (string, error) {
	var system strings.Builder
	var messages []ClaudeMessage

	lines := strings.Split(conversation, "\n")
	var currentRole string
	var currentContent strings.Builder

	flush := func() {
		if currentRole == "" || currentContent.Len() == 0 {
			return
		}
		if currentRole == "system" {
			if system.Len() > 0 {
				system.WriteString("\n\n")
			}
			system.WriteString(currentContent.String())
		} else {
			messages = append(messages, ClaudeMessage{Role: currentRole, Content: currentContent.String()})
		}
		currentContent.Reset()
	}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		switch {
		case strings.HasPrefix(line, "system:"):
			flush()
			currentRole = "system"
			currentContent.WriteString(strings.TrimSpace(line[7:]))
		case strings.HasPrefix(line, "user:"):
			flush()
			currentRole = "user"
			currentContent.WriteString(strings.TrimSpace(line[5:]))
		case strings.HasPrefix(line, "assistant:"):
			flush()
			currentRole = "assistant"
			currentContent.WriteString(strings.TrimSpace(line[10:]))
		default:
			if currentRole != "" {
				currentContent.WriteString(" " + line)
			}
		}
	}
	flush()

	return c.send(ctx, system.String(), messages)
}
