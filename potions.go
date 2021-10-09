package main

import (
	"fmt"
	"time"
)

func (char1 *personnage) TakePot() { // Permet de prendre une potion de vie

	if !Contains(char1.Inventaire, potionDeVie) { // Si inventaire ne contient pas de potion de vie, l'indique au joueur
		fmt.Println("Vous n'avez pas de potion")
	}
	if char1.Point_de_vie_max == char1.Point_de_vie_restant { // Si les PVs max sont égaux aux PVs restants alors l'action est bloqué
		fmt.Println("Vous ne pouvez pas utiliser de potion")

	}
	if Contains(char1.Inventaire, potionDeVie) { // Si potion de vie est dans l'inventaire
		if char1.Point_de_vie_restant < char1.Point_de_vie_max { // Que les PVs restants sont inférieurs aux PVs max alors le joueur peut utiliser sa potion
			char1.Point_de_vie_restant += 25
			fmt.Println("Vous avez utilisé une potion de vie")
			char1.Inventaire = RemoveInventory(char1.Inventaire, potionDeVie) // Retire la potion de vie qui vient d'être utiliser
		}
	}
	if char1.Point_de_vie_restant > char1.Point_de_vie_max { // Les PVs restants ne peuvent pas dépasser les PVs max
		char1.Point_de_vie_restant = char1.Point_de_vie_max

	}
	fmt.Println("PV :", char1.Point_de_vie_restant, "/", char1.Point_de_vie_max)
	fmt.Println(char1.Inventaire)
}

func (char1 *personnage) PoisonPot() { // Permet d'utiliser une potion de poison sur soi-même (le gout du danger)
	if !Contains(char1.Inventaire, potionDePoison) { // Si potion de poison n'est pas contenu dans l'inventaire alors action impossible
		fmt.Println("Vous n'avez pas de potion de poison")
	}
	if Contains(char1.Inventaire, potionDePoison) { // Si la potion est bien présente dans l'inventaire, l'action peut être lancée
		fmt.Println("Applique un poison sur vous même pendant 3 secondes")
		char1.Point_de_vie_restant -= 10
		time.Sleep(1 * time.Second)
		fmt.Println("PV :", char1.Point_de_vie_restant, "/", char1.Point_de_vie_max)
		fmt.Println("Il reste 2 secondes")
		char1.Point_de_vie_restant -= 10
		time.Sleep(1 * time.Second)
		fmt.Println("PV :", char1.Point_de_vie_restant, "/", char1.Point_de_vie_max)
		fmt.Println("Il reste 1 seconde")
		char1.Point_de_vie_restant -= 10
		time.Sleep(1 * time.Second)
		fmt.Println("L'effet de la potion de poison s'estompe")
		char1.Inventaire = RemoveInventory(char1.Inventaire, potionDePoison) // Retire la potion de poison qui vient d'être utilisé
	}
	fmt.Println("PV :", char1.Point_de_vie_restant, "/", char1.Point_de_vie_max)

}

func (char1 *personnage) ThrowPoisonPot(char3 *monstre) { // Permet d'envoyer une potion de poison sur l'ennemi
	if !Contains(char1.Inventaire, potionDePoison) { // Si le joueur n'en a pas dans son inventaire, action impossible, renvoi d'un message d'erreur
		fmt.Println("Vous n'avez pas de potion de poison")
	}
	if Contains(char1.Inventaire, potionDePoison) { // Si l'objet est présent dans l'inventaire alors les effets sont appliqués sur l'ennemi
		fmt.Println("Applique un poison sur l'ennemi pendant 3 secondes")
		char3.Point_de_vie_restant -= 10
		time.Sleep(1 * time.Second)
		fmt.Println("PV :", char3.Point_de_vie_restant, "/", char3.Point_de_vie_max)
		fmt.Println("Il reste 2 secondes")
		char3.Point_de_vie_restant -= 10
		fmt.Println("PV :", char3.Point_de_vie_restant, "/", char3.Point_de_vie_max)
		time.Sleep(1 * time.Second)
		fmt.Println("Il reste 1 seconde")
		char3.Point_de_vie_restant -= 10
		time.Sleep(1 * time.Second)
		fmt.Println("L'effet de la potion de poison s'estompe")
		char1.Inventaire = RemoveInventory(char1.Inventaire, potionDePoison) // Enlève l'objet utilisé de l'inventaire du joueur
	}
	fmt.Println("PV :", char3.Point_de_vie_restant, "/", char3.Point_de_vie_max)
}
func (char1 *personnage) ThrowLifePot(char3 *monstre) { // Permet d'envoyer une potion de soin sur l'ennemi (Pourquoi pas)

	// Does char1.Inventaire contains a potion ?
	// If no -> I dont have potion
	// If pv == pv max -> Potion useless
	// Prend la potion
	// J'enleve la potion
	// si pv > pv max alors pv = pv max

	if !Contains(char1.Inventaire, potionDeVie) { // Si le joueur n'en a pas dans son inventaire alors l'action ne peut pas être réalisée
		fmt.Println("Vous n'avez pas de potion à envoyer")
	}
	if Contains(char1.Inventaire, potionDeVie) { // Si le joueur possède une potion de vie dans son inventaire alors il peut l'envoyer contre l'ennemi
		if char3.Point_de_vie_restant <= char3.Point_de_vie_max {
			char3.Point_de_vie_restant += 25
			fmt.Println("Vous venez de soigner le gobelin, vous avez un QI négatif")
			char1.Inventaire = RemoveInventory(char1.Inventaire, potionDeVie)
		}
	}
	if char3.Point_de_vie_restant > char3.Point_de_vie_max { // Les PVs restants du gobelin ne peuvent pas excéder les PVs maxs
		char3.Point_de_vie_restant = char3.Point_de_vie_max

	}
	fmt.Println("PV :", char3.Point_de_vie_restant, "/", char3.Point_de_vie_max)
	fmt.Println(char1.Inventaire)
}

func (char1 *personnage) TakeManaPot() { //Fonction qui permet d'utiliser une potion de vie

	if !Contains(char1.Inventaire, potionDeMana) {
		fmt.Println("Vous n'avez pas de potion de Mana")
	}
	if char1.Point_de_mana_max == char1.Point_de_mana_restant {
		fmt.Println("Vous ne pouvez pas utiliser de potion de Mana")
	}
	if Contains(char1.Inventaire, potionDeMana) {
		if char1.Point_de_mana_restant < char1.Point_de_mana_max {
			char1.Point_de_mana_restant += 20
			fmt.Println("Vous avez utilisé une potion de Mana")
			char1.Inventaire = RemoveInventory(char1.Inventaire, potionDeMana)
		}
	}
	if char1.Point_de_mana_restant > char1.Point_de_mana_max {
		char1.Point_de_mana_restant = char1.Point_de_mana_max

	}
	fmt.Println("Vous avez maintenant :", char1.Point_de_mana_restant, "/", char1.Point_de_mana_max, "Points de Mana")
	fmt.Println(char1.Inventaire)
}
