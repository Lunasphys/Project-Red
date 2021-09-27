package main

import "fmt"

type Marchand struct {
	Nom        string
	Inventaire []string
}

func (char2 *personnage) DisplayInfoMarchand() {
	fmt.Println(char2.Nom)
	fmt.Println(char2.Inventaire)
}

func (char2 *personnage) InventoryMarchand() {
	for i := 0; i < len(char2.Inventaire); i++ {
		fmt.Println(char2.Inventaire[i])

	}
}
