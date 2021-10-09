package main

import (
	"bufio"
	"fmt"
	"os"
)

var count int = 1 // Le compteur de tour commence toujours à 1

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

func GoblinPattern(char1 *personnage, char3 *monstre, char2 *Marchand) { // Tour du gobelin
	for {
		if count%3 == 0 { // Initie l'attaque critique du gobelin
			char1.Point_de_vie_restant -= (char3.Point_d_attaque * 2)
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Le gobelin a infligé", (char3.Point_d_attaque * 2), "points de dégats à", (char1.Nom))
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Il vous reste :", char1.Point_de_vie_restant, "/", char1.Point_de_vie_max, "PV restants")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			if char1.Point_de_vie_restant <= 0 { // Mort du gobelin et retour au menu
				char1.Dead(char3)
				char1.menu(char2, char3)
			} else { // Print des tours
				fmt.Println("======== Tour ", count, " ========")
				CharTurn(char1, char3, char2)
			}
		} else { // Attaque basique du gobelin
			char1.Point_de_vie_restant -= char3.Point_d_attaque
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Le gobelin a infligé", (char3.Point_d_attaque), "points de dégats à", (char1.Nom))
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Il reste au", char1.Nom, char1.Point_de_vie_restant, "/", char1.Point_de_vie_max, "PV restants")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			if char1.Point_de_vie_restant <= 0 { // Mort du gobelin et retour au menu
				char1.Dead(char3)
				char1.menu(char2, char3)
			} else {
				fmt.Println("======== Tour ", count, " ========") // Affichage des tours
				CharTurn(char1, char3, char2)
			}
		}
	}
}

func CharTurn(char1 *personnage, char3 *monstre, char2 *Marchand) { // Tour du personnage principal
	count++ // Compte les différents tours
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
			scanner.Scan()      // l'utilisateur input dans la console
			m := scanner.Text() // lis ce que l'utilisation a écrit
			switch m {
			case "1":
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Vous avez décidé de réaliser une attaque simple")
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char3.Point_de_vie_restant -= char1.Point_d_attaque
				fmt.Println(char1.Nom, " a infligé", (char1.Point_d_attaque), "points de dégats à", (char3.Nom))
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				fmt.Println("Il reste au", char3.Nom, char3.Point_de_vie_restant, "/", char3.Point_de_vie_max, "PV restants") // Renvoi à l'utilisateur son nombre de PV restants sur
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				if char3.Point_de_vie_restant <= 0 { // Permet la mort du gobelin
					char3.Dead2(char1)
					count = 1
					char1.menu(char2, char3) // Renvoi au menu après la mort du gobelin
				} else {
					GoblinPattern(char1, char3, char2) // Sinon renvoi au tour du gobelin
				}
			case "2":
				ManaCoupdePoing(char1, char3, char2)
				if char3.Point_de_vie_restant <= 0 {
					char3.Dead2(char1)
					count = 1
					char1.menu(char2, char3)
				} else {
					GoblinPattern(char1, char3, char2)
				}
			case "3":
				ManaBouledeFeu(char1, char3, char2)
				if char3.Point_de_vie_restant <= 0 {
					char3.Dead2(char1)
					count = 1
					char1.menu(char2, char3)
				} else {
					GoblinPattern(char1, char3, char2)
				}
			case "4":
				break
			}
		case "2":
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Que souhaitez-vous faire ?")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			char1.AccessInventory()
			fmt.Println("1 = Choisissez une potion de vie")
			fmt.Println("2 = Choisissez une potion de poison")
			fmt.Println("3 = Choisissez une potion de mana")
			fmt.Println("4 = Envoyer une potion de vie sur le Gobelin")
			fmt.Println("5 = Envoyer une potion de poison sur le Gobelin")
			fmt.Println("6 = Apprendre le sort : Boule de feu")
			fmt.Println("7 = Ne rien choisir et quitter")
			scanner.Scan() // l'utilisateur input dans la console
			p := scanner.Text()
			switch p {
			case "1":
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				if char1.Point_de_vie_max == char1.Point_de_vie_restant {
					fmt.Println("Vous ne pouvez pas utiliser de potion")
					count--
					CharTurn(char1, char3, char2)
				}
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				if Contains(char1.Inventaire, potionDeVie) {
					char1.TakePot()
					GoblinPattern(char1, char3, char2)
				} else {
					fmt.Println("Vous n'avez pas de potion de vie")
				}
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			case "2":
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char1.PoisonPot()
				fmt.Println("Vous avez un QI négatif")
				if char3.Point_de_vie_restant <= 0 {
					char3.Dead2(char1)
					count = 1
					char1.menu(char2, char3)
				} else {
					GoblinPattern(char1, char3, char2)
				}
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			case "3":
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				if char1.Point_de_vie_max == char1.Point_de_vie_restant {
					fmt.Println("Vous ne pouvez pas utiliser de potion")
					count--
					CharTurn(char1, char3, char2)
				}
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				if Contains(char1.Inventaire, potionDeMana) {
					char1.TakeManaPot()
					GoblinPattern(char1, char3, char2)
				} else {
					fmt.Println("Vous n'avez pas de potion de mana")
				}
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			case "4":
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				if Contains(char1.Inventaire, potionDeVie) {
					char1.ThrowLifePot(char3)
					GoblinPattern(char1, char3, char2)
				} else {
					fmt.Println("Vous n'avez pas de potion de vie")
				}
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			case "5":
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				if Contains(char1.Inventaire, potionDePoison) {
					char1.ThrowPoisonPot(char3)
					if char3.Point_de_vie_restant <= 0 {
						char3.Dead2(char1)
						count = 1
						char1.menu(char2, char3)
					} else {
						GoblinPattern(char1, char3, char2)
					}
				}
				if !Contains(char1.Inventaire, potionDePoison) {
					fmt.Println("Vous n'avez pas de potion de poison")
					count--
					CharTurn(char1, char3, char2)
				}
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			case "6":
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char1.LearnSkill()
				fmt.Println("Sorts appris :")
				fmt.Println(char1.Skill)
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				if Contains(char1.Skill, bouleDeFeu) {
					fmt.Println(char1.Nom, "a appris Boule de feu")
					GoblinPattern(char1, char3, char2)
					fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				} else {
					fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
					fmt.Println("Vous n'avez pas acheté le Livre de sort : Boule de feu au marchand")
				}
			case "7":
				break
			}
		case "3":
			char3.Init("Gobelin d'entrainement", 40, 40, 5)
			char1.Init("Lunasphys", "Archer", 20, 50, 30, 50, 30, 5, []string{"Coup de Poing"}, []string{"Potion de vie", "Potion de vie", "Potion de vie", "Potion de poison"}, 100)
			count = 1
			char1.menu(char2, char3)
		}
	}
}
