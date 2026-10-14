package main

import "github.com/maxime/pokedexcli/internal/pokecache"

type config struct {
	commands map[string]cliCommand
	nextURL  *string
	prevURL  *string
	cache    *pokecache.Cache
}
