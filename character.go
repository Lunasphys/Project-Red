package main

import "fmt"

func (char *personnage) Init(Nom string, Classe string, Niveau int, Point_de_vie_max int, Point_de_vie_restant int, Skill []string, Inventaire []string, Argent int) {
	char.Nom = Nom
	char.Classe = Classe
	char.Niveau = Niveau
	char.Point_de_vie_max = Point_de_vie_max
	char.Point_de_vie_restant = Point_de_vie_restant
	char.Skill = Skill
	char.Inventaire = Inventaire
	char.Argent = Argent
}

func (char1 *personnage) DisplayInfo() {
	fmt.Println(char1.Nom)
	fmt.Println(char1.Classe)
	fmt.Println("Vous etes level", ":", char1.Niveau)
	fmt.Println("PV max", ":", char1.Point_de_vie_max)
	fmt.Println("PV restant", ":", char1.Point_de_vie_restant)
	fmt.Println("Liste des sorts", ":", char1.Skill)
	fmt.Println("Vous détenez", ":", char1.Inventaire)
	fmt.Println("Vous avez", char1.Argent, "Pièce d'or")
}

func (char1 *personnage) Dead() {
	for char1.Point_de_vie_restant == 0 {
		fmt.Println("Vous etes mort")
		char1.Point_de_vie_restant += (char1.Point_de_vie_max / 2)
		fmt.Println("Vous etes revenue a la vie avec 50% de vos point de vie restant")
		break
	}
}
