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

// Remplace les accents et les minuscules par une majuscule sans accent
func normaliser(r rune) rune {

	// Gestion des accents
	switch r {
	case 'à','á','â','ä','ã','À','Á','Â','Ä','Ã':
		return 'A'
	case 'ç','Ç':
		return 'C'
	case 'è','é','ê','ë','È','É','Ê','Ë':
		return 'E'
	case 'ì','í','î','ï','Ì','Í','Î','Ï':
		return 'I'
	case 'ò','ó','ô','ö','õ','Ò','Ó','Ô','Ö','Õ':
		return 'O'
	case 'ù','ú','û','ü','Ù','Ú','Û','Ü':
		return 'U'
	}

	// Gestion des minuscules
	if r >= 'a' && r <= 'z' {
		return r - ('a'-'A')
	}

	return r
}


func palinPhrase2(phrase string) bool {

	// On doit runifier la phrase sinon phrase[i] renvoie un seul octet
	phraseR := []rune(phrase)

	// Initialiser deux indices au début et à la fin de la chaine
	i, j := 0, len(phraseR)-1

	// Tant que i et j ne se sont pas rejoint
	for i<j {

		// On regarde les caractères i et j
		ci := normaliser(phraseR[i])
		cj := normaliser(phraseR[j])

		// Si ce sont des espaces, on passe au suivant
		if ci == ' ' {
			i++
		} else if cj == ' ' {
			j--
		} else {

			// S'ils sont différents, on arrête et on renvoie false
			if ci != cj {
				return false
			}

			// Sinon on passe aux suivants
			i++
			j--
		}
	}
	
	// Si on a jamais trouvé 2 caractères différents, on renvoie true
	return true
}

func main() {
    phrase := input("Entrez une phrase: ")
    if palinPhrase2(phrase) {
        fmt.Printf("La phrase '%v' est un palindrome.", phrase)
    } else {
        fmt.Printf("La phrase '%v' n'est pas un palindrome.", phrase)
    }
}