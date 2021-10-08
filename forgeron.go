package main

import ("fmt"
	"bufio"
	"os"
)

func CountItem(tab []string, s string) int {
	var res int = 0
	for _, a := range tab {
		if a == s {
			res++
		}
	}
	return res
}

func (char1 *personnage) Craft(quantityA int, quantityB int, itemA string, itemB string) bool {
	if CountItem(char1.Inventaire, itemA) >= quantityA && CountItem(char1.Inventaire, itemB) >= quantityB {
		return true
	}
	return false
}

func menuCraft(char1 *personnage) {
	scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("1 = Craft un chapeau de l'aventurier ")
			fmt.Println("(1 Plume de Corbeau et 1 Cuir de Sanglier)")
			fmt.Println("2 = Craft une tunique de l'aventurier ")
			fmt.Println("(2 Fourrure de Loup et 1 Peau de Troll)")
			fmt.Println("3 = Craft des bottes de l'aventurier ")
			fmt.Println("(1 Fourrure de Loup et 1 Cuir de Sanglier)")
			fmt.Println("4 = Quitter le forgeron")
			scanner.Scan() // l'utilisateur input dans la console
			d := scanner.Text()
			switch d {
			case "1":
				if len(char1.Inventaire) < 10 {
					if char1.Argent >= 5 {
						if char1.Craft(1, 1, "Plume de Corbeau", "Cuir de Sanglier") {
							char1.Inventaire = AddInventory(char1.Inventaire, chapeauAventurier)
							char1.Inventaire = RemoveInventory(char1.Inventaire, "Plume de Corbeau")
							char1.Inventaire = RemoveInventory(char1.Inventaire, "Cuir de Sanglier")
							char1.Argent -= 5
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
							fmt.Println("Vous venez de dépenser 5 pièces d'or")
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
							fmt.Println("Vous venez de créer un Chapeau de l'aventurier")
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
						} else {
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
							fmt.Println("Vous n'avez pas les ressources nécessaire pour créer cet objet")
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
							fmt.Println("Vous venez de créer un Chapeau de l'aventurier")
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

						}
					} else {
						fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
						fmt.Println("Vous êtes trop pauvre pour pouvoir forger un item")
						fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
						fmt.Println("Vous n'avez que :", char1.Argent)
						fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
					}
				} else {
					fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
					fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
					fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
					fmt.Println(char1.Inventaire)
					fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				}
			case "2":
				if len(char1.Inventaire) < 10 {
					if char1.Argent >= 5 {
						if char1.Craft(2, 1, "Fourrure de Loup", "Peau de Troll") {
							char1.Inventaire = AddInventory(char1.Inventaire, "Tunique de l'aventurier")
							char1.Inventaire = RemoveInventory(char1.Inventaire, "Fourrure de Loup")
							char1.Inventaire = RemoveInventory(char1.Inventaire, "Fourrure de Loup")
							char1.Inventaire = RemoveInventory(char1.Inventaire, "Peau de Troll")
							char1.Argent -= 5
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
							fmt.Println("Vous venez de créer une Tunique de l'aventurier")
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
							fmt.Println("Vous venez de dépenser 5 pièces d'or", "il vous reste :", char1.Argent, "pièces d'or")
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
							fmt.Println(char1.Inventaire)
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
						} else {
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
							fmt.Println("Vous n'avez pas les ressources nécessaire pour créer cet objet")
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
							fmt.Println(char1.Inventaire)
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
						}
					} else {
						fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
						fmt.Println("Vous êtes trop pauvre pour pouvoir forger un item")
						fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
						fmt.Println("Vous n'avez que :", char1.Argent)
						fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
					}
				} else {
					fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
					fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
					fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
					fmt.Println(char1.Inventaire)
					fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				}
			case "3":
				if len(char1.Inventaire) < 10 {
					if char1.Argent >= 5 {
						if char1.Craft(1, 1, "Fourrure de Loup", "Cuir de Sanglier") {
							char1.Inventaire = AddInventory(char1.Inventaire, "Bottes de l'Aventurier")
							char1.Inventaire = RemoveInventory(char1.Inventaire, "Fourrure de Loup")
							char1.Inventaire = RemoveInventory(char1.Inventaire, "Cuir de Sanglier")
							char1.Argent -= 5
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
							fmt.Println("Vous venez de créer des Bottes de l'aventurier")
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
							fmt.Println("Vous venez de dépenser 5 pièces d'or", "il vous reste :", char1.Argent, "pièces d'or")
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
							fmt.Println(char1.Inventaire)
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
						} else {
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
							fmt.Println("Vous n'avez pas les ressources nécessaire pour créer cet objet")
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
							fmt.Println(char1.Inventaire)
							fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
						}
					} else {
						fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
						fmt.Println("Vous êtes trop pauvre pour pouvoir forger un item")
						fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
						fmt.Println("Vous n'avez que :", char1.Argent)
						fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
					}
				} else {
					fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
					fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
					fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
					fmt.Println(char1.Inventaire)
					fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				}
			case "4":
				break
			}
}