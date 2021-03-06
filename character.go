package main

import "fmt"

/* Cette fonction permet de gérer les paramètres du personnage principal
 */
func (char *personnage) Init(Nom string, Classe string, Niveau int, Point_de_vie_max int, Point_de_vie_restant int, Point_de_mana_max int, Point_de_mana_restant int, Point_d_attaque int, Initiative int, Skill []string, Inventaire []string, Argent int, Tailleinventairemax int) {
	char.Nom = Nom
	char.Classe = Classe
	char.Niveau = Niveau
	char.Point_de_vie_max = Point_de_vie_max
	char.Point_de_vie_restant = Point_de_vie_restant
	char.Point_de_mana_max = Point_de_mana_max
	char.Point_de_mana_restant = Point_de_mana_restant
	char.Point_d_attaque = Point_d_attaque
	char.Initiative = Initiative
	char.Skill = Skill
	char.Inventaire = Inventaire
	char.Argent = Argent
	char.Tailleinventairemax = Tailleinventairemax
}

func (char1 *personnage) DisplayInfo() { /* Cette fonction permet d'afficher les informations principales de notre personnage lorsque cette fonction est appelée dans le menu*/
	fmt.Println(char1.Nom)
	fmt.Println(char1.Classe)
	fmt.Println("Vous êtes level :", char1.Niveau)
	fmt.Println("PV max :", char1.Point_de_vie_max)
	fmt.Println("PV restant :", char1.Point_de_vie_restant)
	fmt.Println("Mana max :", char1.Point_de_mana_max)
	fmt.Println("Mana restant :", char1.Point_de_mana_restant)
	fmt.Println("Vous avez", char1.Point_d_attaque, "points d'attaque")
	fmt.Println("Vous avez", char1.Initiative, "d'Initiative")
	fmt.Println("Liste des sorts :", char1.Skill)
	fmt.Println("Vous détenez :", char1.Inventaire)
	fmt.Println("Vous avez", char1.Argent, "Pièces d'or")
	fmt.Println("Vous avez", char1.Tailleinventairemax, "place dans votre inventaire")
}

func (char1 *personnage) Dead(char3 *monstre) { /* Permet au perso principal de mourir pendant le combat d'entrainement et de réinitialiser ses stats de base à la fin du combat
	 */
	fmt.Println("Vous etes mort")
	char1.Point_de_vie_restant += (char1.Point_de_vie_max / 2)
	fmt.Println("Vous etes revenue a la vie avec 50% de vos point de vie restant")
	char3.Init("Gobelin d'entrainement", 40, 40, 5, 40)
}

func (char3 *monstre) Dead2() { // Permet au gobelin de mourir et de réinitialiser ses stats de base à la fin du combat
	fmt.Println("Le gobelin est mort")
	char3.Init("Gobelin d'entrainement", 40, 40, 5, 40)
	fmt.Println("Un nouveau Gobelin est disponible")
}

func (char4 *personnage) DisplayInfo1() { /* Cette fonction permet d'afficher les informations principales de notre personnage lorsque cette fonction est appelée dans le menu*/
	fmt.Println(char4.Nom)
	fmt.Println(char4.Classe)
	fmt.Println("Vous êtes level :", char4.Niveau)
	fmt.Println("PV max :", char4.Point_de_vie_max)
	fmt.Println("PV restant :", char4.Point_de_vie_restant)
	fmt.Println("Mana max :", char4.Point_de_mana_max)
	fmt.Println("Mana restant :", char4.Point_de_mana_restant)
	fmt.Println("Vous avez", char4.Point_d_attaque, "points d'attaque")
	fmt.Println("Liste des sorts :", char4.Skill)
	fmt.Println("Vous détenez :", char4.Inventaire)
	fmt.Println("Vous avez", char4.Argent, "Pièces d'or")
}
