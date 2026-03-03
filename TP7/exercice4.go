package main
import (
	"fmt"
)

func decalage(chaine string, n int) string {

	// Décalaration et Initialisation des liste de runes
	chaineR := []rune(chaine)
	resultat := make([]rune, len(chaineR))

	// Parcours des runes
	for i:=0; i<len(chaineR); i++ {

		// Si c'est une lettre, on la décale de n et on prend le modulo 26
		if chaineR[i]>='a' && chaineR[i]<'z' {
			resultat[i] = 'a' + (chaineR[i] - 'a' + rune(n)) % 26
			
		// Pourquoi ce calcul?
		// En ASCII, 'a' à 'z' ne sont pas les caractères 0 à 25, mais 97 à 122 .
		// en faisant chaineR[i] - 'a', j'obtiens le rang dans l'alphabet de chaineR[i] (de 0 à 25).
		// Là je rajoute n pour obtenir le rang dans l'alphabet du caractère décalé.
		// Je prend le modulo 26 pour revenir a 'a' après 'z'.
		// Enfin, je rajoute à nouveau 'a' pour revenir entre 97 et 122.
		
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
	fmt.Printf("Phrase d'origine: %v\nPhrase transformée: %v", chaine, decalage(chaine, 13))
}