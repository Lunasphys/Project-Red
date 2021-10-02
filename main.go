package main

type personnage struct {
	Nom                  string
	Classe               string
	Niveau               int
	Point_de_vie_max     int
	Point_de_vie_restant int
	Skill                []string
	Inventaire           []string
	Argent               int
}

type Marchand struct {
	Nom        string
	Inventaire []string
	Argent     int
}

func main() {
	var char2 Marchand
	var char1 personnage
	char1.Init("Lunasphys", "Archer", 68, 1350, 1150, []string{"coup de poing"}, []string{"Potion de vie", "Potion de vie", "Potion de vie", "Potion de poison"}, 100)
	char2.Init("Marchand", []string{"Potion de vie", "épée de fer", "Potion de poison", "Boule de feu"}, 1000000)
	//char1.TakePot()
	char1.menu(&Marchand{})
}
