package main

import "fmt"

type weapon struct {
	weapon_type string
	attack int
	hit_rate float32
	cri_rate float32
}

type character struct {
	name string
	strength int
	stamina int
	speed int
	weapon
}

func main() {
	var stanly = character{
		name: "Stanly",
		strength: 100,
		stamina: 110,
		speed: 100,
		weapon: weapon{
			weapon_type: "Rifle",
			attack: 88,
			hit_rate: 105.51,
			cri_rate: 11.08,
		},
	}

	fmt.Println(stanly.weapon.weapon_type)
}