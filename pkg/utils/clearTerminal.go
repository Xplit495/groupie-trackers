package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" { //If the user is on windows we call the command cls
		cmd = exec.Command("cmd", "/c", "cls") //The command is cmd /c cls because the command cls is a command of cmd
	} else { //Else we call the command clear
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout //We set the output of the command to the output of the terminal
	err := cmd.Run()       //We execute the command
	if err != nil {
		fmt.Println("Erreur lors de l'éxécution de la commande de nettoyage du terminal")
		return
	}
}
