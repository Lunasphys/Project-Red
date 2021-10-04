package main

import (
	"fmt"
)

func (char1 *personnage) AccessInventory() {
	if len(char1.Inventaire) == 0 {
		fmt.Println("Votre inventaire est vide")
	}
	for i := 0; i < len(char1.Inventaire); i++ {
		fmt.Println(char1.Inventaire[i])
	}
}

func RemoveInventory(tab []string, s string) []string {
	for index, val := range tab {
		if val == s {
			tab[index] = tab[len(tab)-1]
			return tab[:len(tab)-1]
		}
	}
	return tab
}

func AddInventory(tab []string, s string) []string {
	if len(tab)-1 >= 9 {
		fmt.Println("Vous n'avez plus de place dans votre inventaire")
		return tab
	}
	return append(tab, s)
}

