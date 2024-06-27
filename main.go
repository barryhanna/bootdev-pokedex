package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Exiting...")
	defer os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}

func commandMap() error {
	fmt.Println("Command: map")
	return nil
}

func commandMapB() error {
	fmt.Println("Command: mapb")
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get Pokemon locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous Pokemon locations",
			callback:    commandMapB,
		},
	}

	for {
		fmt.Print("pokedex >")
		for scanner.Scan() {
			command := scanner.Text()
			commands[command].callback()
			fmt.Print("pokedex >")
		}
	}
}
