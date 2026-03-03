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

	// On parcours le mots du début au milieu
	n := len(mot)
	for i := 0; i < n/2; i++ {

		// On compare alors le caracères i et n-i-1
		ci := toMaj(rune(mot[i]))
		cj := toMaj(rune(mot[n-1-i]))

		// On renvoie false s'ils sont différents (cela arrete la boucle)
		if ci != cj {
			return false
		}
	}

	// Si on a jamais renvoyé false, c'est bien un palindrome, on renvoie true
	return true
}

func main() {
    mot := input("Entrez un mot (sans espaces): ")
    if palinMot(mot) {
        fmt.Printf("Le mot %v est un palindrome.", mot)
    } else {
        fmt.Printf("Le mot %v n'est pas un palindrome.", mot)
    }
}