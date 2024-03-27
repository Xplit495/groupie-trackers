package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// ClearTerminal clears the terminal screen.
func ClearTerminal() {
	// Declare a variable to hold the command to clear the terminal
	var cmd *exec.Cmd

	// Set the command to execute for clearing the terminal
	cmd = exec.Command("cmd", "/c", "cls")

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
