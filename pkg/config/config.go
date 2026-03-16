package config

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config holds the application configuration
type Config struct {
	// AI provider settings
	AIProvider   string `json:"ai_provider"`
	GeminiAPIKey string `json:"gemini_api_key"`
	GeminiModel  string `json:"gemini_model"`
	OpenAIAPIKey string `json:"openai_api_key"`
	OpenAIModel  string `json:"openai_model"`
	OllamaURL    string `json:"ollama_url"`
	OllamaModel  string `json:"ollama_model"`

	// Terminal settings
	MaxHistorySize           int  `json:"max_history_size"`
	EnableLogging            bool `json:"enable_logging"`
	EnableShellInInteractive bool `json:"enable_shell_in_interactive"`
	CommandFirstMode         bool `json:"command_first_mode"`

	// Agent mode settings
	EnableAgentMode             bool   `json:"enable_agent_mode"`
	EnableAgentREPL             bool   `json:"enable_agent_repl"`
	AgentConfirmBeforeExecution bool   `json:"agent_confirm_before_execution"`
	AgentMaxSteps               int    `json:"agent_max_steps"`
	AgentSafetyLevel            string `json:"agent_safety_level"`

	// Chat settings
	EnableChatREPL bool `json:"enable_chat_repl"`

	// Pipe settings
	EnablePipeProcessing bool `json:"enable_pipe_processing"`

	// System settings
	EnableSystemHealth bool `json:"enable_system_health"`
	EnableSystemReport bool `json:"enable_system_report"`

	// Speed test settings
	EnableSpeedTest  bool `json:"enable_speed_test"`
	SpeedTestTimeout int  `json:"speed_test_timeout"`

	// Desktop assistant settings
	EnableDesktopAssistant bool   `json:"enable_desktop_assistant"`
	DefaultDesktopEnv      string `json:"default_desktop_env"`

	// Server settings
	EnableServer      bool `json:"enable_server"`
	ServerPort        int  `json:"server_port"`
	ServerQuietOutput bool `json:"server_quiet_output"`

	// Authentication settings
	EnableAuth            bool   `json:"enable_auth"`
	JWTSecret             string `json:"jwt_secret"`
	TokenExpirationHours  int    `json:"token_expiration_hours"`
	RefreshExpirationDays int    `json:"refresh_expiration_days"`

	// Application settings
	Debug bool `json:"debug"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		AIProvider:                  "gemini",                 // Default to Gemini
		GeminiAPIKey:                "",                       // Will be loaded from environment
		GeminiModel:                 "gemini-2.5-flash-lite",   // Default Gemini model
		OpenAIAPIKey:                "",                       // Will be loaded from environment
		OpenAIModel:                 "gpt-3.5-turbo",          // Default OpenAI model
		OllamaURL:                   "http://localhost:11434", // Default Ollama URL
		OllamaModel:                 "llama3",                 // Default Ollama model
		MaxHistorySize:              1000,
		EnableLogging:               true,
		EnableShellInInteractive:    false,    // Shell commands disabled in interactive mode by default
		CommandFirstMode:            false,    // Default to AI-first mode (treat input as AI queries by default)
		EnableAgentMode:             true,     // Agent mode enabled by default
		EnableAgentREPL:             true,     // REPL mode enabled by default
		AgentConfirmBeforeExecution: true,     // Confirm before execution by default
		AgentMaxSteps:               10,       // Maximum 10 steps by default
		AgentSafetyLevel:            "medium", // Medium safety level by default
		EnableChatREPL:              true,     // Chat REPL mode enabled by default
		EnablePipeProcessing:        true,     // Pipe processing enabled by default
		EnableSystemHealth:          true,     // System health checks enabled by default
		EnableSystemReport:          true,     // System reports enabled by default
		EnableSpeedTest:             true,     // Speed test feature enabled by default
		SpeedTestTimeout:            30,       // 30 seconds timeout for speed tests
		EnableDesktopAssistant:      true,     // Desktop assistant enabled by default
		DefaultDesktopEnv:           "auto",   // Auto-detect desktop environment by default
		EnableServer:                false,    // REST server disabled by default
		ServerPort:                  7531,     // Default port for the REST server (uncommon port)
		ServerQuietOutput:           true,     // Suppress server log messages by default
		EnableAuth:                  true,     // Authentication enabled by default
		JWTSecret:                   "",       // Will be generated on first run
		TokenExpirationHours:        24,       // 24 hours token expiration
		RefreshExpirationDays:       7,        // 7 days refresh token expiration
		Debug:                       false,
	}
}

// Load loads the configuration from file and environment variables
func Load() (*Config, error) {
	// Start with default config
	cfg := DefaultConfig()

	// Try to load from config file
	if err := cfg.loadFromFile(); err != nil {
		// If file doesn't exist, create it with default values
		if os.IsNotExist(err) {
			if err := cfg.Save(); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: Could not create config file: %v\n", err)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Warning: Could not load config file: %v\n", err)
		}
	}

	// Load API keys from environment variables
	if geminiKey := os.Getenv("GEMINI_API_KEY"); geminiKey != "" {
		cfg.GeminiAPIKey = geminiKey
	}

	if openaiKey := os.Getenv("OPENAI_API_KEY"); openaiKey != "" {
		cfg.OpenAIAPIKey = openaiKey
	}

	// Generate JWT secret if not set
	if cfg.JWTSecret == "" {
		// Generate a random 32-byte secret
		secretBytes := make([]byte, 32)
		_, err := rand.Read(secretBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to generate JWT secret: %w", err)
		}
		cfg.JWTSecret = base64.StdEncoding.EncodeToString(secretBytes)

		// Save the updated configuration
		if err := cfg.Save(); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Could not save JWT secret to config file: %v\n", err)
		}
	}

	return cfg, nil
}

// loadFromFile loads configuration from the config file
func (c *Config) loadFromFile() error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// Check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return err
	}

	// Read file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	// Parse JSON
	return json.Unmarshal(data, c)
}

// Save saves the configuration to file
func (c *Config) Save() error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// Create directory if it doesn't exist
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	// Marshal to JSON
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	// Write to file
	return os.WriteFile(configPath, data, 0644)
}

// getConfigFilePath returns the path to the config file
func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".config", "lumo", "config.json"), nil
}
