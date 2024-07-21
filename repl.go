package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type config struct {
	page    int
	cache   *Cache
	pokedex map[string]Pokemon
}

var cfg config

func startRepl() {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	cfg = config{
		page:    0,
		cache:   NewCache(5 * time.Minute),
		pokedex: make(map[string]Pokemon),
	}
	for {
		fmt.Print("Pokedex > ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		args := words[1:]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(args)
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
	callback    func(args []string) error
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
			callback: func(args []string) error {
				return commandMap(&cfg)
			},
		},
		"mapb": {
			name:        "mapb",
			description: "Scroll maps backwards",
			callback: func(args []string) error {
				return commandMapb(&cfg)
			},
		},
		"explore": {
			name:        "explore",
			description: "Explore area passed as arguement",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch pokemon passed as arguement",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught pokemon passed as arguement",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View your pokedex",
			callback:    commandPokedex,
		},
	}
}
