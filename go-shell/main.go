package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var availableCommands = []string{
	"exit", "pwd", "cd", "mkdir", "ls", "touch", "rm", "echo", "cat", "clear",
}

func main() {
	printAvailableCommands()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		displayPrompt()
		userInput, _ := inputReader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		if userInput == "" {
			continue
		}
		handleCommand(userInput)
	}
}

func printAvailableCommands() {
	fmt.Println("Welcome to Go Shell! Available commands:")
	for _, command := range availableCommands {
		fmt.Println("-", command)
	}
	fmt.Println("Type a command to get started.")
}

func displayPrompt() {
	fmt.Print("$ ")
}

func handleCommand(userInput string) {
	commandParts := strings.Fields(userInput)
	if len(commandParts) == 0 {
		return
	}
	command := commandParts[0]
	switch command {
	case "exit":
		exitShell()
	case "pwd":
		displayCurrentDirectory()
	case "cd":
		changeDirectory(commandParts)
	case "mkdir":
		createDirectory(commandParts)
	case "ls":
		listFiles()
	case "touch":
		createFile(commandParts)
	case "rm":
		removeFile(commandParts)
	case "echo":
		displayMessage(commandParts)
	case "cat":
		readFile(commandParts)
	case "clear":
		clearTerminal()
	default:
		executeExternalCommand(commandParts)
	}
}

func exitShell() {
	os.Exit(0)
}

func displayCurrentDirectory() {
	currentDirectory, _ := os.Getwd()
	fmt.Println(currentDirectory)
}

func changeDirectory(commandParts []string) {
	if len(commandParts) < 2 {
		fmt.Println("cd: missing operand")
		return
	}
	if err := os.Chdir(commandParts[1]); err != nil {
		fmt.Println("cd:", err)
	}
}

func createDirectory(commandParts []string) {
	if len(commandParts) < 2 {
		fmt.Println("mkdir: missing operand")
		return
	}
	for _, directoryName := range commandParts[1:] {
		if err := os.Mkdir(directoryName, 0755); err != nil {
			fmt.Println("mkdir:", err)
		}
	}
}

func listFiles() {
	files, _ := os.ReadDir(".")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func createFile(commandParts []string) {
	if len(commandParts) < 2 {
		fmt.Println("touch: missing operand")
		return
	}
	for _, fileName := range commandParts[1:] {
		file, _ := os.Create(fileName)
		file.Close()
	}
}

func removeFile(commandParts []string) {
	if len(commandParts) < 2 {
		fmt.Println("rm: missing operand")
		return
	}
	for _, fileName := range commandParts[1:] {
		os.Remove(fileName)
	}
}

func displayMessage(commandParts []string) {
	fmt.Println(strings.Join(commandParts[1:], " "))
}

func readFile(commandParts []string) {
	if len(commandParts) < 2 {
		fmt.Println("cat: missing operand")
		return
	}
	for _, fileName := range commandParts[1:] {
		data, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("cat:", err)
			continue
		}
		fmt.Print(string(data))
	}
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func executeExternalCommand(commandParts []string) {
	externalCmd := exec.Command(commandParts[0], commandParts[1:]...)
	externalCmd.Stdout = os.Stdout
	externalCmd.Stderr = os.Stderr
	externalCmd.Stdin = os.Stdin
	err := externalCmd.Run()
	if err != nil {
		fmt.Println("command not found or failed:", err)
	}
}
