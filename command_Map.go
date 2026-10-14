package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMap(_ string, cfg *config) error {
	var pkm pokeMap
	var Url string
	Url = "https://pokeapi.co/api/v2/location-area"
	if cfg.nextURL != nil {
		Url = *cfg.nextURL
	}
	data, ok := cfg.cache.Get(Url)
	if ok {

		if err := json.Unmarshal(data, &pkm); err != nil {
			log.Fatalf("Error unmarshalling JSON: %v", err)
		}
	} else {

		res, err := http.Get(Url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		if err := json.Unmarshal(body, &pkm); err != nil {
			log.Fatalf("Error unmarshalling JSON: %v", err)
		}
		cfg.cache.Add(Url, body)
	}
	cfg.nextURL = pkm.Next
	cfg.prevURL = pkm.Previous

	for _, location := range pkm.Results {
		fmt.Printf("%v \n", location.Name)

	}

	return nil

}

type pokeMap struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
