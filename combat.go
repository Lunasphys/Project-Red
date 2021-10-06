package main 

import "fmt"

func TrainingFight(char1 *personnage, char3 *monstre) { // Initialisation combat d'entrainement
    fmt.Println(char1.Nom, " engage le combat d'entrainement")
    fmt.Println()
    fmt.Println("A vos armes !!!!")
    for count := 1; ; count++ { // Condition de fin de combat
        if char1.Point_de_vie_restant > 0 || char3.Point_de_vie_restant > 0 {
            fmt.Println("======== Tour ", count, " ========") // Initialisation nombre de tours
            CharTurn(char1, char3)
            GoblinPattern(char1, char3, count)
        }
    }
}

func (char3 *monstre) GoblinPattern() {

}
