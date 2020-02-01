package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const prompt = "$ "

var (
	noPathError = errors.New("path required.")
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)

	for scanner.Scan() {
		fmt.Print(prompt)

		input := scanner.Text()
		if input == "" {
			continue
		}

		args := strings.Split(input, " ")
		command := args[0]
		args = args[1:]
		switch command {
		case "cd":
			if len(args) == 0 {
				panic(noPathError)
			}
			os.Chdir(args[0])
		case "exit":
			os.Exit(0)
		default:
			cmd := exec.Command(command, args...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout

			cmd.Run()
		}
	}

}
