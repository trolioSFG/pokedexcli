package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"net/http"
	"encoding/json"
	"io"
)

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

type config struct {
	Next string
	Previous string
}


type locations struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	}
}

var commands = map[string]cliCommand {
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
	"help": {
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	},
	"map": {
		name: "map",
		description: "List 20 locations",
		callback: commandMap,
	},
	"mapb": {
		name: "mapb",
		description: "List 20 PREVIOUS locations",
		callback: commandMapBack,
	},
}

var helpMsg = ""


func cleanInput(text string) []string {
	// return []string{}
	clean := strings.ToLower(text)
	return strings.Fields(clean)
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println(helpMsg)
	return nil
}

func commandMap(c *config) error {
	var url string
	if c.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = c.Next
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error GET %v\n", err.Error())
		return err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error ReadAll")
		return err
	}
	defer res.Body.Close()

	// fmt.Println(string(data))

	resLocations := locations{}

	err = json.Unmarshal(data, &resLocations)
	if err != nil {
		return err
	}

	// fmt.Println(resLocations.Count)

	c.Next = resLocations.Next
	if resLocations.Previous != "" {
		c.Previous = resLocations.Previous
	} else {
		c.Previous = ""
	}

	for _, l := range resLocations.Results {
		fmt.Println(l.Name)
	}

	return nil
}


func commandMapBack(c *config) error {
	var url string
	if c.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	} else {
		url = c.Previous
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error GET %v\n", err.Error())
		return err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error ReadAll")
		return err
	}
	defer res.Body.Close()

	// fmt.Println(string(data))

	resLocations := locations{}

	err = json.Unmarshal(data, &resLocations)
	if err != nil {
		return err
	}

	// fmt.Println(resLocations.Count)

	c.Next = resLocations.Next
	if resLocations.Previous != "" {
		c.Previous = resLocations.Previous
	} else {
		c.Previous = ""
	}

	for _, l := range resLocations.Results {
		fmt.Println(l.Name)
	}

	return nil
}

	
func main() {

	helpMsg = `Welcome to the Pokedex!
Usage:

`

	for _, cmd := range commands {
		helpMsg += fmt.Sprintf("%v: %s\n", cmd.name, cmd.description)
	}

	var conf = config {}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		// command += scanner.Text()
		words := cleanInput(scanner.Text())
		// fmt.Println("Your command was:", strings.ToLower(words[0]))
		command, found := commands[words[0]]

		if !found {
			fmt.Println("Unknown command")
		} else {
			command.callback(&conf)
		}
	}

}

