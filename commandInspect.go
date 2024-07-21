package main

import "fmt"

func commandInspect(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("inspect command requires exactly one argument")
	}
	pokemonName := args[0]

	pokemon, exists := cfg.pokedex[pokemonName]
	if !exists {
		fmt.Println("You have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Name, stat.Value)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t)
	}

	return nil
}
