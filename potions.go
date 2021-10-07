package main

import (
	"fmt"
	"time"
)

func (char1 *personnage) TakePot() {

	// Does char1.Inventaire contains a potion ?
	// If no -> I dont have potion
	// If pv == pv max -> Potion useless
	// Prend la potion
	// J'enleve la potion
	// si pv > pv max alors pv = pv max

	if !Contains(char1.Inventaire, potionDeVie) {
		fmt.Println("Vous n'avez pas de potion")
	}
	if char1.Point_de_vie_max == char1.Point_de_vie_restant {
		fmt.Println("Vous ne pouvez pas utiliser de potion")
	}
	if Contains(char1.Inventaire, potionDeVie) {
		if char1.Point_de_vie_restant < char1.Point_de_vie_max {
			char1.Point_de_vie_restant += 25
			fmt.Println("Vous avez utilisé une potion de vie")
			char1.Inventaire = RemoveInventory(char1.Inventaire, potionDeVie)
		}
	}
	if char1.Point_de_vie_restant > char1.Point_de_vie_max {
		char1.Point_de_vie_restant = char1.Point_de_vie_max

	}
	fmt.Println("PV :", char1.Point_de_vie_restant, "/", char1.Point_de_vie_max)
	fmt.Println(char1.Inventaire)
}

func (char1 *personnage) PoisonPot() {
	if !Contains(char1.Inventaire, potionDePoison) {
		fmt.Println("Vous n'avez pas de potion de poison")
	}
	if Contains(char1.Inventaire, potionDePoison) {
		fmt.Println("Applique un poison pendant 3 secondes")
		char1.Point_de_vie_restant -= 10
		time.Sleep(1 * time.Second)
		fmt.Println("Il reste 2 secondes")
		char1.Point_de_vie_restant -= 10
		time.Sleep(1 * time.Second)
		fmt.Println("Il reste 1 seconde")
		char1.Point_de_vie_restant -= 10
		time.Sleep(1 * time.Second)
		fmt.Println("L'effet de la potion de poison s'estompe")
		char1.Inventaire = RemoveInventory(char1.Inventaire, potionDePoison)
	}
	fmt.Println("PV :", char1.Point_de_vie_restant, "/", char1.Point_de_vie_max)

}
