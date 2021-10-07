package main

type personnage struct {
	Nom                  string
	Classe               string
	Niveau               int
	Point_de_vie_max     int
	Point_de_vie_restant int
	Point_d_attaque      int
	Skill                []string
	Inventaire           []string
	Argent               int
	Equipement           equipement
}

type Marchand struct {
	Nom        string
	Inventaire []string
	Argent     int
}

type equipement struct {
	Tete   string
	Torse  string
	Bottes string
}

type monstre struct {
	Nom                  string
	Point_de_vie_max     int
	Point_de_vie_restant int
	Point_d_attaque      int
}

func main() {
	var char3 monstre
	var char2 Marchand
	var char1 personnage
	char1.Init("Lunasphys", "Archer", 20, 100, 70, 5, []string{"Coup de Poing"}, []string{"Potion de vie", "Potion de vie", "Potion de vie", "Potion de poison"}, 100)
	char2.Init("Marchand", []string{"Potion de vie", "épée de fer", "Potion de poison", "Livre de sort : Boule de feu", "Fourrure de loup", "Peau de troll", "Cuir de Sanglier", "Plume de Corbeau"}, 1000000)
	char3.Init("Gobelin d'entrainement", 40, 40, 5)
	//char1.TakePot()
	char1.menu(&char2, &char3)
}
