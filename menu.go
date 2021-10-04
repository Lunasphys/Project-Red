package main

import (
	"bufio"
	"fmt"
	"os"
)

func (char1 *personnage) menu(char2 *Marchand) {
	for {
		fmt.Println("0 = Afficher les informations du personnages")
		fmt.Println("1 = Acceder au menue de l'inventaire")
		fmt.Println("2 = Marchand")
		fmt.Println("3 = Acceder a la liste de sorts")
		fmt.Println("4 = Acceder au menu du forgeron")
		fmt.Println("5 = Quitter")
		// créer une var scanner qui va lire ce que l'utilisateur va écrire
		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan() // l'utilisateur input dans la console

		// lis ce que l'utilisateur a écrit
		v := scanner.Text()
		switch v {
		case "0":
			char1.DisplayInfo()
		case "1":
			char1.AccessInventory()
			fmt.Println("1 = Choisissez une potion de vie")
			fmt.Println("2 = Choisissez une potion de poison")
			fmt.Println("3 = Apprendre le sort : Boule de feu")
			fmt.Println("4 = Ne rien choisir et quitter")
			scanner.Scan() // l'utilisateur input dans la console
			n := scanner.Text()
			switch n {
			case "1":
				char1.TakePot()
			case "2":
				char1.PoisonPot()
			case "3":
				char1.LearnSkill()
			case "4":
				break
			}
		case "2":
			char2.Inventaire2()
			fmt.Println("1 = Acheter une potion de vie (3 pièces d'or)")
			fmt.Println("2 = Acheter une potion de poison (6 pièces d'or)")
			fmt.Println("3 = Acheter Livre de sort : Boule de feu (25 pièces d'or)")
			fmt.Println("4 = Acheter une Fourrure de Loup (4 pièces d'or)")
			fmt.Println("5 = Acheter une Peau de Troll (7 pièces d'or)")
			fmt.Println("6 = Acheter une Cuir de Sanglier (3 pièces d'or)")
			fmt.Println("7 = Acheter une Plume de Corbeau (6 pièces d'or)")
			fmt.Println("8 = Quitter l'inventaire du marchand")
			scanner.Scan() // l'utilisateur input dans la console
			b := scanner.Text()
			switch b {
			case "1":
				if len(char1.Inventaire) < 10 {
					if char1.Argent >= 3 {
						char1.Inventaire = AddInventory(char1.Inventaire, "Potion de vie")
						char1.Argent -= 3
						fmt.Println(char1.Inventaire)
					} else {
						fmt.Println("Vous êtes trop pauvre pour pouvoir acheter un item")
					}
				} else {
					fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
				}
			case "2":
				if len(char1.Inventaire) < 10 {
					if char1.Argent >= 6 {
						char1.Inventaire = AddInventory(char1.Inventaire, "Potion de poison")
						char1.Argent -= 6
					} else {
						fmt.Println("Vous êtes trop pauvre pour pouvoir acheter un item")
					}
				} else {
					fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
				}
			case "3":
				if len(char1.Inventaire) < 10 {
					if char1.Argent >= 25 {
						char1.Inventaire = AddInventory(char1.Inventaire, "Livre de sort : Boule de feu")
						char1.Argent -= 25
						fmt.Println(char1.Inventaire)
					} else {
						fmt.Println("Vous êtes trop pauvre pour pouvoir acheter un item")
					}
				} else {
					fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
				}
			case "4":
				if len(char1.Inventaire) < 10 {
					if char1.Argent >= 4 {
						char1.Inventaire = AddInventory(char1.Inventaire, "Fourrure de Loup")
						char1.Argent -= 4
					} else {
						fmt.Println("Vous êtes trop pauvre pour pouvoir acheter un item")
					}
				} else {
					fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
				}
			case "5":
				if len(char1.Inventaire) < 10 {
					if char1.Argent >= 7 {
						char1.Inventaire = AddInventory(char1.Inventaire, "Peau de Troll")
						char1.Argent -= 7
					} else {
						fmt.Println("Vous êtes trop pauvre pour pouvoir acheter un item")
					}
				} else {
					fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
				}
			case "6":
				if len(char1.Inventaire) < 10 {
					if char1.Argent >= 3 {
						char1.Inventaire = AddInventory(char1.Inventaire, "Cuir de Sanglier")
						char1.Argent -= 3
						fmt.Println(char1.Inventaire)
					} else {
						fmt.Println("Vous êtes trop pauvre pour pouvoir acheter un item")
					}
				} else {
					fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
				}
			case "7":
				if len(char1.Inventaire) < 10 {
					if char1.Argent >= 1 {
						char1.Inventaire = AddInventory(char1.Inventaire, "Plume de Corbeau")
						char1.Argent -= 1
					} else {
						fmt.Println("Vous êtes trop pauvre pour pouvoir acheter un item")
					}
				} else {
					fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
				}
			case "8":
				break
			}
		case "3":
			char1.AccessSkill()
			fmt.Println("1 = ")
			fmt.Println("2 = Quitter la liste de sort")
			scanner.Scan() // l'utilisateur input dans la console
			c := scanner.Text()
			switch c {
			case "1":
			case "2":
				break

			}
		case "4":
			fmt.Println("1 = Craft un chapeau de l'aventurier ")
			fmt.Println("2 = Craft une tunique de l'aventurier ")
			fmt.Println("3 = Craft des bottes de l'aventurier ")
			fmt.Println("4 = Quitter le forgeron")
			scanner.Scan() // l'utilisateur input dans la console
			d := scanner.Text()
			switch d {
			case "1":
				if len(char1.Inventaire) < 10 {
					if char1.Argent >= 5 {
						if char1.Craft(1, 1, "Plume de Corbeau", "Cuir de Sanglier") {
							char1.Inventaire = AddInventory(char1.Inventaire, "Chapeau de l'aventurier")
							char1.Inventaire = RemoveInventory(char1.Inventaire, "Plume de Corbeau")
							char1.Inventaire = RemoveInventory(char1.Inventaire, "Cuir de Sanglier")
							char1.Argent -= 5
							fmt.Println("Vous venez de créer un Chapeau de l'Aventurier")
						} else {
							fmt.Println("Vous n'avez pas les ressources nécessaire pour créer cet objet")
						}
					} else {
						fmt.Println("Vous êtes trop pauvre pour pouvoir acheter un item")
					}
				} else {
					fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
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
							fmt.Println("Vous venez de créer une Tunique de l'aventurier")
						} else {
							fmt.Println("Vous n'avez pas les ressources nécessaire pour créer cet objet")
						}
					} else {
						fmt.Println("Vous êtes trop pauvre pour pouvoir acheter un item")
					}
				} else {
					fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
				}
			case "3":
				if len(char1.Inventaire) < 10 {
					if char1.Argent >= 5 {
						if char1.Craft(1, 1, "Fourrure de Loup", "Cuir de Sanglier") {
							char1.Inventaire = AddInventory(char1.Inventaire, "Bottes de l'Aventurier")
							char1.Inventaire = RemoveInventory(char1.Inventaire, "Fourrure de Loup")
							char1.Inventaire = RemoveInventory(char1.Inventaire, "Cuir de Sanglier")
							char1.Argent -= 5
							fmt.Println("Vous venez de créer des Bottes de l'aventurier")
						} else {
							fmt.Println("Vous n'avez pas les ressources nécessaire pour créer cet objet")
						}
					} else {
						fmt.Println("Vous êtes trop pauvre pour pouvoir acheter un item")
					}
				} else {
					fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
				}
			case "4":
				break
			}
		case "5":
			os.Exit(2)
		}
	}
}
