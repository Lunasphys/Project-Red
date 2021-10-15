package main

type personnage struct { // Stucture du personnage principal
	Nom                   string
	Classe                string
	Niveau                int
	Point_de_vie_max      int
	Point_de_vie_restant  int
	Point_de_mana_max     int
	Point_de_mana_restant int
	Point_d_attaque       int
	Initiative            int
	Skill                 []string
	Inventaire            []string
	Argent                int
	Equipement            equipement
	Tailleinventairemax   int
}

type Marchand struct { // Informations du marchand
	Nom        string
	Inventaire []string
	Argent     int
}

type equipement struct { // Equipement du joueur
	Tete   string
	Torse  string
	Bottes string
}

type monstre struct { // Informations du Gobelin d'entrainement
	Nom                  string
	Point_de_vie_max     int
	Point_de_vie_restant int
	Point_d_attaque      int
	Initiative           int
}

func main() { // Informations du personnage principal, du marchand et du gobelin d'entrainement
	var char4 personnage
	var char3 monstre
	var char2 Marchand
	var char1 personnage
	char1.Init("Lunasphys", "Archer", 20, 50, 30, 50, 40, 5, 50, []string{"Coup de Poing"}, []string{"Potion de vie", "Potion de vie", "Potion de vie", "Potion de poison"}, 1000, 10)
	char2.Init("Marchand", []string{"Potion de vie", "Potion de poison", "Potion de Mana", "Livre de sort : Boule de feu", "Fourrure de loup", "Peau de troll", "Cuir de Sanglier", "Plume de Corbeau"}, 1000000)
	char3.Init("Gobelin d'entrainement", 40, 40, 5, 60)
	char4.Init("", "", 0, 0, 0, 0, 0, 0, 0, []string{}, []string{}, 0, 0)
	char4 = char1
	char1.start(&char2, &char3, &char4)
	char1.menu(&char2, &char3, &char4)
}
