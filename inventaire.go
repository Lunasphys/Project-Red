package main

import (
	"bufio"
	"fmt"
	"os"
)

func (char1 *personnage) AccessInventory() { // Permet d'accéder à l'inventaire
	if len(char1.Inventaire) == 0 { // S'il n'y a rien dans l'inventaire, renvoi au joueur que son inventaire est vide
		fmt.Println("Votre inventaire est vide")
	}
	for i := 0; i < len(char1.Inventaire); i++ { // Permet de compter le nombre d'objet présent dans l'inventaire
		fmt.Println(char1.Inventaire[i])
	}
}

func RemoveInventory(tab []string, s string) []string { // Permet de retirer des objets de l'inventaire
	for index, val := range tab {
		if val == s {
			tab[index] = tab[len(tab)-1]
			return tab[:len(tab)-1]
		}
	}
	return tab
}

func (char *personnage) AddInventory(tab []string, s string) []string { // Permet d'ajouter des objets dans l'inventaire
	if len(tab) >= char.Tailleinventairemax {
		fmt.Println("Vous n'avez plus de place dans votre inventaire")
		return tab
	}
	return append(tab, s)
}

func (char1 *personnage) UpgradeInventorySlot() {
	if char1.Tailleinventairemax == 40 {
		fmt.Println("Vous ne pouvez plus augmentez votre Inventaire")
	}
	if !Contains(char1.Inventaire, "Sac a patate") {
		fmt.Println("Vous n'avez pas de Sac a patate")
	}
	if Contains(char1.Inventaire, "Sac a patate") && len(char1.Inventaire)-1 < char1.Tailleinventairemax && char1.Tailleinventairemax < 40 {
		char1.Tailleinventairemax += 10
		char1.Inventaire = RemoveInventory(char1.Inventaire, "Sac a patate")
		fmt.Println("Vous avez augmenté votre inventaire de 10 places")

	}
}

func (char1 *personnage) ReturnInventaire(char3 *monstre) {
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("------------------INVENTAIRE---------------------")
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	SlowPrint("Que souhaitez-vous faire ?\n")
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	char1.AccessInventory() // Accès à l'inventaire et ses différentes actions
	fmt.Println("1 = Choisissez une potion de vie")
	fmt.Println()
	fmt.Println("2 = Choisissez une potion de poison")
	fmt.Println()
	fmt.Println("3 = Choisissez une potion de mana")
	fmt.Println()
	fmt.Println("4 = Choisissez un sac a patate")
	fmt.Println()
	fmt.Println("5 = Apprendre le sort : Boule de feu")
	fmt.Println()
	fmt.Println("6 = Ne rien choisir et quitter")
	// créer une var scanner qui va lire ce que l'utilisateur va écrire
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // l'utilisateur input dans la console
	n := scanner.Text()
	switch n {
	case "1":
		char1.TakePot() // Permet d'utiliser une potion de vie
		char1.ReturnInventaire(char3)
	case "2":
		char1.PoisonPot() // Permet d'utiliser une potion de poison (pourquoi pas)
		char1.Dead(char3)
		char1.ReturnInventaire(char3)
	case "3":
		char1.TakeManaPot() // Permet d'utiliser une potion de mana
		char1.ReturnInventaire(char3)
	case "4":
		char1.UpgradeInventorySlot() // Permet d'augmenter son inventaire de 10 slots
		char1.ReturnInventaire(char3)
	case "5":
		char1.LearnSkill()       // Permet d'apprendre un sort
		fmt.Println(char1.Skill) // Renvoie liste de sorts
		char1.ReturnInventaire(char3)
	case "6":
		break
	}

}
