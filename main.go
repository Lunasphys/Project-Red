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
	var char1 personnage
	char1.Init("Lunasphys", "Archer", 68, 1350, 1000, []string{"Potion", "Potion", "Potion"}, 100)
	//char1.TakePot()
	char1.menu()
}
