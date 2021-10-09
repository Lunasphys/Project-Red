package main

import (
	"fmt"
	"time"
)

func (char1 *personnage) TakePot() { //Fonction qui permet d'utiliser une potion de vie

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

func (char1 *personnage) PoisonPot() { // Fonction qui permet d'utiliser une potion de poison
	if !Contains(char1.Inventaire, potionDePoison) {
		fmt.Println("Vous n'avez pas de potion de poison")
	}
	if Contains(char1.Inventaire, potionDePoison) {
		fmt.Println("Applique un poison pendant 3 secondes")
		char1.Point_de_vie_restant -= 10
		time.Sleep(1 * time.Second)
		fmt.Println("PV :", char1.Point_de_vie_restant, "/", char1.Point_de_vie_max)
		fmt.Println("Il reste 2 secondes")
		char1.Point_de_vie_restant -= 10
		time.Sleep(1 * time.Second)
		fmt.Println("PV :", char1.Point_de_vie_restant, "/", char1.Point_de_vie_max)
		fmt.Println("Il reste 1 seconde")
		char1.Point_de_vie_restant -= 10
		time.Sleep(1 * time.Second)
		fmt.Println("L'effet de la potion de poison s'estompe")
		char1.Inventaire = RemoveInventory(char1.Inventaire, potionDePoison)
	}
	fmt.Println("PV :", char1.Point_de_vie_restant, "/", char1.Point_de_vie_max)

}

func (char1 *personnage) ThrowPoisonPot(char3 *monstre) { //Fonction qui permet d'envoyer une potion de poison sur le gobelin d'entrainement
	if !Contains(char1.Inventaire, potionDePoison) {
		fmt.Println("Vous n'avez pas de potion de poison")
	}
	if Contains(char1.Inventaire, potionDePoison) {
		fmt.Println("Applique un poison sur l'ennemi pendant 3 secondes")
		char3.Point_de_vie_restant -= 10
		time.Sleep(1 * time.Second)
		fmt.Println("PV :", char3.Point_de_vie_restant, "/", char3.Point_de_vie_max)
		fmt.Println("Il reste 2 secondes")
		char3.Point_de_vie_restant -= 10
		fmt.Println("PV :", char3.Point_de_vie_restant, "/", char3.Point_de_vie_max)
		time.Sleep(1 * time.Second)
		fmt.Println("Il reste 1 seconde")
		char3.Point_de_vie_restant -= 10
		time.Sleep(1 * time.Second)
		fmt.Println("L'effet de la potion de poison s'estompe")
		char1.Inventaire = RemoveInventory(char1.Inventaire, potionDePoison)
	}
	fmt.Println("PV :", char3.Point_de_vie_restant, "/", char3.Point_de_vie_max)
}

func (char1 *personnage) ThrowLifePot(char3 *monstre) { //Fonction qui permet d'envoyer une potion de vie sur le gobelin d'entrainement

	if !Contains(char1.Inventaire, potionDeVie) {
		fmt.Println("Vous n'avez pas de potion à envoyer")
	}
	if Contains(char1.Inventaire, potionDeVie) {
		if char3.Point_de_vie_restant <= char3.Point_de_vie_max {
			char3.Point_de_vie_restant += 25
			fmt.Println("Vous venez de soigner le gobelin, vous avez un QI négatif")
			char1.Inventaire = RemoveInventory(char1.Inventaire, potionDeVie)
		}
	}
	if char3.Point_de_vie_restant > char3.Point_de_vie_max {
		char3.Point_de_vie_restant = char3.Point_de_vie_max

	}
	fmt.Println("PV :", char3.Point_de_vie_restant, "/", char3.Point_de_vie_max)
	fmt.Println(char1.Inventaire)
}

func (char1 *personnage) TakeManaPot() { //Fonction qui permet d'utiliser une potion de vie

	if !Contains(char1.Inventaire, potionDeMana) {
		fmt.Println("Vous n'avez pas de potion de Mana")
	}
	if char1.Point_de_mana_max == char1.Point_de_mana_restant {
		fmt.Println("Vous ne pouvez pas utiliser de potion de Mana")
	}
	if Contains(char1.Inventaire, potionDeMana) {
		if char1.Point_de_mana_restant < char1.Point_de_mana_max {
			char1.Point_de_mana_restant += 20
			fmt.Println("Vous avez utilisé une potion de Mana")
			char1.Inventaire = RemoveInventory(char1.Inventaire, potionDeMana)
		}
	}
	if char1.Point_de_mana_restant > char1.Point_de_mana_max {
		char1.Point_de_mana_restant = char1.Point_de_mana_max

	}
	fmt.Println("Vous avez maintenant :", char1.Point_de_mana_restant, "/", char1.Point_de_mana_max, "Points de Mana")
	fmt.Println(char1.Inventaire)
}
