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

func AddSkill(sorts []string, sort string) []string {
	for _, letter := range sorts {
		if letter == sort {
			fmt.Println("dsl t'a déja", sort)
		}
	}
	return append(sorts, sort)
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
	for _, letter := range char1.Skill {
		if letter == ("Boule de feu") {
			fmt.Println("dsl t'a déja", "Boule de feu")

		} else {
			char1.Skill = append(char1.Skill, "Boule de feu")
		}
	}
	fmt.Println(char1.Skill)
}
