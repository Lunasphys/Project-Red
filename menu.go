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
			menuCraft(char1)
		case "5":
			menuEquipement(char1)
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
