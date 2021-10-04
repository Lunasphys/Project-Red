package main

func CountItem(tab []string, s string) int {
	var res int = 0
	for _, a := range tab {
		if a == s {
			res++
		}
	}
	return res
}

func (char1 *personnage) Craft(quantityA int, quantityB int, itemA string, itemB string) bool {
	if CountItem(char1.Inventaire, itemA) >= quantityA && CountItem(char1.Inventaire, itemB) >= quantityB {
		return true
	}
	return false
}
