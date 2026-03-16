package chat

import (
	"context"
	"fmt"
	"sync"

	"github.com/agnath18K/lumo/pkg/ai"
)

// Manager handles chat conversations
type Manager struct {
	conversations      map[string]*Conversation
	activeConversation string
	maxConversations   int
	maxMessagesPerConv int
	mu                 sync.Mutex
	aiClient           ai.Client
}

// NewManager creates a new chat manager
func NewManager(aiClient ai.Client, maxConversations, maxMessagesPerConv int) *Manager {
	// Set default values if not specified
	if maxConversations <= 0 {
		maxConversations = 5
	}
	if maxMessagesPerConv <= 0 {
		maxMessagesPerConv = 20
	}

	return &Manager{
		conversations:      make(map[string]*Conversation),
		maxConversations:   maxConversations,
		maxMessagesPerConv: maxMessagesPerConv,
		aiClient:           aiClient,
	}
}

// StartNewConversation starts a new conversation and makes it active
func (m *Manager) StartNewConversation() *Conversation {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Create a new conversation with the chat system instructions
	conv := NewConversation(ai.ChatInstructions+"\n\n"+ai.EnvContext(), m.maxMessagesPerConv)

	// Add the conversation to the map
	m.conversations[conv.ID] = conv

	// Set it as the active conversation
	m.activeConversation = conv.ID

	// Trim conversations if needed
	m.trimConversationsIfNeeded()

	return conv
}

// GetActiveConversation returns the active conversation
// If there is no active conversation, it creates a new one
func (m *Manager) GetActiveConversation() *Conversation {
	m.mu.Lock()
	defer m.mu.Unlock()

	// If there is no active conversation or it doesn't exist, create a new one
	if m.activeConversation == "" || m.conversations[m.activeConversation] == nil {
		conv := NewConversation(ai.ChatInstructions+"\n\n"+ai.EnvContext(), m.maxMessagesPerConv)
		m.conversations[conv.ID] = conv
		m.activeConversation = conv.ID
	}

	return m.conversations[m.activeConversation]
}

// SetActiveConversation sets the active conversation
// Returns false if the conversation doesn't exist
func (m *Manager) SetActiveConversation(id string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.conversations[id]; !exists {
		return false
	}

	m.activeConversation = id
	return true
}

// GetConversation returns a conversation by ID
// Returns nil if the conversation doesn't exist
func (m *Manager) GetConversation(id string) *Conversation {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.conversations[id]
}

// ListConversations returns a list of all conversation IDs
func (m *Manager) ListConversations() []string {
	m.mu.Lock()
	defer m.mu.Unlock()

	var ids []string
	for id := range m.conversations {
		ids = append(ids, id)
	}

	return ids
}

// DeleteConversation deletes a conversation by ID
// If the active conversation is deleted, there will be no active conversation
func (m *Manager) DeleteConversation(id string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.conversations[id]; !exists {
		return false
	}

	delete(m.conversations, id)

	// If the active conversation was deleted, clear the active conversation
	if m.activeConversation == id {
		m.activeConversation = ""
	}

	return true
}

// ClearAllConversations deletes all conversations
func (m *Manager) ClearAllConversations() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.conversations = make(map[string]*Conversation)
	m.activeConversation = ""
}

// ProcessMessage processes a user message in the active conversation
// and returns the AI response
func (m *Manager) ProcessMessage(ctx context.Context, message string) (string, error) {
	// Get the active conversation (creates a new one if needed)
	conv := m.GetActiveConversation()

	// Add the user message to the conversation
	conv.AddUserMessage(message)

	// Create a prompt for the AI based on the conversation history
	prompt := m.createPromptFromConversation(conv)

	// Get response from AI
	response, err := m.aiClient.GetCompletion(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("failed to get AI completion: %w", err)
	}

	// Add the assistant response to the conversation
	conv.AddAssistantMessage(response)

	return response, nil
}

// createPromptFromConversation creates a prompt for the AI based on the conversation history
func (m *Manager) createPromptFromConversation(conv *Conversation) string {
	var prompt string

	// For simplicity, we'll just concatenate all messages with role prefixes
	// In a real implementation, you might want to format this differently based on the AI provider
	for _, msg := range conv.GetMessages() {
		prompt += fmt.Sprintf("%s: %s\n\n", msg.Role, msg.Content)
	}

	return prompt
}

// trimConversationsIfNeeded removes the oldest conversations if the number exceeds the maximum
func (m *Manager) trimConversationsIfNeeded() {
	if len(m.conversations) <= m.maxConversations {
		return
	}

	// Find the oldest conversations to remove
	type convAge struct {
		id  string
		age int64
	}

	var convAges []convAge
	for id, conv := range m.conversations {
		// Skip the active conversation
		if id == m.activeConversation {
			continue
		}

		// Use the timestamp of the first message as the age
		var age int64
		if len(conv.Messages) > 0 {
			age = conv.Messages[0].Timestamp.UnixNano()
		}

		convAges = append(convAges, convAge{id: id, age: age})
	}

	// Sort by age (oldest first)
	for i := 0; i < len(convAges)-1; i++ {
		for j := i + 1; j < len(convAges); j++ {
			if convAges[i].age > convAges[j].age {
				convAges[i], convAges[j] = convAges[j], convAges[i]
			}
		}
	}

	// Remove oldest conversations until we're under the limit
	for i := 0; i < len(convAges) && len(m.conversations) > m.maxConversations; i++ {
		delete(m.conversations, convAges[i].id)
	}
}
