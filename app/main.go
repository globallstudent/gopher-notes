package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command)
		if command == "" {
			continue
		}

		if strings.HasPrefix(command, "exit") {
			parts := strings.Fields(command)
			if len(parts) == 2 {
				exitCode, err := strconv.Atoi(parts[1])
				if err == nil {
					os.Exit(exitCode)
				}
			}
			os.Exit(0)
		}

		if strings.HasPrefix(command, "echo ") {
			fmt.Println(strings.TrimPrefix(command, "echo "))
			continue
		}

		if strings.HasPrefix(command, "type ") {
			parts := strings.Fields(command)
			if len(parts) == 2 {
				builtinCommands := map[string]bool{
					"echo":     true,
					"exit":     true,
					"type":     true,
					"pwd":      true,
					"cd":       true,
					"ls":       true,
					"mkdir":    true,
					"rm":       true,
					"rmdir":    true,
					"touch":    true,
					"cat":      true,
					"clear":    true,
					"help":     true,
					"whoami":   true,
					"hostname": true,
				}
				if builtinCommands[parts[1]] {
					fmt.Println(parts[1] + " is a shell builtin")
				} else {
					fmt.Println(parts[1] + ": not found")
				}
				continue
			}
		}

		if command == "pwd" {
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error getting current directory:", err)
				continue
			}
			fmt.Println(dir)
			continue
		}

		if strings.HasPrefix(command, "cd ") {
			parts := strings.Fields(command)
			if len(parts) == 2 {
				err := os.Chdir(parts[1])
				if err != nil {
					fmt.Fprintln(os.Stderr, "cd:", err)
				}
			}
			continue
		}

		if command == "ls" {
			files, err := ioutil.ReadDir(".")
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error reading directory:", err)
				continue
			}
			for _, file := range files {
				fmt.Println(file.Name())
			}
			continue
		}

		if strings.HasPrefix(command, "mkdir ") {
			parts := strings.Fields(command)
			if len(parts) == 2 {
				err := os.Mkdir(parts[1], 0755)
				if err != nil {
					fmt.Fprintln(os.Stderr, "mkdir:", err)
				}
			}
			continue
		}

		if strings.HasPrefix(command, "rm ") {
			parts := strings.Fields(command)
			if len(parts) == 2 {
				err := os.Remove(parts[1])
				if err != nil {
					fmt.Fprintln(os.Stderr, "rm:", err)
				}
			}
			continue
		}

		if command == "clear" {
			fmt.Print("\033[H\033[2J")
			continue
		}

		if command == "whoami" {
			user, err := exec.Command("whoami").Output()
			if err == nil {
				fmt.Print(string(user))
			}
			continue
		}

		if command == "hostname" {
			host, err := os.Hostname()
			if err == nil {
				fmt.Println(host)
			}
			continue
		}

		if command == "help" {
			fmt.Println("Supported commands: pwd, cd, ls, mkdir, rm, rmdir, touch, cat, echo, exit, clear, whoami, hostname, type")
			continue
		}

		// Execute external commands
		args := strings.Fields(command)
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println(command + ": command not found")
		}
	}
}
