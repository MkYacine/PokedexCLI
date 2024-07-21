package main

import "fmt"

func commandPokedex(args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("pokedex command takes no arguement")
	}
	fmt.Println("Your Pokedex:")
	for k := range cfg.pokedex {
		fmt.Printf("- %s\n", k)
	}
	return nil
}
