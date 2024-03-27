package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// ClearTerminal clears the terminal screen.
func ClearTerminal() {
	// Declare a variable to hold the command to clear the terminal
	var cmd *exec.Cmd

	// Check the operating system and set the appropriate command
	if runtime.GOOS == "windows" {
		// Windows command to clear the terminal
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		// Unix/Linux command to clear the terminal
		cmd = exec.Command("clear")
	}

	// Set the standard output of the command to os.Stdout
	cmd.Stdout = os.Stdout

	// Run the command to clear the terminal
	err := cmd.Run()

	// Check if there was an error during the execution of the command
	if err != nil {
		// Print an error message if there was an error
		fmt.Println("Error while executing the terminal clearing command")
		return
	}
}
