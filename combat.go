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
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Le gobelin a infligé", (char3.Point_d_attaque * 2), "points de dégats à", (char1.Nom))
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Il vous reste :", char1.Point_de_vie_restant, "/", char1.Point_de_vie_max, "PV restants")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		} else {
			char1.Point_de_vie_restant -= char3.Point_d_attaque
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Le gobelin a infligé", (char3.Point_d_attaque), "points de dégats à", (char1.Nom))
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			fmt.Println("Il reste au", char3.Nom, char3.Point_de_vie_restant, "/", char3.Point_de_vie_max, "PV restants")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		}
	}
}

func CharTurn(char1 *personnage, char3 *monstre) {
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
				char1.TakePot()
                fmt.Println("Vous avez décidé de réaliser une attaque simple")
			case "2":
				char1.PoisonPot()
                fmt.Println("Vous avez décidé d'utiliser Coup de poing")
			case "3":
				char1.LearnSkill()
				fmt.Println(char1.Skill)
			case "4":
				break
			}
			
		case "2":