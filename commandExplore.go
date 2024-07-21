package main

import (
	"encoding/json"
	"fmt"
)

func commandExplore(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("explore command requires exactly one argument")
	}
	areaName := args[0]

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", areaName)
	body, err := sendRequest(url)
	if err != nil {
		return err
	}

	var locationArea struct {
		PokemonEncounters []struct {
			Pokemon struct {
				Name string `json:"name"`
			} `json:"pokemon"`
		} `json:"pokemon_encounters"`
	}

	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
