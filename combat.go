package main

import "fmt"

func TrainingFight(char1 *personnage, char3 *monstre) { // Initialisation combat d'entrainement
	fmt.Println(char1.Nom, " engage le combat d'entrainement")
	fmt.Println()
	fmt.Println("A vos armes !!!!")
	for count := 1; ; count++ { // Condition de fin de combat
		if char1.Point_de_vie_restant > 0 || char3.Point_de_vie_restant > 0 {
			fmt.Println("======== Tour ", count, " ========") // Initialisation nombre de tours
			CharTurn(char1, char3)
			//
			GoblinPattern(char1, char3, count)
		}
		if char3.Point_de_vie_restant >= 0 {
			fmt.Println("Vous avez gagné l'entrainement")
		}
		if char1.Point_de_vie_restant >= 0 {
			fmt.Println("Vous avez perdu l'entrainement")
		}
	}
}

func GoblinPattern(char1 *personnage, char3 *monstre, count int) {
	for count := 1; ; count++ {
		if count%3 == 0 {
			char1.Point_de_vie_restant -= (char3.Point_d_attaque * 2)
			fmt.Println("Le gobelin a infligé", (char3.Point_d_attaque * 2), "points de dégats à", (char1.Nom))
		} else {
			char1.Point_de_vie_restant -= char3.Point_d_attaque
			fmt.Println("Le gobelin a infligé", (char3.Point_d_attaque), "points de dégats à", (char1.Nom))
		}
	}
}

func CharTurn(char1 *personnage, char3 *monstre) {
}
