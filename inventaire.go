package main

import (
	"fmt"
)

func (char1 personnage) AccessInventory() {
	if len(char1.Inventaire) == 0 {
		fmt.Println("Votre inventaire est vide")
	}
	for i := 0; i < len(char1.Inventaire); i++ {
		fmt.Println(char1.Inventaire[i])
	}
}

func RemoveInventory(tab []string, s string) []string {
	for index, val := range tab {
		if val == s {
			tab[index] = tab[len(tab)-1]
			return tab[:len(tab)-1]
		}
	}
	return tab
}

func AddInventory(tab []string, s string) []string {
	for index, val := range tab {
		if val == s {
			tab[index] = tab[len(tab)+1]
			return tab[:len(tab)+1]
		}
	}
	return tab
}

func (char1 *personnage) TakePot() {
	for _, objet := range char1.Inventaire {
		if objet == "Potion" {
			if (char1.Point_de_vie_restant) < (char1.Point_de_vie_max) {
				char1.Inventaire = RemoveInventory(char1.Inventaire, objet)
				char1.Point_de_vie_restant += 150
				fmt.Println("Vous avez utilisé une potion")
				if char1.Point_de_vie_restant >= char1.Point_de_vie_max {
					char1.Point_de_vie_restant = char1.Point_de_vie_max
					fmt.Println("Vous ne pouvez pas utiliser de potion")
					if objet != "Potion" {
						fmt.Println("Vous n'avez pas de potion")
					}
				}
			}
			break
		}
	}
	char1.AccessInventory()
	fmt.Println(char1.Point_de_vie_restant)
}
