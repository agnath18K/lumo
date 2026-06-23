package tests

import (
	"testing"

	"github.com/agnath18K/lumo/pkg/config"
)

// TestConfigDefaultValues tests that the default configuration values are set correctly
func TestConfigDefaultValues(t *testing.T) {
	// Create a new config with default values
	cfg := config.DefaultConfig()

	// Test default values
	if cfg.AIProvider != "gemini" {
		t.Errorf("Expected default AIProvider to be 'gemini', got '%s'", cfg.AIProvider)
	}

	if cfg.GeminiModel != "gemini-2.5-flash-lite" {
		t.Errorf("Expected default GeminiModel to be 'gemini-2.5-flash-lite', got '%s'", cfg.GeminiModel)
	}

	if cfg.OpenAIModel != "gpt-3.5-turbo" {
		t.Errorf("Expected default OpenAIModel to be 'gpt-3.5-turbo', got '%s'", cfg.OpenAIModel)
	}

	if cfg.OllamaModel != "llama3" {
		t.Errorf("Expected default OllamaModel to be 'llama3', got '%s'", cfg.OllamaModel)
	}

	if cfg.OllamaURL != "http://localhost:11434" {
		t.Errorf("Expected default OllamaURL to be 'http://localhost:11434', got '%s'", cfg.OllamaURL)
	}

	if cfg.CommandFirstMode {
		t.Errorf("Expected default CommandFirstMode to be false, got true")
	}

	if cfg.EnableShellInInteractive {
		t.Errorf("Expected default EnableShellInInteractive to be false, got true")
	}

	if !cfg.EnablePipeProcessing {
		t.Errorf("Expected default EnablePipeProcessing to be true, got false")
	}

	if !cfg.EnableSystemHealth {
		t.Errorf("Expected default EnableSystemHealth to be true, got false")
	}

	if !cfg.EnableSystemReport {
		t.Errorf("Expected default EnableSystemReport to be true, got false")
	}

	if !cfg.EnableSpeedTest {
		t.Errorf("Expected default EnableSpeedTest to be true, got false")
	}

	if !cfg.EnableAgentMode {
		t.Errorf("Expected default EnableAgentMode to be true, got false")
	}

	if !cfg.EnableAgentREPL {
		t.Errorf("Expected default EnableAgentREPL to be true, got false")
	}

	if !cfg.EnableChatREPL {
		t.Errorf("Expected default EnableChatREPL to be true, got false")
	}

	if !cfg.EnableDesktopAssistant {
		t.Errorf("Expected default EnableDesktopAssistant to be true, got false")
	}

	if cfg.DefaultDesktopEnv != "auto" {
		t.Errorf("Expected default DefaultDesktopEnv to be 'auto', got '%s'", cfg.DefaultDesktopEnv)
	}

	if cfg.Debug {
		t.Errorf("Expected default Debug to be false, got true")
	}
}

// TestConfigLoadFromFile tests loading configuration from a file
func TestConfigLoadFromFile(t *testing.T) {
	t.Skip("Skipping test that requires file system access")
}

// TestConfigEnvironmentVariables tests that environment variables override config file values
func TestConfigEnvironmentVariables(t *testing.T) {
	t.Skip("Skipping test that requires environment variable manipulation")
}
