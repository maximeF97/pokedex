package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	pokecache "github.com/maxime/pokedexcli/internal/pokecache"
)

func main() {
	commandsMap := commands()
	cfg := &config{
		commands: commandsMap,
		cache:    pokecache.NewCache(5 * time.Second),
	}

	startRepl(cfg)
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			return
		}

		input := scanner.Text()
		cleanS := cleanInput(input)
		if len(cleanS) == 0 {
			continue
		}

		commandName := cleanS[0]
		reste := strings.Join(cleanS[1:], " ")

		if len(reste) == 0 {
			fmt.Println("Type something else")
			continue
		}
		command, exists := cfg.commands[commandName]

		if exists {
			err := command.callback(reste, cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue

		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	cleanWorld := []string{}
	words := strings.Fields(text)
	for _, i := range words {
		trimWorld := strings.ToLower(i)
		cleanWorld = append(cleanWorld, trimWorld)
	}
	return cleanWorld
}

func commandExit(_ string, cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(string, *config) error
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the map",
			callback:    commandMap,
		},
		"explore": {
			name:        "explore",
			description: "explore each location",
			callback:    commandExplore,
		},
	}
}
