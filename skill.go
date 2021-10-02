package main

import (
	"fmt"
)

func AddSkill(tab []string, s string) []string {
	return append(tab, s)
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

func (char1 *personnage) SpellBook(s string) {
	for _, sort := range char1.Skill {
		if sort != s {
			char1.Skill = AddSkill(char1.Skill, s)
			fmt.Println("Vous venez d'apprendre le sort :", s)
		} else {
			fmt.Println("Vous connaissez déjà le sort :", s)
		}
	}
}
