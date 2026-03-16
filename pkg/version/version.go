package version

import (
	"fmt"
	"runtime"
)

// Version information
var (
	// Version is the current version of Lumo
	Version = "1.1.0"

	// BuildDate is the date when the binary was built
	BuildDate = "Mar 16 2026"

	// GitCommit is the git commit hash when the binary was built
	GitCommit = "HEAD"

	// GoVersion is the version of Go used to build the binary
	GoVersion = runtime.Version()
)

// GetVersion returns the full version string
func GetVersion() string {
	return fmt.Sprintf("%s (built: %s, commit: %s, %s)",
		Version, BuildDate, GitCommit, GoVersion)
}

// GetShortVersion returns just the version number
func GetShortVersion() string {
	return Version
}

// PrintVersion prints the version information
func PrintVersion() {
	fmt.Printf("Lumo version %s\n", GetVersion())
}
