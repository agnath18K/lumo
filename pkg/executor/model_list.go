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

  Any model string is accepted; these are suggestions.
  Newer Gemini releases (e.g. 3.x) may also be available.

╰──────────────────────────────────────────────────────────╯
`
	case "claude":
		output = `
╭─────────────── 🐦 Available Claude Models ───────────────╮

  • claude-opus-4-8        (Most capable Opus tier)
  • claude-sonnet-4-6      (Best speed/intelligence balance)
  • claude-haiku-4-5       (Fastest, most cost-effective)
  • claude-fable-5         (Most capable; premium pricing)

  Current model: ` + e.config.ClaudeModel + `

  Any model string is accepted; these are suggestions.

╰──────────────────────────────────────────────────────────╯
`
	case "openai-compatible":
		output = `
╭───────── 🐦 OpenAI-Compatible Provider ──────────╮

  Endpoint: ` + e.config.CompatibleBaseURL + `
  Current model: ` + e.config.CompatibleModel + `

  Any model the endpoint serves is accepted, e.g.:
    • xAI:        grok-*
    • DeepSeek:   deepseek-*
    • Mistral:    mistral-*-latest, codestral-latest
    • Groq:       llama-*, gpt-oss-*
    • OpenRouter: <vendor>/<model>

  Set the endpoint with: config:compatible set-url <url>

╰──────────────────────────────────────────────────╯
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

  • gpt-4o                 (Advanced capabilities)
  • gpt-4o-mini            (Fast, cost-effective — default)
  • gpt-4.1                (Strong coding & instruction following)

  Current model: ` + e.config.OpenAIModel + `

  Any model string is accepted; these are suggestions.
  Newer GPT releases may also be available.

╰──────────────────────────────────────────────────────────╯
`
	}

	return &Result{
		Output:     output,
		IsError:    false,
		CommandRun: cmd.RawInput,
	}, nil
}
