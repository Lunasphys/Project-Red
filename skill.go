package main

import (
	"fmt"
)

func (char1 *personnage) AccessSkill() {
	if len(char1.Skill) == 0 {
		fmt.Println("Vous n'avez appris aucun sort")
	}
	for i := 0; i < len(char1.Skill); i++ {
		fmt.Println(char1.Skill[i])
	}
}

func AddSkill(sorts []string, s string) []string {
	return append(sorts, s)
}

func Contains(tab []string, s string) bool {
	for _, a := range tab {
		if a == s {
			return true
		}
	}
	return false
}

func RemoveDuplicates(tab []string) []string {
	list := []string{}
	for _, sort := range tab {
		fmt.Println(sort)
		if !Contains(list, sort) {
			list = append(list, sort)
		}
	}
	return list
}

func (char1 *personnage) SpellBook() {
    // fonction qui nous permet d'ajouter ou repertorier les sorts (spell)
	for _, sort := range char1.Skill {
		if sort == "Boule de feu" {
			if (char1.Point_de_vie_restant) < (char1.Point_de_vie_max) {
				char1.Skill = RemoveDuplicates( 
				fmt.Println("Vous avez apris le sort :", sort)
				if char1.Point_de_vie_restant >= char1.Point_de_vie_max {
					char1.Point_de_vie_restant = char1.Point_de_vie_max
					fmt.Println("Vous ne pouvez pas utiliser de potion")
					if objet != "Potion de vie" {
						fmt.Println("Vous n'avez pas de potion")
					}
				}
			}
			break
		}
	}
	char1.AccessS()
	fmt.Println(char1.Point_de_vie_restant)
}
