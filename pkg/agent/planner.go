package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/agnath18K/lumo/pkg/ai"
	"github.com/agnath18K/lumo/pkg/config"
)

// Planner handles the generation of execution plans
type Planner struct {
	config   *config.Config
	aiClient ai.Client
}

// NewPlanner creates a new planner instance
func NewPlanner(cfg *config.Config, aiClient ai.Client) *Planner {
	return &Planner{
		config:   cfg,
		aiClient: aiClient,
	}
}

// CreatePlan generates a plan for the given task
func (p *Planner) CreatePlan(ctx context.Context, task *Task) (*Plan, error) {
	// Create the prompt for the AI
	prompt := fmt.Sprintf(`
You are Lumo, an AI-powered command-line assistant.
Create a step-by-step plan to accomplish the following task using shell commands:

Task: %s

%s

Provide a detailed plan with the following structure:
1. A brief description of the overall approach
2. A numbered list of shell commands to execute
3. For each command, include:
   - The exact command to run
   - A brief explanation of what the command does
   - Whether the command is critical for the task

IMPORTANT: Your response MUST be a valid JSON object with the following structure:
{
  "description": "Overall approach description",
  "steps": [
    {
      "id": 1,
      "command": "exact shell command",
      "description": "what this command does",
      "isCritical": true/false
    },
    ...
  ]
}

Do not include any text before or after the JSON object. The response must be parseable as JSON.
Do not include markdown formatting, code blocks, or any other non-JSON content.

Ensure all commands are safe to execute and won't cause data loss or system damage.
Use relative paths when possible and avoid commands that require sudo.
Limit the plan to at most %d steps.
`, task.Description, ai.EnvContext(), p.config.AgentMaxSteps)

	// Get response from AI
	response, err := p.aiClient.GetCompletion(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to get AI completion: %w", err)
	}

	// Extract JSON from the response
	jsonStart := -1
	jsonEnd := -1

	// Find the start of the JSON object
	for i := 0; i < len(response); i++ {
		if response[i] == '{' {
			jsonStart = i
			break
		}
	}

	// Find the end of the JSON object
	if jsonStart >= 0 {
		braceCount := 1
		for i := jsonStart + 1; i < len(response); i++ {
			if response[i] == '{' {
				braceCount++
			} else if response[i] == '}' {
				braceCount--
				if braceCount == 0 {
					jsonEnd = i + 1
					break
				}
			}
		}
	}

	// Extract the JSON part
	var jsonData string
	if jsonStart >= 0 && jsonEnd > jsonStart {
		jsonData = response[jsonStart:jsonEnd]
	} else {
		return nil, fmt.Errorf("failed to extract JSON from AI response")
	}

	// Parse the JSON response
	var planData struct {
		Description string `json:"description"`
		Steps       []struct {
			ID          int    `json:"id"`
			Command     string `json:"command"`
			Description string `json:"description"`
			IsCritical  bool   `json:"isCritical"`
		} `json:"steps"`
	}

	if err := json.Unmarshal([]byte(jsonData), &planData); err != nil {
		return nil, fmt.Errorf("failed to parse AI response: %w", err)
	}

	// Create the plan
	plan := &Plan{
		Task:        task,
		Description: planData.Description,
		CreatedAt:   time.Now(),
		Steps:       make([]*Step, len(planData.Steps)),
	}

	// Add steps to the plan
	for i, stepData := range planData.Steps {
		plan.Steps[i] = &Step{
			ID:          stepData.ID,
			Command:     stepData.Command,
			Description: stepData.Description,
			IsCritical:  stepData.IsCritical,
			Executed:    false,
		}
	}

	return plan, nil
}
