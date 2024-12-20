package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type config struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

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
	apiLocationEndpoint := "https://pokeapi.co/api/v2/location"
	res, err := http.Get(apiLocationEndpoint)

	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with statuscode: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}
	var c config
	err = json.Unmarshal(body, &c)
	if err != nil {
		fmt.Println("error: could not unmarshal JSON result")
	}
	for _, item := range c.Results {
		fmt.Println(item.Name)
	}
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
			fmt.Print("Pokedex >")
		}
	}
}
