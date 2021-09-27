package main

import "fmt"

type Marchand struct {
	Nom        string
	Inventaire []string
}

func (char2 *personnage) DisplayInfoMarchand() {
	fmt.Println(char2.Nom)
	fmt.Println(char2.Classe)
	fmt.Println(char2.Niveau)
	fmt.Println(char2.Point_de_vie_max)
	fmt.Println(char2.Point_de_vie_restant)
	fmt.Println(char2.Inventaire)
	fmt.Println(char2.Money)
}

func (char2 *personnage) InventoryMarchand() {
	for i := 0; i < len(char2.Inventaire); i++ {
		fmt.Println(char2.Inventaire[i])
	}
}
