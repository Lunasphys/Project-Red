package main

import (
	"bufio"
	"fmt"
	"os"
)

var count int = 1

func TrainingFight(char1 *personnage, char3 *monstre, char2 *Marchand) { // Initialisation combat d'entrainement
	fmt.Println(char1.Nom, " engage le combat d'entrainement")
	fmt.Println()
	fmt.Println("A vos armes !!!!")
	for { //
		// Condition de fin de combat
		if char1.Point_de_vie_restant > 0 || char3.Point_de_vie_restant > 0 {
			fmt.Println("======== Tour ", count, " ========") // Initialisation nombre de tours
			CharTurn(char1, char3, char2)
		}
	}
}

func GoblinPattern(char1 *personnage, char3 *monstre, char2 *Marchand) {

	for {
		if count%3 == 0 {
			char1.Point_de_vie_restant -= (char3.Point_d_attaque * 2)
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Le gobelin a infligé", (char3.Point_d_attaque * 2), "points de dégats à", (char1.Nom))
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Il vous reste :", char1.Point_de_vie_restant, "/", char1.Point_de_vie_max, "PV restants")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			CharTurn(char1, char3, char2)

		} else {
			char1.Point_de_vie_restant -= char3.Point_d_attaque
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Le gobelin a infligé", (char3.Point_d_attaque), "points de dégats à", (char1.Nom))
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Il reste au", char1.Nom, char1.Point_de_vie_restant, "/", char1.Point_de_vie_max, "PV restants")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			if char1.Point_de_vie_restant <= 0 {
				char1.Dead()
				os.Exit(1)
			} else {
				fmt.Println("======== Tour ", count, " ========")
				CharTurn(char1, char3, char2)
			}
		}
	}
}

func CharTurn(char1 *personnage, char3 *monstre, char2 *Marchand) {
	count++
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("Que souhaitez-vous faire ?")
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	for {
		fmt.Println("1 = Attaquer")
		fmt.Println("2 = Accéder à votre inventaire")
		fmt.Println("3 = Prendre la fuite")
		// créer une var scanner qui va lire ce que l'utilisateur va écrire
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // l'utilisateur input dans la console

		// lis ce que l'utilisateur a écrit
		v := scanner.Text()
		switch v {
		case "1":
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Choisissez une attaque")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("1 = Attaque simple")
			fmt.Println("2 = Coup de poing")
			fmt.Println("3 = Boule de feu")
			scanner.Scan() // l'utilisateur input dans la console
			m := scanner.Text()
			switch m {
			case "1":
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous avez décidé de réaliser une attaque simple")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char3.Point_de_vie_restant -= char1.Point_d_attaque
				fmt.Println(char1.Nom, " a infligé", (char1.Point_d_attaque), "points de dégats à", (char3.Nom))
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Il reste au", char3.Nom, char3.Point_de_vie_restant, "/", char3.Point_de_vie_max, "PV restants")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				if char3.Point_de_vie_restant <= 0 {
					char3.Dead2(char1)
					count = 1
					char1.menu(char2, char3)
				} else {
					GoblinPattern(char1, char3, char2)
				}
			case "2":
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous avez décidé d'utiliser Coup de poing")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char3.Point_de_vie_restant -= 8
				fmt.Println(char1.Nom, " a infligé 8 points de dégats à", (char3.Nom))
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Il reste au", char3.Nom, char3.Point_de_vie_restant, "/", char3.Point_de_vie_max, "PV restants")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				if char3.Point_de_vie_restant <= 0 {
					char3.Dead2(char1)
					count = 1
					char1.menu(char2, char3)
				} else {
					GoblinPattern(char1, char3, char2)
				}
			case "3":
				if Contains(char1.Skill, "Boule de feu") {
					char3.Point_de_vie_restant -= 18
					fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
					fmt.Println(char1.Nom, " a infligé 18 points de dégats à", (char3.Nom))
					fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
					fmt.Println("Il reste au", char3.Nom, char3.Point_de_vie_restant, "/", char3.Point_de_vie_max, "PV restants")
					fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				} else {
					fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
					fmt.Println("Vous n'avez pas appris ce sort")
					fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				}
				if char3.Point_de_vie_restant <= 0 {
					char3.Dead2(char1)
					count = 1
					char1.menu(char2, char3)
				} else {
					GoblinPattern(char1, char3, char2)
				}
		case "2":
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Que souhaitez-vous faire ?")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			char1.AccessInventory()
			fmt.Println("1 = Choisissez une potion de vie")
			fmt.Println("2 = Choisissez une potion de poison")
			fmt.Println("3 = Apprendre le sort : Boule de feu")
			fmt.Println("4 = Ne rien choisir et quitter")
			scanner.Scan() // l'utilisateur input dans la console
			p := scanner.Text()
			switch p {
			case "1":
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char1.TakePot()
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			case "2":
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char1.PoisonPot()
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			case "3":
				GoblinPattern(char1, char3, char2)
			}
		case "3":
			os.Exit(45434)
			}
		}
	}
}
