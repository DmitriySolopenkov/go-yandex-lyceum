package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func clear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	for {
		clear()
		var input string
		fmt.Print("КЛЮЧ?!:")
		fmt.Scanln(&input)
		if input == "exit" || input == "выход" {
			break
		}
	}
	fmt.Println("Вы ввели правильный ключ!")
}
