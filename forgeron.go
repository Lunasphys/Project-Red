package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountItem(tab []string, s string) int { // Permet de compter les objets de même nom
	var res int = 0
	for _, a := range tab {
		if a == s {
			res++
		}
	}
	return res
}

func (char1 *personnage) Craft(quantityA int, quantityB int, itemA string, itemB string) bool { // Permet de savoir s'il a le nombre d'objet demandé en vu de fabriquer un item
	if CountItem(char1.Inventaire, itemA) >= quantityA && CountItem(char1.Inventaire, itemB) >= quantityB {
		return true
	}
	return false
}

func menuCraft(char1 *personnage) { // Ce que le joueur peut forger comme item et les composants demandés
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("--------------------FORGERON---------------------")
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	SlowPrint("Bienvenue chez le forgeron\n")
	SlowPrint("Que voulez-vous forger\n ")
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("1 = Craft un chapeau de l'aventurier ")
	fmt.Println("(1 Plume de Corbeau et 1 Cuir de Sanglier)")
	fmt.Println()
	fmt.Println("2 = Craft une tunique de l'aventurier ")
	fmt.Println("(2 Fourrure de Loup et 1 Peau de Troll)")
	fmt.Println()
	fmt.Println("3 = Craft des bottes de l'aventurier ")
	fmt.Println("(1 Fourrure de Loup et 1 Cuir de Sanglier)")
	fmt.Println()
	fmt.Println("4 = Quitter le forgeron")
	scanner.Scan() // l'utilisateur input dans la console
	d := scanner.Text()
	switch d {
	case "1":
		char1.Making(5, 1, 1, "Plume de Corbeau", "Cuir de Sanglier", chapeauAventurier)
	case "2":
		char1.Making(5, 2, 1, "Fourrure de Loup", "Peau de Troll", tuniqueAventurier)
	case "3":
		char1.Making(5, 1, 1, "Fourrure de Loup", "Cuir de Sanglier", bottesAventurier)
	case "4":
		break
	}
}
func (char1 *personnage) Making(argent1 int, nbobjet1 int, nbobjet2 int, objet1 string, objet2 string, equipement string) {
	if len(char1.Inventaire) < char1.Tailleinventairemax { // Voir si il y a assez de place dans l'inventaire
		if char1.Argent >= argent1 { // Verifie si le joueur a assez d'argent pour effectuer cette action
			if char1.Craft(nbobjet1, nbobjet2, objet1, objet2) {
				char1.Inventaire = char1.AddInventory(char1.Inventaire, equipement) // Ajoute l'item dans l'inventaire
				char1.Inventaire = RemoveInventory(char1.Inventaire, objet1)        // Retire l'objet utilisé pour le craft de l'inventaire
				char1.Inventaire = RemoveInventory(char1.Inventaire, objet2)        // Retire l'objet utilisé pour le craft de l'inventaire
				char1.Argent -= argent1
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous venez de dépenser", argent1, " d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous venez de créer un ", equipement)
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			} else { // Indique au joueur qu'il n'a pas les objets nécessaires pour effectuer cette action
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous n'avez pas les ressources nécessaire pour créer cet objet")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Votre inventaire ne contient que les objets suivants :", char1.Inventaire) // Renvoi l'inventaire du joueur
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

			}
		} else { // Indique au joueur qu'il n'a pas assez de pièces d'or pour forger un item
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Vous êtes trop pauvre pour pouvoir forger un item")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Vous n'avez que :", char1.Argent, "pièces d'or")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		}
	} else { // Indique au joueur qu'il a atteint la limite de son inventaire
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println("Votre inventaire contient les objets suivants :", char1.Inventaire) // Renvoi l'inventaire du joueur
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	}
}
