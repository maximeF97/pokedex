package main

import (
	"fmt"
)

func commandInspect(pokemon string, cfg *config) error {
	pkm, ok := cfg.pokedex[pokemon]
	if ok {
		fmt.Printf("Name: %v\n", pkm.Name)
		fmt.Printf("Height: %v\n", pkm.Height)
		fmt.Printf("Weight: %v\n", pkm.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pkm.Stats {
			fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, types := range pkm.Types {
			fmt.Printf("  -%v\n", types.Type.Name)
		}

	} else {
		fmt.Print("you have not caught that pokemon\n")
	}
	return nil
}
