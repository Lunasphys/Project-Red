package main

import (
	"fmt"
)

type personnage struct {
	Nom                  string
	Classe               string
	Niveau               int
	Point_de_vie_max     int
	Point_de_vie_restant int
	Inventaire           []string
}

func (char *personnage) Init(Nom string, Classe string, Niveau int, Point_de_vie_max int, Point_de_vie_restant int, Inventaire []string) {
	char.Nom = Nom
	char.Classe = Classe
	char.Niveau = Niveau
	char.Point_de_vie_max = Point_de_vie_max
	char.Point_de_vie_restant = Point_de_vie_restant
	char.Inventaire = Inventaire
}

func main() {
	var char1 personnage
	char1.Init("Lunasphys", "Archer", 68, 1350, 1152, []string{"Potion", "Potion", "Potion"})
	char1.DisplayInfo()
	char1.AccessInventory()
	char1.TakePot()
}

func (char1 *personnage) DisplayInfo() {
	fmt.Println(char1.Nom)
	fmt.Println(char1.Classe)
	fmt.Println(char1.Niveau)
	fmt.Println(char1.Point_de_vie_max)
	fmt.Println(char1.Point_de_vie_restant)
	fmt.Println(char1.Inventaire)
}

func (char1 personnage) AccessInventory() {
	if len(char1.Inventaire) == 0 {
		fmt.Println("Votre inventaire est vide")
	}
	for i := 0; i < len(char1.Inventaire); i++ {
		fmt.Println(char1.Inventaire[i])
	}
}

func Remove(tab []string, s string) []string {
	for index, val := range tab {
		if val == s {
			tab[index] = tab[len(tab)-1]
			return tab[:len(tab)-1]
		}
	}
	return tab
}

func (char1 *personnage) TakePot() {
	for _, objet := range char1.Inventaire {
		if objet == "Potion" {
			if (char1.Point_de_vie_restant) < (char1.Point_de_vie_max) {
				char1.Inventaire = Remove(char1.Inventaire, objet)
				char1.Point_de_vie_restant += 150
				fmt.Println("Vous avez utilisé une potion")
				if char1.Point_de_vie_restant >= char1.Point_de_vie_max {
					char1.Point_de_vie_restant = char1.Point_de_vie_max
					fmt.Println("Vous ne pouvez pas utiliser de potion")
			} else {
				fmt.Println("Vous n'avez pas de potion")
				}
			}	
		}
	}
	char1.AccessInventory()
	fmt.Println(char1.Point_de_vie_restant)
}
