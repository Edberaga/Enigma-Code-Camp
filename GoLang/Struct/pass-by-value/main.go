package main

import "fmt"

type potion struct {
	name string
	desc string
	price int
	recover_HP int
	recover_MP int
	sellable bool
}

func main() {
	var hp_potion = potion{
		name: "HP Potion",
		desc: "Made by green herbs from forest. Restore small ammount HP",
		price: 90,
		recover_HP: 100,
		recover_MP: 0,
		sellable: true,
	}
	var grand_hp_potion = hp_potion 

	fmt.Println("-------------Sebelum diubah------------")
	fmt.Println("HP Potion", hp_potion)
	fmt.Println("Grand HP Potion", hp_potion)
	fmt.Printf("HP Potion memory address: %p\n", &hp_potion)
	fmt.Printf("Grand HP Potion memory address: %p\n", &grand_hp_potion)

	grand_hp_potion.desc = "Well advanced mix of green herbs. Restore large ammount HP"
	grand_hp_potion.price= 250
	grand_hp_potion.recover_HP = 300

	fmt.Println("-------------Setelah diubah------------")
	fmt.Println("HP Potion", hp_potion)
	fmt.Println("Grand HP Potion", hp_potion)
	fmt.Printf("HP Potion memory address: %p\n", &hp_potion)
	fmt.Printf("Grand HP Potion memory address: %p\n", &grand_hp_potion)

	var mp_potion = potion{
		name: "MP Potion",
		desc: "Made by several magic mushroom. Restore small ammount MP",
		price: 240,
		recover_HP: 0,
		recover_MP: 100,
		sellable: true,
	}
	fmt.Println(mp_potion)
	enchanceElixir(mp_potion)
}

func enchanceElixir(elixir potion) {
	elixir.name = "MP Elixir"
	elixir.desc = "Potion enchanced by holy water. Fully restore MP"
	elixir.recover_MP = 999
	elixir.sellable = false
	fmt.Println(elixir)
}