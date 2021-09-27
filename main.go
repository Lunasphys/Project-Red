package main

type personnage struct {
	Nom                  string
	Classe               string
	Niveau               int
	Point_de_vie_max     int
	Point_de_vie_restant int
	Inventaire           []string
	Money                int
}

func main() {
	var char2 personnage
	var char1 personnage
	char1.Init("Lunasphys", "Archer", 68, 1350, 1150, []string{"Potion", "Potion", "Potion", "Potion de poison"}, 100)
	char2.Init("Marchand", "Marchand", 100000, 10000, 10000, []string{"Potion", "épée de fer", "Potion de Poison"}, 1000000)
	//char1.TakePot()
	char1.menu()
}
