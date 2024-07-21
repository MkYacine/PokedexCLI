package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
)

type Stat struct {
	Name  string
	Value int
}

type Pokemon struct {
	Name           string
	BaseExperience int
	Height         int
	Weight         int
	Stats          []Stat
	Types          []string
}

func calculateCatchChance(baseExperience int) float64 {

	// Calculate catch chance
	// This will give a range from 0.9 (for baseExperience = 0) to 0.1 (for baseExperience = 255 or higher)
	catchChance := 0.9 - float64(baseExperience)*0.003125

	// Ensure catch chance is between 0.1 and 0.9
	if catchChance < 0.1 {
		catchChance = 0.1
	} else if catchChance > 0.9 {
		catchChance = 0.9
	}

	return catchChance
}

func commandCatch(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("catch command requires exactly one argument")
	}
	pokemonName := args[0]

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)
	body, err := sendRequest(url)
	if err != nil {
		return err
	}

	var pokemonData struct {
		Name           string `json:"name"`
		BaseExperience int    `json:"base_experience"`
		Height         int    `json:"height"`
		Weight         int    `json:"weight"`
		Stats          []struct {
			BaseStat int `json:"base_stat"`
			Stat     struct {
				Name string `json:"name"`
			} `json:"stat"`
		} `json:"stats"`
		Types []struct {
			Type struct {
				Name string `json:"name"`
			} `json:"type"`
		} `json:"types"`
	}

	err = json.Unmarshal(body, &pokemonData)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	catchChance := calculateCatchChance(pokemonData.BaseExperience)
	fmt.Printf("%f catch chance.\n", catchChance)
	if rand.Float64() < catchChance {
		stats := make([]Stat, len(pokemonData.Stats))
		for i, stat := range pokemonData.Stats {
			stats[i] = Stat{
				Name:  stat.Stat.Name,
				Value: stat.BaseStat,
			}
		}

		types := make([]string, len(pokemonData.Types))
		for i, t := range pokemonData.Types {
			types[i] = t.Type.Name
		}
		cfg.pokedex[pokemonName] = Pokemon{
			Name:           pokemonData.Name,
			BaseExperience: pokemonData.BaseExperience,
			Height:         pokemonData.Height,
			Weight:         pokemonData.Weight,
			Stats:          stats,
			Types:          types,
		}
		fmt.Printf("%s was caught!\n", pokemonName)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
