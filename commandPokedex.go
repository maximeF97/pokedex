package main

import (
	"fmt"
)

func commandPokedex(pokemon string, cfg *config) error {
	_, ok := cfg.pokedex[pokemon]

	if ok {
		fmt.Print("Your Pokedex:\n")
		for _, pkm := range cfg.pokedex {

			fmt.Printf(" - %v\n", pkm)
		}
	} else {
		fmt.Print("Pokedex is empty\n")
	}
	return nil
}
