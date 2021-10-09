package main

import (
	"bufio"
	"fmt"
	"os"
)

func (char *Marchand) Init(Nom string, Inventaire []string, Argent int) { 
	char.Nom = Nom
	char.Inventaire = Inventaire
	char.Argent = Argent
}

func (char2 *Marchand) DisplayInfoMarchand() { // Rappelle les informations du marchand
	fmt.Println(char2.Nom)
	fmt.Println(char2.Inventaire)
	fmt.Println(char2.Argent)
}

func (char2 *Marchand) Inventaire2() { // Inventaire du marchand
	for i := 0; i < len(char2.Inventaire); i++ {
		fmt.Println(char2.Inventaire[i])
	}
}

func (char2 *Marchand) returnMarchand(char1 *personnage) { // Permet de ne pas quitter l'interface d'achat du marchand
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("Que voulez-vous acheter ?")
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("1 = Acheter une potion de vie (3 pièces d'or)")
	fmt.Println("2 = Acheter une potion de poison (6 pièces d'or)")
	fmt.Println("3 = Acheter une potion de mana (5 pièces d'or)")
	fmt.Println("4 = Acheter Livre de sort : Boule de feu (25 pièces d'or)")
	fmt.Println("5 = Acheter une Fourrure de Loup (4 pièces d'or)")
	fmt.Println("6 = Acheter une Peau de Troll (7 pièces d'or)")
	fmt.Println("7 = Acheter une Cuir de Sanglier (3 pièces d'or)")
	fmt.Println("8 = Acheter une Plume de Corbeau (6 pièces d'or)")
	fmt.Println("9 = Quitter l'inventaire du marchand")
	scanner.Scan() // l'utilisateur input dans la console
	b := scanner.Text()
	switch b {
	case "1": // Acheter une potion de vie
		if len(char1.Inventaire) < 10 { // Vérifie si inventaire pas plein
			if char1.Argent >= 3 {
				char1.Inventaire = AddInventory(char1.Inventaire, "Potion de vie") // Ajoute l'objet à l'inventaire
				char1.Argent -= 3
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous venez d'acheter une potion de vie et vous avez dépensé 3 pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Il vous reste :", char1.Argent, "pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println(char1.Inventaire)// Renvoi l'inventaire
				char2.returnMarchand(char1) // Permet de rester dans le menu du marchand
			} else {
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous êtes trop pauvre pour pouvoir acheter un item")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Il vous reste :", char1.Argent, "pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char2.returnMarchand(char1)// Permet de rester dans le menu du marchand
			}
		} else {
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println(char1.Inventaire)// Renvoi l'inventaire
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			char2.returnMarchand(char1)// Permet de rester dans le menu du marchand
		}
	case "2": // Acheter une potion de poison
		if len(char1.Inventaire) < 10 {
			if char1.Argent >= 6 {
				char1.Inventaire = AddInventory(char1.Inventaire, "Potion de poison")
				char1.Argent -= 6
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous venez d'acheter une potion de poison et vous avez dépensé 6 pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Il vous reste :", char1.Argent, "pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char2.returnMarchand(char1)// Permet de rester dans le menu du marchand
			} else {
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous êtes trop pauvre pour pouvoir acheter un item")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Il vous reste :", char1.Argent, "pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char2.returnMarchand(char1)// Permet de rester dans le menu du marchand

			}
		} else {
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println(char1.Inventaire)// Renvoi l'inventaire
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			char2.returnMarchand(char1)// Permet de rester dans le menu du marchand

		}
	case "3": // Acheter le sort boule de feu
		if len(char1.Inventaire) < 10 {
			if char1.Argent >= 5 {
				char1.Inventaire = AddInventory(char1.Inventaire, potionDeMana)
				char1.Argent -= 5
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous venez d'acheter une potion de mana et vous avez dépensé 5 pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Il vous reste :", char1.Argent, "pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println(char1.Inventaire)
				char2.returnMarchand(char1)
			} else {
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous êtes trop pauvre pour pouvoir acheter un item")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Il vous reste :", char1.Argent, "pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char2.returnMarchand(char1)
			}
		} else {
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println(char1.Inventaire)
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			char2.returnMarchand(char1)
		}
	case "4":
		if len(char1.Inventaire) < 10 {
			if char1.Argent >= 25 {
				char1.Inventaire = AddInventory(char1.Inventaire, "Livre de sort : Boule de feu")
				char1.Argent -= 25
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous venez d'acheter le Livre de sort : Boule de feu et avez dépensé 25 pièces d'or")
				fmt.Println("Il vous reste :", char1.Argent, "pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println(char1.Inventaire)// Renvoi l'inventaire
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char2.returnMarchand(char1)// Permet de rester dans le menu du marchand
			} else {
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous êtes trop pauvre pour pouvoir acheter un item")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Il vous reste :", char1.Argent, "pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char2.returnMarchand(char1)// Permet de rester dans le menu du marchand
			}
		} else {
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println(char1.Inventaire)// Renvoi l'inventaire
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			char2.returnMarchand(char1)// Permet de rester dans le menu du marchand
		}
<<<<<<< HEAD
	case "5":
=======
	case "4": // Acheter une fourrure de loup
>>>>>>> 6638a5300c89dd04d9f941db4ff07c00e6192fe4
		if len(char1.Inventaire) < 10 {
			if char1.Argent >= 4 {
				char1.Inventaire = AddInventory(char1.Inventaire, "Fourrure de Loup")
				char1.Argent -= 4
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous venez d'acheter une Fourrure de Loup et avez dépensé 4 pièces d'or")
				fmt.Println("Il vous reste :", char1.Argent, "pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println(char1.Inventaire)// Renvoi l'inventaire
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char2.returnMarchand(char1) // Permet de rester dans le menu du marchand
			} else {
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous êtes trop pauvre pour pouvoir acheter un item")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Il vous reste :", char1.Argent, "pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char2.returnMarchand(char1) // Permet de rester dans le menu du marchand
			}
		} else {
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println(char1.Inventaire)// Renvoi l'inventaire
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			char2.returnMarchand(char1) // Permet de rester dans le menu du marchand
		}
<<<<<<< HEAD
	case "6":
=======
	case "5": // Acheter une peau de troll
>>>>>>> 6638a5300c89dd04d9f941db4ff07c00e6192fe4
		if len(char1.Inventaire) < 10 {
			if char1.Argent >= 7 {
				char1.Inventaire = AddInventory(char1.Inventaire, "Peau de Troll")
				char1.Argent -= 7
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous venez d'acheter une Peau de Troll et avez dépensé 7 pièces d'or")
				fmt.Println("Il vous reste :", char1.Argent, "pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println(char1.Inventaire)// Renvoi l'inventaire
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char2.returnMarchand(char1) // Permet de rester dans le menu du marchand
			} else {
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous êtes trop pauvre pour pouvoir acheter un item")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Il vous reste :", char1.Argent, "pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char2.returnMarchand(char1) // Permet de rester dans le menu du marchand
			}
		} else {
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println(char1.Inventaire)// Renvoi l'inventaire
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			char2.returnMarchand(char1) // Permet de rester dans le menu du marchand
		}
	case "7":
		if len(char1.Inventaire) < 10 {
			if char1.Argent >= 3 {
				char1.Inventaire = AddInventory(char1.Inventaire, "Cuir de Sanglier")
				char1.Argent -= 3
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous venez d'acheter un Cuir de Sanglier et avez dépensé 3 pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Il vous reste :", char1.Argent, "pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println(char1.Inventaire)// Renvoi l'inventaire
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char2.returnMarchand(char1)// Permet de rester dans le menu du marchand
			} else {
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous êtes trop pauvre pour pouvoir acheter un item")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Il vous reste :", char1.Argent, "pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char2.returnMarchand(char1)// Permet de rester dans le menu du marchand
			}
		} else {
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println(char1.Inventaire)// Renvoi l'inventaire
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			char2.returnMarchand(char1)// Permet de rester dans le menu du marchand
		}
<<<<<<< HEAD
	case "8":
=======
	case "7": // Acheter une plume de corbeau
>>>>>>> 6638a5300c89dd04d9f941db4ff07c00e6192fe4
		if len(char1.Inventaire) < 10 {
			if char1.Argent >= 6 {
				char1.Inventaire = AddInventory(char1.Inventaire, "Plume de Corbeau")
				char1.Argent -= 6
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous venez d'acheter une Plume de Corbeau et vous avez dépensé 6 pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Il vous reste :", char1.Argent, "pièces")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char2.returnMarchand(char1)// Permet de rester dans le menu du marchand
			} else {
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous êtes trop pauvre pour pouvoir acheter un item")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Il vous reste :", char1.Argent, "pièces d'or")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char2.returnMarchand(char1)// Permet de rester dans le menu du marchand
			}
		} else {
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Vous n'avez pas assez de place dans votre inventaire")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println(char1.Inventaire) // Renvoi l'inventaire
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			char2.returnMarchand(char1)// Permet de rester dans le menu du marchand
		}
<<<<<<< HEAD
	case "9":
=======
	case "8": // Partir du marchand
>>>>>>> 6638a5300c89dd04d9f941db4ff07c00e6192fe4
		break
	}
}
