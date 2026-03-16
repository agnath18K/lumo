package executor

import (
	"fmt"

	"github.com/agnath18K/lumo/pkg/ai"
	"github.com/agnath18K/lumo/pkg/nlp"
)

// handleModelList handles the model list command
func (e *Executor) handleModelList(cmd *nlp.Command) (*Result, error) {
	var output string
	switch e.config.AIProvider {
	case "gemini":
		output = `
╭─────────────── 🐦 Available Gemini Models ───────────────╮

  • gemini-2.5-pro         (Most capable, complex tasks)
  • gemini-2.5-flash       (Fast, great balance of speed/quality)
  • gemini-2.5-flash-lite  (Lightweight, fastest in 2.5 series)

  Current model: ` + e.config.GeminiModel + `

╰──────────────────────────────────────────────────────────╯
`
	case "ollama":
		// Try to get the list of models from Ollama
		ollamaClient := ai.NewOllamaClient(e.config.OllamaURL, e.config.OllamaModel)
		models, err := ollamaClient.ListModels()

		if err != nil {
			return &Result{
				Output:     fmt.Sprintf("Error getting models from Ollama: %v", err),
				IsError:    true,
				CommandRun: cmd.RawInput,
			}, nil
		}

		// Build the model list
		modelList := ""
		for _, model := range models {
			modelList += "  • " + model + "\n"
		}

		if modelList == "" {
			modelList = "  No models found. Use 'ollama pull <model>' to download models.\n"
		}

		output = fmt.Sprintf(`
╭─────────────── 🐦 Available Ollama Models ───────────────╮

%s
  Current model: %s

  Note: To download more models, use 'ollama pull <model>'
  Example: ollama pull llama3

╰──────────────────────────────────────────────────────────╯
`, modelList, e.config.OllamaModel)

	default: // OpenAI
		output = `
╭─────────────── 🐦 Available OpenAI Models ───────────────╮

  • gpt-3.5-turbo          (Fast, cost-effective)
  • gpt-4o                 (Advanced capabilities, slower)
  • gpt-4o-mini            (Balanced performance and quality)

  Current model: ` + e.config.OpenAIModel + `

╰──────────────────────────────────────────────────────────╯
`
	}

	return &Result{
		Output:     output,
		IsError:    false,
		CommandRun: cmd.RawInput,
	}, nil
}
