package utils

import (
	"os/exec"
	"runtime"
)

// OpenBrowser opens the default web browser to the specified URL.
func OpenBrowser(url string) error {
	var cmd string
	var args []string

	// Determine the appropriate command and arguments based on the operating system
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	args = append(args, url)

	// Execute the command with the specified arguments
	return exec.Command(cmd, args...).Start()
}
