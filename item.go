package main

const (
	chapeauAventurier = "Chapeau de l'aventurier"
	tuniqueAventurier = "Tunique de l'aventurier"
	bottesAventurier = "Bottes de l'aventurier"
)

func (char1 *personnage) AddItem(item string) {
	if char1.Equipement.Tete == chapeauAventurier {
		char1.Point_de_vie_max += 10
		char1.EquipHead(chapeauAventurier)
		RemoveInventory(char1.Inventaire, chapeauAventurier)
	}
	if char1.Equipement.Torse == tuniqueAventurier {
		char1.Point_de_vie_max += 25
		char1.EquipChest(tuniqueAventurier)
		RemoveInventory(char1.Inventaire, tuniqueAventurier)
	}
	if char1.Equipement.Bottes == bottesAventurier {
		char1.Point_de_vie_max += 15
		char1.EquipBoots(bottesAventurier)
		RemoveInventory(char1.Inventaire, bottesAventurier)
	}
}

func (char1 *personnage) RemoveItem(item string) {
	if char1.Equipement.Tete == chapeauAventurier {
		char1.Point_de_vie_max -= 10
		char1.EquipHead("")
		AddInventory(char1.Inventaire, chapeauAventurier)
	}
	if char1.Equipement.Torse == tuniqueAventurier {
		char1.Point_de_vie_max -= 25
		char1.EquipChest("")
		AddInventory(char1.Inventaire, tuniqueAventurier)
	}
	if char1.Equipement.Bottes == bottesAventurier {
		char1.Point_de_vie_max -= 15
		char1.EquipBoots("")
		AddInventory(char1.Inventaire, bottesAventurier)
	}
}

