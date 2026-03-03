package main
import (
	"fmt"
)

func decalage1(chaine string) string {

	// Décalaration et Initialisation des liste de runes
	chaineR := []rune(chaine)
	resultat := make([]rune, len(chaineR))

	// Parcours des runes
	for i:=0; i<len(chaineR); i++ {

		// Si c'est une lettre, on met celle d'après
		if chaineR[i]>='a' && chaineR[i]<'z' {
			resultat[i] = chaineR[i]+1

		// Sauf pour 'z', on met le caractère 'a'
		} else if chaineR[i]=='z' {
			resultat[i] = 'a'
		
		// Sinon on ne change pas le caractère
		} else {
			resultat[i] = chaineR[i]
		}
	}

	// On oublie pas de convertir
	return string(resultat)	
}

func main() {
	chaine := "bonjour monsieur !"
	fmt.Printf("Phrase d'origine: %v\nPhrase transformée: %v", chaine, decalage1(chaine))
}