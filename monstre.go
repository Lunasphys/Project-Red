package main

import "fmt"

func (char *monstre) Init(Nom string,Point_de_vie_max int, Point_de_vie_restant int, Point_d_attaque int) {
	char.Nom = Nom
	char.Point_de_vie_max = Point_de_vie_max
	char.Point_de_vie_restant = Point_de_vie_restant
	char.Point_d_attaque = Point_d_attaque
}

func (char1 *monstre) DisplayInfo() {
	fmt.Println(char1.Nom)
	fmt.Println(char1.Point_de_vie_max)
	fmt.Println(char1.Point_de_vie_restant)
	fmt.Println(char1.Point_d_attaque)
}

