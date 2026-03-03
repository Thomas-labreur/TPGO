package main
import (
	"fmt"
)

func remplacer(chaine string, c1, c2 rune) string {

	// Décalaration et Initialisation des liste de runes
	chaineR := []rune(chaine)
	resultat := make([]rune, len(chaineR))

	// Parcours des runes
	for i:=0; i<len(chaineR); i++ {

		// S'il y a c1, on met c2 dans resultat
		if chaineR[i]== c1 {
			resultat[i] = c2

		// Sinon on met le caractère d'origine
		} else {
			resultat[i] = chaineR[i]
		}
	}

	// On oublie pas de convertir
	return string(resultat)	
}

func main() {
	chaine := "Cet élève va avoir 08 en programmation."
	fmt.Printf("Phrase d'origine: %v\nPhrase transformée: %v", chaine, remplacer(chaine, '1', '0'))
}