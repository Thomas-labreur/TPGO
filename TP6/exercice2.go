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

func main() {
    c := 'a'
	fmt.Printf("La version majuscule de %c est %c.", c, toMaj(c))