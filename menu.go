package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Clear() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}
func SlowPrint(str ...string) {
	for _, strpart := range str {
		for _, char := range strpart {
			fmt.Print(string(char))
			time.Sleep(40_000_000 * time.Nanosecond)
		}
	}
}
func (char1 *personnage) start(char2 *Marchand, char3 *monstre, char4 *personnage) {
	SlowPrint("Bonjour à vous jeune aventurier.\n", "Bienvenue dans cette démo jouable qui représente notre premier projet de l année.\n", "Nous sommes heureux de partager cette expérience avec vous et vous montrer nos balbutiements en langage informatique ! \n", "Vous assistez à nos premiers pas dans ce monde, quel honneur.\n Après de nombreuses heures à se creuser la tête, à bloquer devant des messages d’erreurs et à courir après le temps, voici enfin la première ébauche « jouable » de notre projet. \n Merci d’avoir pris le temps de l ouvrir, en espérant que vous passerez un bon moment malgré le manque de contenu. \n", "\n FORCE ET HONNEUR \n")
	fmt.Println("1 = Commencer l'aventure ?")
	fmt.Println("2 = Non, j'ai une vie")
	// créer une var scanner qui va lire ce que l'utilisateur va écrire
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan() // l'utilisateur input dans la console

	// lis ce que l'utilisateur a écrit
	o := scanner.Text()
	switch o {
	case "1":
		Clear()
		char1.menu(char2, char3, char4)
	case "2":
		os.Exit(2)
	}
}

func (char1 *personnage) menu(char2 *Marchand, char3 *monstre, char4 *personnage) { // Le menu du jeu
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("----------------------MENU-------------------------")
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	SlowPrint("Que souhaitez-vous faire ?\n")
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	SlowPrint("0 = Afficher les informations du personnage\n")
	fmt.Println()
	SlowPrint("1 = Acceder au menu de l'inventaire\n")
	fmt.Println()
	SlowPrint("2 = Marchand\n")
	fmt.Println()
	SlowPrint("3 = Acceder a la liste de sorts\n")
	fmt.Println()
	SlowPrint("4 = Acceder au menu du forgeron\n")
	fmt.Println()
	SlowPrint("5 = Acceder a l'équipement\n")
	fmt.Println()
	SlowPrint("6 = Engager un combat d'entrainement\n")
	fmt.Println()
	SlowPrint("7 = Qui sont-ils?\n")
	fmt.Println()
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
		char1.ReturnInventaire(char3)
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
		TrainingFight(char1, char3, char2, char4) // Permet de lancer un combat contre le gobelin d'entraineme    *
	case "7":
		fmt.Println("Dans la partie 2 : Le premier artiste caché est ABBA")
		fmt.Println("Dans la partie 3 : le 2ème artiste caché est le réalisateur Spielberg")
	case "8":
		os.Exit(2)
	}
}
