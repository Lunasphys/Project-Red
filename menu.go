package main

import (
	"bufio"
	"fmt"
	"os"
)

func (char1 *personnage) menu(char2 *Marchand, char3 *monstre) {
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("Que souhaitez-vous faire ?")
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	for {
		fmt.Println("0 = Afficher les informations du personnages")
		fmt.Println("1 = Acceder au menu de l'inventaire")
		fmt.Println("2 = Marchand")
		fmt.Println("3 = Acceder a la liste de sorts")
		fmt.Println("4 = Acceder au menu du forgeron")
		fmt.Println("5 = Acceder a l'équipement")
		fmt.Println("6 = Engager un combat d'entrainement")
		fmt.Println("7 = Quitter")
		// créer une var scanner qui va lire ce que l'utilisateur va écrire
		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan() // l'utilisateur input dans la console

		// lis ce que l'utilisateur a écrit
		v := scanner.Text()
		switch v {
		case "0":
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Informations de votre personnage")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			char1.DisplayInfo()
		case "1":
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Que souhaitez-vous faire ?")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
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
				fmt.Println(char1.Skill)
			case "4":
				break
			}
		case "2":
			char2.returnMarchand(char1)
		case "3":
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Quel sort souhaitez vous utiliser ?")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
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
			fmt.Println("1 Plume de Corbeau et 1 Cuir de Sanglier")
			fmt.Println("2 = Craft une tunique de l'aventurier ")
			fmt.Println("2 Fourrure de Loup et 1 Peau de Troll")
			fmt.Println("3 = Craft des bottes de l'aventurier ")
			fmt.Println("1 Fourrure de Loup et 1 Cuir de Sanglier")
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
							fmt.Println("Vous venez de dépenser 5 pièces d'or", char1.Argent)
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
							fmt.Println("Vous venez de dépenser 5 pièces d'or", char1.Argent)
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
		case "5":
			char1.ShowEquipement()
			fmt.Println("1 = Équiper un chapeau de l'aventurier")
			fmt.Println("2 = Équiper une tunique de l'aventurier")
			fmt.Println("3 = Équiper des bottes de l'aventurier")
			fmt.Println("4 = Quitter le menu de l'équipement")
			scanner.Scan() // l'utilisateur input dans la console
			e := scanner.Text()
			switch e {
			case "1":
				char1.RemoveItem(char1.Equipement.Tete)
				char1.AddItem(chapeauAventurier)
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println(char1.Inventaire)
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

			case "2":
				char1.RemoveItem(char1.Equipement.Torse)
				char1.AddItem(tuniqueAventurier)
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println(char1.Inventaire)
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

			case "3":
				char1.RemoveItem(char1.Equipement.Bottes)
				char1.AddItem(bottesAventurier)
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println(char1.Inventaire)
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			case "4":
				break
			}
		case "6":
			TrainingFight(char1, char3, char2)
			fmt.Println("1 = Quitter l'entrainement")
			scanner.Scan() // l'utilisateur input dans la console
			f := scanner.Text()
			switch f {
			case "1":
				break
			}

		case "7":
			os.Exit(2)
		}
	}
}
