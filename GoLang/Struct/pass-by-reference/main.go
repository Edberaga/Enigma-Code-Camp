package main

import "fmt"

type sword struct {
	name string
	attack int
	hit_rate int
	cri_rate int
	desc string
	price int
	sellable bool
}

func main() {
	var saber = sword{
		name: "Silver Saber",
		attack: 112,
		hit_rate: 95,
		cri_rate: 12,
		desc: "Sword craft by high value silvers. Used by Empire Soldier",
		price: 1250,
		sellable: true,
	}
	var long_saber *sword = &saber
	fmt.Println(saber)
	fmt.Println(long_saber) //Long saber value is the Silver Saber memory address
	fmt.Printf("Silver Saber memory: %p\n", &saber) //Silver saber and Long Saber variable has different memory
	fmt.Printf("Long Saber memory: %p\n", &long_saber)

	saber.price = 500
	long_saber.name = "Long Saber"
	fmt.Println("-------------Sebelum diubah------------")
	fmt.Println(saber)
	fmt.Println(long_saber)

	fmt.Println("-------------Setelah diubah------------")
	enchanceWeapon(&saber)
	fmt.Println(saber)
	fmt.Println(long_saber)
}

func enchanceWeapon(enchance *sword) {
	enchance.attack *= 2
	enchance.hit_rate += 3
	enchance.cri_rate += 5
	fmt.Println(enchance)
}