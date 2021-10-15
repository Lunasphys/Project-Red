package main

import (
	"bufio"
	"fmt"
	"os"
)

func (char *equipement) Init(Tete string, Torse string, Bottes string) { // Equipement du personnage
	char.Tete = Tete
	char.Torse = Torse
	char.Bottes = Bottes
}

func (char1 *personnage) ShowEquipement() { // Permet de print la liste de l'équipement équipé
	fmt.Println("Tête :", char1.Equipement.Tete)
	fmt.Println("Torse :", char1.Equipement.Torse)
	fmt.Println("Bottes :", char1.Equipement.Bottes)
}

func (char1 *personnage) EquipHead(s string) string { // Permet d'équiper la tête
	if Contains(char1.Inventaire, s) {
		char1.Equipement.Tete = s
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println("Vous venez d'équiper : ", s)
		fmt.Println(char1.Inventaire)
	}
	if !(Contains(char1.Inventaire, s)) {
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println("Vous n'avez rien à équiper")
		fmt.Println(char1.Inventaire)
	}
	return char1.Equipement.Tete
}

func (char1 *personnage) EquipChest(s string) string { // Permet d'équiper le torse
	if Contains(char1.Inventaire, s) {
		char1.Equipement.Torse = s
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println("Vous venez d'équiper : ", s)
		fmt.Println(char1.Inventaire)
	}
	if !(Contains(char1.Inventaire, s)) {
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println("Vous n'avez rien à équiper")
		fmt.Println(char1.Inventaire)
	}
	return char1.Equipement.Torse
}

func (char1 *personnage) EquipBoots(s string) string { // Permet d'équiper les bottes
	if Contains(char1.Inventaire, s) {
		char1.Equipement.Bottes = s
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println("Vous venez d'équiper : ", s)
		fmt.Println(char1.Inventaire)
	}
	if !(Contains(char1.Inventaire, s)) {
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println("Vous n'avez rien à équiper")
		fmt.Println(char1.Inventaire)
	}
	return char1.Equipement.Bottes
}

func menuEquipement(char1 *personnage) { // Fonction utiliser dans le menu pour équiper les équipements
	scanner := bufio.NewScanner(os.Stdin)
	char1.ShowEquipement()
	fmt.Println("-------------------EQUIPEMENT--------------------")
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	SlowPrint("Que souhaitez vous équiper")
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("1 = Équiper un chapeau de l'aventurier (+10 PV max)")
	fmt.Println()
	fmt.Println("2 = Équiper une tunique de l'aventurier (+25 PV max)")
	fmt.Println()
	fmt.Println("3 = Équiper des bottes de l'aventurier (+15 PV max)")
	fmt.Println()
	fmt.Println("4 = Quitter le menu de l'équipement")
	scanner.Scan() // l'utilisateur input dans la console
	e := scanner.Text()
	switch e {
	case "1":
		if !Contains(char1.Inventaire, chapeauAventurier) {
			char1.EquipHead("")
			break
		}
		char1.RemoveItem(char1.Equipement.Tete) // Enlève équipement tête
		char1.AddItem(chapeauAventurier)        // Ajout item
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println(char1.Inventaire) // Imprime l'inventaire
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

	case "2":
		if !Contains(char1.Inventaire, tuniqueAventurier) { // Si ne contient pas le dit item
			char1.EquipChest("")
			break
		}
		char1.RemoveItem(char1.Equipement.Torse) // Enlever équipement torse
		char1.AddItem(tuniqueAventurier)         // Ajout item dans équipement
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println(char1.Inventaire) // Imprime l'inventaire
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

	case "3":
		if !Contains(char1.Inventaire, bottesAventurier) { // Si ne contient pas le dit item
			char1.EquipBoots("")
			break
		}
		char1.RemoveItem(char1.Equipement.Bottes) // Enlève les bottes
		char1.AddItem(bottesAventurier)           // Ajoute l'item des bottes dans l'équipement
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println(char1.Inventaire)
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	case "4":
		break
	}
}
