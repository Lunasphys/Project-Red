package main

import (
	"bufio"
	"fmt"
	"os"
)

var count int = 1 // Le compteur de tour commence toujours à 1
func TrainingFight(char1 *personnage, char3 *monstre, char2 *Marchand, char4 *personnage) { // Initialisation combat d'entrainement
	fmt.Println(char1.Nom, " engage le combat d'entrainement")
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("A vos armes !!!!")
	for { //
		// Condition de fin de combat
		if char1.Point_de_vie_restant > 0 || char3.Point_de_vie_restant > 0 {
			fmt.Println("======== Tour ", count, " ========") // Initialisation nombre de tours
		}
		if char1.Initiative > char3.Initiative {
			CharTurn(char1, char3, char2, char4)
		}
		if char3.Initiative > char1.Initiative {
			GoblinPattern(char1, char3, char2, char4)
		}
	}
}

func GoblinPattern(char1 *personnage, char3 *monstre, char2 *Marchand, char4 *personnage) { // Tour du gobelin
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
				char1 = char4
				char1.menu(char2, char3, char4)
			} else { // Print des tours
				fmt.Println("======== Tour ", count, " ========")
				CharTurn(char1, char3, char2, char4)
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
				char1 = char4
				char1.menu(char2, char3, char4)
			} else {
				fmt.Println("======== Tour ", count, " ========") // Affichage des tours
				CharTurn(char1, char3, char2, char4)
			}
		}
	}
}

func CharTurn(char1 *personnage, char3 *monstre, char2 *Marchand, char4 *personnage) { // Tour du personnage principal
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
			fmt.Println("4 = Quittez le menu d'attaque")
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
					char3.Dead2()
					char1 = char4
					count = 1
					char1.menu(char2, char3, char4) // Renvoi au menu après la mort du gobelin
				} else {
					GoblinPattern(char1, char3, char2, char4) // Sinon c'est au tour du gobelin de jouer
				}
			case "2":
				ManaCoupdePoing(char1, char3, char2, char4)
				if char3.Point_de_vie_restant <= 0 { // Si le gobelin meurt le compteur de tour revient à 0 et l'utilisateur retourne au menu
					char1 = char4
					char3.Dead2()
					count = 1
					char1.menu(char2, char3, char4)
				} else {
					GoblinPattern(char1, char3, char2, char4) // Sinon c'est au tour du gobelin de jouer
				}

			case "3":
				ManaBouledeFeu(char1, char3, char2, char4)
				if char3.Point_de_vie_restant <= 0 { // Si gobelin mort,  renvoi au menu et réinitialisation du compteur de tour
					char3.Dead2()
					char1 = char4
					count = 1
					char1.menu(char2, char3, char4)
				} else { // Sinon, c'est au tour du gobelin de jouer
					GoblinPattern(char1, char3, char2, char4)
				}

			case "4":
				break
			}
		case "2":
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Que souhaitez-vous faire ?")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			char1.AccessInventory() // Accès à l'inventaire et ses diverses possibilités
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
					CharTurn(char1, char3, char2, char4)
				}
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				if Contains(char1.Inventaire, potionDeVie) { // Vérifie si potion de vie est présent dans l'inventaire
					char1.TakePot()                           // Permet d'uiliser cette potion
					GoblinPattern(char1, char3, char2, char4) // Si cette action a pu se dérouler, alors c'est au tour du gobelin d'attaquer
				} else { // Si potion de vie n'est pas présent dans l'inventaire, renvoi un message au joueur
					fmt.Println("Vous n'avez pas de potion de vie")
				}
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			case "2":
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char1.PoisonPot()                      // Permet d'utiliser une potion de poison sur soi même
				fmt.Println("Vous avez un QI négatif") // Condescendance des développeurs parce qu'on est trop drôle
				if char3.Point_de_vie_restant <= 0 {   // Si goeblin meurt, le compteur de tour revient à 1 et l'utilisateur se retrouve au menu
					char3.Dead2()
					char1 = char4
					count = 1
					char1.menu(char2, char3, char4)
				} else {
					GoblinPattern(char1, char3, char2, char4) // Sinon c'est au tour du gobelin d'attaquer
				}
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			case "3":
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				if char1.Point_de_vie_max == char1.Point_de_vie_restant {
					fmt.Println("Vous ne pouvez pas utiliser de potion")
					count--
					CharTurn(char1, char3, char2, char4)
				}
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				if Contains(char1.Inventaire, potionDeMana) {
					char1.TakeManaPot()
					GoblinPattern(char1, char3, char2, char4)
				} else {
					fmt.Println("Vous n'avez pas de potion de mana")
				}
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			case "4":
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				if Contains(char1.Inventaire, potionDeVie) { // Si l'inventaire contient  bien une potion de vie
					char1.ThrowLifePot(char3)                 // Alors il pourra la jeter sur le gobelin
					GoblinPattern(char1, char3, char2, char4) // Tour du gobelin d'attaquer
				} else {
					fmt.Println("Vous n'avez pas de potion de vie")
				}
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			case "5":
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				if Contains(char1.Inventaire, potionDePoison) { // Si l'inventaire contient bien une potion de poison
					char1.ThrowPoisonPot(char3)          // Si ok, alors action d'envoi de poison sur le gobelin
					if char3.Point_de_vie_restant <= 0 { // Si le gobelin meurt après cette action, le compteur est remit à 1 et renvoi l'utilisateur au menu
						char3.Dead2()
						char1 = char4
						count = 1
						char1.menu(char2, char3, char4)
					} else {
						GoblinPattern(char1, char3, char2, char4) // S'il ne meurt pas, c'est au gobelin d'attaquer
					}
				}
				if !Contains(char1.Inventaire, potionDePoison) { // Si l'inventaire ne contient pas de potion de poison
					fmt.Println("Vous n'avez pas de potion de poison") // Print à l'uilisateur qu'il n'en possède pas
					count--                                            // Décompte 1 tour
					CharTurn(char1, char3, char2, char4)               // Permet à l'utilisateur d'effectuer une nouvelle action
				}
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			case "6":
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				char1.LearnSkill() // Fonction qui permet d'apprendre un sort au personnage principal
				fmt.Println("Sorts appris :")
				fmt.Println(char1.Skill) // Liste des sorts connus par le personnage principal
				fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				if Contains(char1.Skill, bouleDeFeu) { // Si la liste de sort contient bien boule de feu
					fmt.Println(char1.Nom, "a appris Boule de feu") // Indique à l'utilisateur qu'il a bien appris ce soir
					GoblinPattern(char1, char3, char2, char4)       // Renvoi au tour du gobelin
					fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				} else { // Sinon, indique au joueur qu'il n'a pas acheté de livre et lui indique comment obtenir cet item
					fmt.Println("Vous n'avez pas acheté le Livre de sort : Boule de feu au marchand")
					fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
				}
			case "7":
				break
			}
		case "3": // Reinitialisation des caractéristiques de base du gobelin et du joueur ainsi que du compteur de tour. Renvoi l'utilisateur au menu
			char3.Init("Gobelin d'entrainement", 40, 40, 5, 40)
			char1 = char4
			char1.menu(char2, char3, char4)
		}
	}
}
