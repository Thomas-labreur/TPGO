package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func input(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(message)
	scanner.Scan()
	return scanner.Text()
}

func toMaj(c rune) rune {
	// Les runes sont en fait des entiers, on peut leur faire
	// faire des opérations comme 'a'+1 -> 'b'.
	// La différence entre une minuscule et une majuscule est toujours la même
	// On peut donc enlever cette différence au caractères si celui-ci est une
	// minuscule.
    if c>='a' && c <= 'z' {
        return c - ('a'-'A')
    } else {
        return c
    }
}

func palinMot(mot string) bool {
	n := len(mot)
	for i := 0; i < n/2; i++ {
		ci := toMaj(rune(mot[i]))
		cj := toMaj(rune(mot[n-1-i]))
		if ci != cj {
			return false
		}
	}
	return true
}

func palinPhrase1(phrase string) bool {
	// Retrait des espaces
	sans_espace := strings.ReplaceAll(phrase, " ", "")

	// appel a palinMot
	return palinMot(sans_espace)
}

func palinPhrase2(phrase string) bool {

	// Initialiser deux indices au début et à la fin de la chaine
	i, j := 0, len(phrase)-1

	// Tant que i et j ne se sont pas rejoint
	for i<j {

		// On regarde les caractères i et j
		ci := toMaj(rune(phrase[i]))
		cj := toMaj(rune(phrase[j]))

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