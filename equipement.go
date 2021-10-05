package main

func (char *equipement) Init(Tete string, Torse string, Bottes string) {
	char.Tete = Tete
	char.Torse = Torse
	char.Bottes = Bottes
}

func AddEquip(tab []string, s string) []string {
	if len(tab)-1 >= 2 {
		fmt.Println("Vous avez déjà équipé un item")
		return tab
	}
	return append(tab, s)
}


func (char *equipement) EquipHead() {
	for _, chapeau := range char1.Equipement {
		if chapeau == "Chapeau de Sorcier" {
}