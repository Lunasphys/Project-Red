package main

import (
	"fmt"
)

func (char1 *personnage) AccessInventory() { // Permet d'accéder à l'inventaire
	if len(char1.Inventaire) == 0 { // S'il n'y a rien dans l'inventaire, renvoi au joueur que son inventaire est vide
		fmt.Println("Votre inventaire est vide")
	}
	for i := 0; i < len(char1.Inventaire); i++ { // Permet de compter le nombre d'objet présent dans l'inventaire 
		fmt.Println(char1.Inventaire[i])
	}
}

func RemoveInventory(tab []string, s string) []string { // Permet de retirer des objets de l'inventaire
	for index, val := range tab {
		if val == s {
			tab[index] = tab[len(tab)-1]
			return tab[:len(tab)-1]
		}
	}
	return tab
}

func AddInventory(tab []string, s string) []string { // Permet d'ajouter des objets dans l'inventaire
	if len(tab)-1 >= 9 {
		fmt.Println("Vous n'avez plus de place dans votre inventaire")
		return tab
	}
	return append(tab, s)
}

