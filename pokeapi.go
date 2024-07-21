package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocationArea struct {
	Name string `json:"name"`
}

func sendRequest(url string) ([]byte, error) {
	if cachedData, ok := cfg.cache.Get(url); ok {
		// fmt.Println("Cache hit!")
		return cachedData, nil
	}
	// fmt.Println("Cache miss")

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode == 404 {
		return nil, fmt.Errorf("no more locations")
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s", body)
	cfg.cache.Add(url, body)
	return body, nil
}

func commandMap(cfg *config) error {

	for i := 0; i < 20; i++ {
		cfg.page += 1
		url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d", cfg.page)
		body, _ := sendRequest(url)

		var locationArea LocationArea
		err := json.Unmarshal(body, &locationArea)
		if err != nil {
			return fmt.Errorf("error parsing JSON: %v", err)
		}

		fmt.Printf("City name: %s\n", locationArea.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.page-40 < 0 {
		return fmt.Errorf("this is the first page, cannot go back further.")
	}
	cfg.page -= 40
	for i := 0; i < 20; i++ {
		cfg.page += 1
		url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d", cfg.page)
		body, _ := sendRequest(url)

		var locationArea LocationArea
		err := json.Unmarshal(body, &locationArea)
		if err != nil {
			return fmt.Errorf("error parsing JSON: %v", err)
		}

		fmt.Printf("City name: %s\n", locationArea.Name)
	}

	return nil
}
