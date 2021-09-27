package main

import (
	"fmt"
	"time"
)

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

func (char1 *personnage) PoisonPot() {
	for _, objet := range char1.Inventaire {
		if objet == "Potion de poison" {
			fmt.Println("Applique un poison pendant 3 secondes")
			char1.Point_de_vie_restant -= 10
			time.Sleep(1 * time.Second)
			char1.Point_de_vie_restant -= 10
			time.Sleep(1 * time.Second)
			char1.Point_de_vie_restant -= 10
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println(char1.Point_de_vie_restant, "/" , char1.Point_de_vie_max)
}
