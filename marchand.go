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
	fmt.Println("9 = Acheter un Sac a patate (30 pièces d'or)")
	fmt.Println("10 = Quitter l'inventaire du marchand")
	scanner.Scan() // l'utilisateur input dans la console
	b := scanner.Text()
	switch b {
	case "1": // Acheter une potion de vie
		char1.Vente(char2, 3, potionDeVie)
	case "2": // Acheter une potion de poison
		char1.Vente(char2, 6, potionDePoison)
	case "3":
		char1.Vente(char2, 5, potionDeMana)
	case "4": // Acheter le sort boule de feu
		char1.Vente(char2, 25, livreDeSortBouleDeFeu)
	case "5": // Acheter une fourrure de loup
		char1.Vente(char2, 4, "Fourrure de Loup")
	case "6": // Acheter une peau de troll
		char1.Vente(char2, 7, "Peau de Troll")
	case "7":
		char1.Vente(char2, 3, "Cuir de Sanglier")
	case "8": // Acheter une plume de corbeau
		char1.Vente(char2, 6, "Plume de Corbeau")
	case "9":
		char1.Vente(char2, 30, "Sac a patate")
	case "10": // Partir du marchand
		break
	}
}

func (char1 *personnage) Vente(char2 *Marchand, arg int, objet string) {

	if len(char1.Inventaire) < char1.Tailleinventairemax { // Vérifie si inventaire pas plein
		if char1.Argent >= arg {
			char1.Inventaire = char1.AddInventory(char1.Inventaire, objet) // Ajoute l'objet à l'inventaire
			char1.Argent -= arg
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Vous venez d'acheter une potion de vie et vous avez dépensé", arg, "pieces d'or")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Il vous reste :", char1.Argent, "pièces d'or")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println(char1.Inventaire) // Renvoi l'inventaire
			char2.returnMarchand(char1)   // Permet de rester dans le menu du marchand
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
		fmt.Println(char1.Inventaire) // Renvoi l'inventaire
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		char2.returnMarchand(char1) // Permet de rester dans le menu du marchand
	}
}
