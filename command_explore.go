package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandExplore(area string, cfg *config) error {
	var location areaMap
	Url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", area)
	data, ok := cfg.cache.Get(Url)
	if ok {

		if err := json.Unmarshal(data, &location); err != nil {
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
		if err := json.Unmarshal(body, &location); err != nil {
			log.Fatalf("Error unmarshalling JSON: %v", err)
		}
		cfg.cache.Add(Url, body)
	}
	for _, pkm := range location.PokemonEncounters {
		fmt.Printf("- %v \n", pkm.Pokemon.Name)

	}
	return nil
}

type areaMap struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel int `json:"min_level"`
				MaxLevel int `json:"max_level"`
				Chance   int `json:"chance"`
				Method   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				ConditionValues []any `json:"condition_values"`
				PokemonDetails  any   `json:"pokemon_details"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
