package main

type personnage struct {
	Nom                  string
	Classe               string
	Niveau               int
	Point_de_vie_max     int
	Point_de_vie_restant int
	Skill                []string
	Argent               int
	Inventaire           []string
}

type Marchand struct {
	Nom        string
	Inventaire []string
	Argent     int
}

func main() {
	var char2 Marchand
	var char1 personnage
	char1.Init("Lunasphys", "Archer", 68, 1350, 1150, []string{"Coup de poing"}, 100, []string{"Potion de vie", "Potion de vie", "Potion de vie", "Potion de poison"})
	char2.Init("Marchand", []string{"Potion de vie", "épée de fer", "Potion de poison", "Livre de sort : Boule de feu"}, 1000000)
	//char1.TakePot()
	char1.menu(&Marchand{})
}
