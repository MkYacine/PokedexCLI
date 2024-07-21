package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type config struct {
	page  int
	cache *Cache
}

var cfg config

func startRepl() {
	reader := bufio.NewReader(os.Stdin)
	cfg = config{
		page:  0,
		cache: NewCache(5 * time.Minute),
	}
	for {
		fmt.Print("Pokedex > ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		fmt.Println("You entered:", input)
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}
		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
			description: "Scroll map forward",
			callback: func() error {
				return commandMap(&cfg)
			},
		},
		"mapb": {
			name:        "mapb",
			description: "Scroll maps backwards",
			callback: func() error {
				return commandMapb(&cfg)
			},
		},
	}
}
