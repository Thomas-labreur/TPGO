package main
import (
    "fmt"
    "bufio"
    "os"
    )

func input(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(message)
	scanner.Scan()
	return scanner.Text()
}

// Fonction de vérification de la saisie (non demandé)
func saisie_incorrecte(chaine string) bool {
	chaineR := []rune(chaine)
	caractere_invalide := false
	for _, c := range chaineR {
		if (c<'a' || c>'z') && (c<'A' || c>'z') && c!='?' && c!=':' && c!=',' && c!=';' && c!='!' && c!='.' {
			caractere_invalide = true
		}
	}
	return len(chaineR) > 20 || caractere_invalide
}

// Fonction de saisie
func saisie() string {
	chaine := input("Entrez une chaine de moins de 20 caractères:")
	for saisie_incorrecte(chaine) {
		chaine = input("Saisie invalide ! Entrez une chaine de moins de 20 caractères:")
	}
	return chaine
}

// Fonction pour créer les 3 listes
func creer_liste(chaine string) ([]rune, []rune, []rune) {
	chaineR := []rune(chaine)
	voyelles, consonnes, ponctuation := []rune{}, []rune{}, []rune{}
	for _, r := range chaineR {
		switch r {
		case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U', 'y', 'Y':
			voyelles = append(voyelles, r)
		case '.', '?', '!', ',', ';', ':':
			ponctuation = append(ponctuation, r)
		default:
			consonnes = append(consonnes, r)
		}
	}
	return voyelles, consonnes, ponctuation
	// RQUE: Si on ne veut pas de doublons dans les listes, on doit créer une fonction qui
	// vérifie si une rune est déja présente dans une liste de rune donnée. 
	// On l'applique a chaque rune croisée avec la liste correspondante et on ajoute la rune
	// Que dans le cas ou ça renvoie false
}

// Fonction principale
func main() {
	chaine := saisie()
	voyelles, consonnes, ponctuation := creer_liste(chaine)
	fmt.Printf("Voyelles: %c\nConsonnes: %c\nPonctuation: %c\n", voyelles, consonnes, ponctuation)
}

