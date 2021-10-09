package main

import "fmt"

func ManaCoupdePoing(char1 *personnage, char3 *monstre, char2 *Marchand) {

	if char1.Point_de_mana_restant < 5 {
		fmt.Println("Vous n'avez pas le mana requis pour lancez coup de poing")
		fmt.Println("Il vous reste seulement", char1.Point_de_mana_restant, "/", char1.Point_de_mana_max, "points de mana")
		count--
		CharTurn(char1, char3, char2)
	}
	if char1.Point_de_mana_restant > 4 {
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println("Vous avez décidé d'utiliser Coup de poing")
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		char1.Point_de_mana_restant -= 5
		fmt.Println("Vous avez dépensez 5 points de mana")
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println("Il vous reste", char1.Point_de_mana_restant, "/", char1.Point_de_mana_max, "points de mana")
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		char3.Point_de_vie_restant -= 8
		fmt.Println(char1.Nom, "a infligé 8 points de dégats à", (char3.Nom))
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println("Il reste au", char3.Nom, char3.Point_de_vie_restant, "/", char3.Point_de_vie_max, "PV restants")
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	}

}

func ManaBouledeFeu(char1 *personnage, char3 *monstre, char2 *Marchand) {

	if char1.Point_de_mana_restant < 5 {
		fmt.Println("Il vous reste seulement", char1.Point_de_mana_restant, "/", char1.Point_de_mana_max, "points de mana")
		fmt.Println("Vous n'avez pas le mana requis pour lancez boule de feu")
		count--
		CharTurn(char1, char3, char2)
	}
	if Contains(char1.Skill, "Boule de feu") {
		char3.Point_de_vie_restant -= 18
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		char1.Point_de_mana_restant -= 10
		fmt.Println("Vous avez dépensez 10 points de mana")
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println("Il vous reste", char1.Point_de_mana_restant, "/", char1.Point_de_mana_max, "points de mana")
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println(char1.Nom, " a infligé 18 points de dégats à", (char3.Nom))
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println("Il reste au", char3.Nom, char3.Point_de_vie_restant, "/", char3.Point_de_vie_max, "PV restants")
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

	} else {
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println("Vous n'avez pas appris ce sort")
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		count--
		CharTurn(char1, char3, char2)
	}
}
