package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func SlowPrint(str ...string) {
	for _, strpart := range str {
		for _, char := range strpart {
			fmt.Print(string(char))
			time.Sleep(40_000_000 * time.Nanosecond)
		}
	}
}

func (char1 *personnage) menu(char2 *Marchand, char3 *monstre, char4 *personnage) { // Le menu du jeu
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	SlowPrint("Que souhaitez-vous faire ?\n")
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	for {
		SlowPrint("0 = Afficher les informations du personnages\n")
		SlowPrint("1 = Acceder au menu de l'inventaire\n")
		SlowPrint("2 = Marchand\n")
		SlowPrint("3 = Acceder a la liste de sorts\n")
		SlowPrint("4 = Acceder au menu du forgeron\n")
		SlowPrint("5 = Acceder a l'équipement\n")
		SlowPrint("6 = Engager un combat d'entrainement\n")
		SlowPrint("7 = Qui sont ils?\n")
		SlowPrint("8 = Quitter\n")
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
			SlowPrint("Que souhaitez-vous faire ?\n")
			fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			char1.AccessInventory() // Accès à l'inventaire et ses différentes actions
			SlowPrint("1 = Choisissez une potion de vie\n")
			SlowPrint("2 = Choisissez une potion de poison\n")
			SlowPrint("3 = Choisissez une potion de mana\n")
			SlowPrint("4 = Choisissez un sac a patate\n")
			SlowPrint("5 = Apprendre le sort : Boule de feu\n")
			SlowPrint("6 = Ne rien choisir et quitter\n")
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
				char1.UpgradeInventorySlot() // Permet d'augmenter son inventaire de 10 slots
			case "5":
				char1.LearnSkill()       // Permet d'apprendre un sort
				fmt.Println(char1.Skill) // Renvoie liste de sorts
			case "6":
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
			fmt.Println("Dans la partie 2 : Le premier artiste caché est ABBA")
			fmt.Println("Dans la partie 3 : le 2ème artiste caché est le réalisateur Spielberg")
		case "8":
			os.Exit(2)
		}
	}
}
