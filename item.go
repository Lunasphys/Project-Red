package main

const (
	chapeauAventurier = "Chapeau de l'aventurier"
	tuniqueAventurier = "Tunique de l'aventurier"
	bottesAventurier  = "Bottes de l'aventurier"
	potionDeVie       = "Potion de vie"
	potionDePoison    = "Potion de poison"
	bouleDeFeu        = "Boule de feu"
)

func (char1 *personnage) AddItem(item string) {
	if item == chapeauAventurier {
		char1.Point_de_vie_max += 10
		char1.EquipHead(chapeauAventurier)
		char1.Inventaire = RemoveInventory(char1.Inventaire, chapeauAventurier)
	} else if item == tuniqueAventurier {
		char1.Point_de_vie_max += 25
		char1.EquipChest(tuniqueAventurier)
		char1.Inventaire = RemoveInventory(char1.Inventaire, tuniqueAventurier)
	} else if item == bottesAventurier {
		char1.Point_de_vie_max += 15
		char1.EquipBoots(bottesAventurier)
		char1.Inventaire = RemoveInventory(char1.Inventaire, bottesAventurier)
	}
}

func (char1 *personnage) RemoveItem(item string) {
	if item == chapeauAventurier {
		char1.Point_de_vie_max -= 10
		char1.EquipHead("")
		char1.Inventaire = AddInventory(char1.Inventaire, chapeauAventurier)
	} else if item == tuniqueAventurier {
		char1.Point_de_vie_max -= 25
		char1.EquipChest("")
		char1.Inventaire = AddInventory(char1.Inventaire, tuniqueAventurier)
	} else if item == bottesAventurier {
		char1.Point_de_vie_max -= 15
		char1.EquipBoots("")
		char1.Inventaire = AddInventory(char1.Inventaire, bottesAventurier)
	}
}
