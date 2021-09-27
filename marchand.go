package main

import "fmt"

func (char *Marchand) Init(Nom string, Inventaire []string, Argent int) {
	char.Nom = Nom
	char.Inventaire = Inventaire
	char.Argent = Argent
}

func (char2 *Marchand) DisplayInfoMarchand() {
	fmt.Println(char2.Nom)
	fmt.Println(char2.Inventaire)
	fmt.Println(char2.Argent)
}

func (char2 *Marchand) InventoryMarchand() {
	for i := 0; i < len(char2.Inventaire); i++ {
		fmt.Println(char2.Inventaire[i])
	}
}
