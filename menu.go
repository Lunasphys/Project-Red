package main

import (
	"bufio"
	"fmt"
	"os"
)

func (char1 *personnage) menu(char2 *Marchand) {
	for {
		fmt.Println("0 = Afficher les informations du personnages")
		fmt.Println("1 = Acceder au menue de l'inventaire")
		fmt.Println("2 = Marchand")
		fmt.Println("3 = Quitter")
		// créer une var scanner qui va lire ce que l'utilisateur va écrire
		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan() // l'utilisateur input dans la console

		// lis ce que l'utilisateur a écrit
		v := scanner.Text()
		switch v {
		case "0":
			char1.DisplayInfo()
		case "1":
			char1.AccessInventory()
			fmt.Println("1 = Choisissez une potion de vie")
			fmt.Println("2 = Choisissez une potion de poison")
			fmt.Println("3 = Ne rien choisir et quitter")
			scanner.Scan() // l'utilisateur input dans la console
			n := scanner.Text()
			switch n {
			case "1":
				char1.TakePot()
			case "2":
				char1.PoisonPot()
			case "3":
				break
			}
		case "2":
			char2.Inventaire2()
			fmt.Println("1 = Acheter une potion de vie")
			fmt.Println("2 = Acheter une potion de poison")
			fmt.Println("3 = Quitter l'inventaire du marchand")
			scanner.Scan() // l'utilisateur input dans la console
			b := scanner.Text()
			switch b {
			case "1":
				char1.Inventaire = AddInventory(char1.Inventaire, "Potion de vie")
				fmt.Println(char1.Inventaire)
			case "2":
				char1.Inventaire = AddInventory(char1.Inventaire, "Potion de poison")
			case "3":
				break
			}
		case "3":
			os.Exit(2)
		}
	}
}
