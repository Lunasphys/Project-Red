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
	if !Contains(sorts, s) {
		sorts = append(sorts, s)
	}
	return sorts
}

func Contains(tab []string, s string) bool {
	for _, a := range tab {
		if a == s {
			return true
		}
	}
	return false
}

func (char1 *personnage) LearnSkill() {
	for _, skill := range char1.Inventaire {
		if skill == "Livre de sort : Boule de feu" {
			char1.Skill = AddSkill(char1.Skill, skill)
			fmt.Println("Vous avez appris le sort Boule de feu")
		}
	}
}
