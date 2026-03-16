package ai

import (
	"os"
	"runtime"
	"strings"
	"sync"
)

var (
	envContext     string
	envContextOnce sync.Once
)

// EnvContext returns a formatted string describing the user's OS, architecture,
// and shell. Computed once and cached for the process lifetime.
func EnvContext() string {
	envContextOnce.Do(func() {
		shell := detectShell()
		envContext = "User Environment: OS=" + runtime.GOOS + ", Arch=" + runtime.GOARCH + ", Shell=" + shell
	})
	return envContext
}

func detectShell() string {
	// Try the SHELL environment variable (macOS, Linux, most Unixes)
	shell := os.Getenv("SHELL")
	if shell != "" {
		parts := strings.Split(shell, "/")
		return parts[len(parts)-1]
	}

	// On Windows, check ComSpec
	comspec := os.Getenv("ComSpec")
	if comspec != "" {
		if strings.Contains(strings.ToLower(comspec), "cmd.exe") {
			return "cmd"
		}
		parts := strings.Split(comspec, "\\")
		return parts[len(parts)-1]
	}

	if runtime.GOOS == "windows" {
		return "powershell"
	}

	return "unknown"
}
