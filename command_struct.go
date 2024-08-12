package main

type cliCommand struct {
	name        string
	description string
	callback    func(apiState *config, params []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Shows the next page of all locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous page of all locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Show the possible encounters in a location, the location must be passed as a parameter",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon, the pokemon name must be passed as a parameter",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Shows caught pokemon information, a pokemon must be passed as a parameter",
			callback:    commandInspect,
		},
	}
}
