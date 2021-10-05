package main

import "fmt"

func (char *equipement) Init(Tete string, Torse string, Bottes string) {
	char.Tete = Tete
	char.Torse = Torse
	char.Bottes = Bottes
}

func (char1 *personnage) ShowEquipement() {
	println(char1.Equipement.Tete)
	println(char1.Equipement.Torse)
	println(char1.Equipement.Bottes)
}

func (char1 *personnage) EquipHead(s string) string {
	for _, tete := range char1.Inventaire {
		if tete == s {
			char1.Equipement.Tete = s
			fmt.Println("Vous venez d'équiper : ", s)
			char1.Point_de_vie_max += 10
		}
	}
	if char1.Equipement.Tete != s {
		fmt.Println("Vous n'avez rien à équiper")
	}
	return char1.Equipement.Tete
}

func (char1 *personnage) EquipChest(s string) string {
	for _, tete := range char1.Inventaire {
		if tete == s {
			char1.Equipement.Torse = s
			fmt.Println("Vous venez d'équiper : ", s)
			char1.Point_de_vie_max += 25
		}
	}
	fmt.Println("Vous n'avez rien à équiper")
	return char1.Equipement.Torse
}

func (char1 *personnage) EquipBoots(s string) string {
	for _, tete := range char1.Inventaire {
		if tete == s {
			char1.Equipement.Bottes = s
			fmt.Println("Vous venez d'équiper : ", s)
			char1.Point_de_vie_max += 15
		}
	}
	fmt.Println("Vous n'avez rien à équiper")
	return char1.Equipement.Tete
}
