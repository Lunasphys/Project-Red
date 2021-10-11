package main

import (
	"bufio"
	"fmt"
	"os"
)

func (char1 *personnage) menu(char2 *Marchand, char3 *monstre, char4 *personnage) { // Le menu du jeu
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
			char1.AccessInventory() // Accès à l'inventaire et ses différentes actions
			fmt.Println("1 = Choisissez une potion de vie")
			fmt.Println("2 = Choisissez une potion de poison")
			fmt.Println("3 = Choisissez une potion de mana")
			fmt.Println("4 = Apprendre le sort : Boule de feu")
			fmt.Println("5 = Ne rien choisir et quitter")
			scanner.Scan() // l'utilisateur input dans la console
			n := scanner.Text()
			switch n {
			case "1":
				char1.TakePot() // Permet d'utiliser une potion de vie
			case "2":
				char1.PoisonPot() // Permet d'utiliser une potion de poison (pourquoi pas)
				char1.Dead(char3)
			case "3":
				char1.TakeManaPot() // Permet d'utiliser une potion de mana
			case "4":
				char1.LearnSkill()       // Permet d'apprendre un sort
				fmt.Println(char1.Skill) // Renvoie liste de sorts
			case "5":
				break
			}
		case "2":
			char2.returnMarchand(char1) // Rentre dans le menu du marchand
		case "3":
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Voici votre liste de sorts ?")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			char1.AccessSkill()
		case "4":
			menuCraft(char1) // Menu du forgeron
		case "5":
			menuEquipement(char1) // Menu de l'équipement
		case "6":
			char1.DisplayInfo()
			char3.DisplayInfo()
			TrainingFight(char1, char3, char2, char4) // Permet de lancer un combat contre le gobelin d'entraineme    *
		case "7":
			os.Exit(2)
		}
	}
}
