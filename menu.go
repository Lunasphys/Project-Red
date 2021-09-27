package main

import (
	"bufio"
	"fmt"
	"os"
)

func (char1 *personnage) menu() {
	for {
		fmt.Println("0 = Afficher les informations du personnages")
		fmt.Println("1 = Acceder au menue de l'inventaire")
		fmt.Println("2 = Quitter")
		fmt.Println("3 = Prendre une potion")
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
		case "2":
			os.Exit(2)
		case "3":
			char1.TakePot()
		}
	}
}
