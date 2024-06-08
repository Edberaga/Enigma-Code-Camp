package main

import "fmt"

func main() {
	edbert := map[string]int{
		"Strength": 126,
		"Dexterity": 85,
		"Intellect": 114,
	}
	fmt.Println("Edbert Stat: ", edbert)
	fmt.Println("========================================================================")

	sydnia := make(map[string]int)
	fmt.Println(sydnia)
	sydnia["Strength"] = 61
	sydnia["Dexterity"] = 109
	sydnia["Intellect"] = 135
	fmt.Println("Sydnia Stat: ", sydnia)

	fmt.Println("Edbert Strength: ", edbert["Strength"])
	fmt.Println("Sydnia Dexterity: ", sydnia["Dexterity"])
	fmt.Println("Edbert Charisma", edbert["Charisma"])

	sydnia["Strength"] = 74
	fmt.Println("Sydnia Stat: ", sydnia)
	//delete(sydnia, "Dexterity")
	fmt.Println("Sydnia Stat: ", sydnia)
	
	for stat := range edbert {
		fmt.Println(stat, ": ", edbert[stat])
	}

	for stat, num := range sydnia {
		fmt.Println(stat, ": ", num)
	}
	fmt.Println("========================================================================")


}