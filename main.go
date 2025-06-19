package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"net/http"
	"encoding/json"
	"io"
	"github.com/trolioSFG/internal/pokecache"
	"time"
	"math/rand"
)


type config struct {
	Next string
	Previous string
	args []string
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

type Encounter struct {
	Encounter_Method_Rates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`

	} `json:"encounter_method_rates"`

	Game_Index int `json:"game_index"`
	Id int `json:"id"`
	
	Location struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"location"`
	
	Name string `json:"name"`
	
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`

	Pokemon_Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance int `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel int `json:"max_level"`
				Method struct {
					Name string `json:"name"`
					URL string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version struct {
				Name string `json:"name"`
				URL string `json:"url"`
			} `json:"version"`
		} `json:version_details"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Cries          struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Height    int `json:"height"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			Order        any `json:"order"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name          string `json:"name"`
	Order         int    `json:"order"`
	PastAbilities []struct {
		Abilities []struct {
			Ability  any  `json:"ability"`
			IsHidden bool `json:"is_hidden"`
			Slot     int  `json:"slot"`
		} `json:"abilities"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"past_abilities"`
	PastTypes []any `json:"past_types"`
	Species   struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string `json:"back_default"`
		BackFemale       string `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  string `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      string `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale string `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string `json:"front_default"`
				FrontFemale  any    `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     string `json:"front_default"`
				FrontFemale      string `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale string `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
			Showdown struct {
				BackDefault      string `json:"back_default"`
				BackFemale       string `json:"back_female"`
				BackShiny        string `json:"back_shiny"`
				BackShinyFemale  any    `json:"back_shiny_female"`
				FrontDefault     string `json:"front_default"`
				FrontFemale      string `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale string `json:"front_shiny_female"`
			} `json:"showdown"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault      string `json:"back_default"`
					BackGray         string `json:"back_gray"`
					BackTransparent  string `json:"back_transparent"`
					FrontDefault     string `json:"front_default"`
					FrontGray        string `json:"front_gray"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault      string `json:"back_default"`
					BackGray         string `json:"back_gray"`
					BackTransparent  string `json:"back_transparent"`
					FrontDefault     string `json:"front_default"`
					FrontGray        string `json:"front_gray"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault           string `json:"back_default"`
					BackShiny             string `json:"back_shiny"`
					BackShinyTransparent  string `json:"back_shiny_transparent"`
					BackTransparent       string `json:"back_transparent"`
					FrontDefault          string `json:"front_default"`
					FrontShiny            string `json:"front_shiny"`
					FrontShinyTransparent string `json:"front_shiny_transparent"`
					FrontTransparent      string `json:"front_transparent"`
				} `json:"crystal"`
				Gold struct {
					BackDefault      string `json:"back_default"`
					BackShiny        string `json:"back_shiny"`
					FrontDefault     string `json:"front_default"`
					FrontShiny       string `json:"front_shiny"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"gold"`
				Silver struct {
					BackDefault      string `json:"back_default"`
					BackShiny        string `json:"back_shiny"`
					FrontDefault     string `json:"front_default"`
					FrontShiny       string `json:"front_shiny"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      string `json:"back_default"`
						BackFemale       string `json:"back_female"`
						BackShiny        string `json:"back_shiny"`
						BackShinyFemale  string `json:"back_shiny_female"`
						FrontDefault     string `json:"front_default"`
						FrontFemale      string `json:"front_female"`
						FrontShiny       string `json:"front_shiny"`
						FrontShinyFemale string `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  string `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}


type cliCommand struct {
	name string
	description string
	callback func(*config) error
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
	"explore": {
		name: "explore",
		description: "List pokemons in location argument",
		callback: commandExplore,
	},
	"catch": {
		name: "catch",
		description: "Catch a pokemon",
		callback: commandCatch,
	},
	"inspect": {
		name: "inspect",
		description: "Info about a CATCHED Pokemon",
		callback: commandInspect,
	},
	"pokedex": {
		name: "pokedex",
		description: "Pokedex",
		callback: commandPokedex,
	},
}

var helpMsg = ""

var mycache pokecache.Cache
var pokedex = map[string]Pokemon {}


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

	data, found := mycache.Get(url)

	if !found {
		fmt.Println("Cache miss")

		res, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error GET %v\n", err.Error())
			return err
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error ReadAll")
			return err
		}
		defer res.Body.Close()

		mycache.Add(url, data)
	}

	// fmt.Println(string(data))

	resLocations := locations{}

	err := json.Unmarshal(data, &resLocations)
	if err != nil {
		return err
	}

	// fmt.Println(resLocations.Count)

	c.Next = resLocations.Next
	if resLocations.Previous != "" {
		c.Previous = resLocations.Previous
	} else {
		c.Previous = url
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

	data, found := mycache.Get(url)

	if !found {
		res, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error GET %v\n", err.Error())
			return err
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error ReadAll")
			return err
		}
		defer res.Body.Close()

		mycache.Add(url, data)
	}

	// fmt.Println(string(data))

	resLocations := locations{}

	err := json.Unmarshal(data, &resLocations)
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

func commandExplore(c *config) error {
	if len(c.args) == 0 {
		fmt.Println("Missing location")
		return fmt.Errorf("Missing required location")
	}

	// fmt.Println(c.args)
	url := "https://pokeapi.co/api/v2/location-area/" + c.args[0]

	data, found := mycache.Get(url)

	if !found {
		res, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error GET %v\n", err.Error())
			return err
		}

		if res.StatusCode > 299 {
			fmt.Println("Bad petition ", res.Status)
			return fmt.Errorf("Bad petition")
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error ReadAll")
			return err
		}
		defer res.Body.Close()

		mycache.Add(url, data)
	}

	// fmt.Println(string(data))

	encounter := Encounter{}
	err := json.Unmarshal(data, &encounter)
	if err != nil {
		return err
	}

	for _, v := range encounter.Pokemon_Encounters {
		fmt.Println(v.Pokemon.Name)
	}

	return nil
}

func commandCatch(c *config) error {
	if len(c.args) < 1 {
		fmt.Println("Missing pokemon name")
		return fmt.Errorf("Missing pokemon name")
	}

	// fmt.Println(c.args)
	pokemonName := c.args[0]
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	data, found := mycache.Get(url)

	if !found {
		res, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error GET %v\n", err.Error())
			return err
		}

		if res.StatusCode > 299 {
			fmt.Println("Bad petition ", res.Status)
			return fmt.Errorf("Bad petition")
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error ReadAll")
			return err
		}
		defer res.Body.Close()

		mycache.Add(url, data)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemon := Pokemon{}
	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		fmt.Println("Unmarshal error")
		return err
	}

	// Typical base experience 31-255
	result := rand.Intn(500)
	fmt.Printf("Base experience: %d Result: %d\n", pokemon.BaseExperience, result)

	if result < (500 - pokemon.BaseExperience) {
		// Catched
		fmt.Printf("%s was caught!\n", pokemonName)
		pokedex[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

func commandInspect(c *config) error {
	if len(c.args) < 1 {
		// fmt.Printf("Missing pokemon name")
		return fmt.Errorf("Missing pokemon name")
	}

	name := c.args[0]
	poke, ok := pokedex[name]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Println("Name:", poke.Name)
	fmt.Println("Height:", poke.Height)
	fmt.Println("Weight:", poke.Weight)
	fmt.Println("Stats:")
	fmt.Println("\t-hp:", poke.Stats[0].BaseStat)
	fmt.Println("\t-attack:", poke.Stats[1].BaseStat)
	fmt.Println("\t-defense:", poke.Stats[2].BaseStat)
	fmt.Println("\t-special-attack:", poke.Stats[3].BaseStat)
	fmt.Println("\t-special-defense:", poke.Stats[4].BaseStat)
  	fmt.Println("\t-speed:", poke.Stats[5].BaseStat)
	fmt.Println("Types:")
	for _,v := range poke.Types {
		fmt.Println("\t-", v.Type.Name)
	}

	return nil
}

func commandPokedex(c *config) error {
	fmt.Println("Your Pokedex:")
	for name, _ := range pokedex {
		fmt.Println("-", name)
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
	mycache = pokecache.NewCache(5 * time.Second)

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
			if len(words) > 1 {
				conf.args = words[1:]
			} else {
				conf.args = []string{}
			}

			err := command.callback(&conf)
			if err != nil {
				fmt.Printf("ERROR: %v\n", err)
			}
		}
	}

}

