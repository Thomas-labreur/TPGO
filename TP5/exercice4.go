package main
import "fmt"

func fusion(a, b []int) []int {

	// Déclarer et initialiser les variables
	fusion := []int{}
	i, j := 0, 0

	// Parcours des deux listes
	for i<len(a) && j<len(b) {

		// On ajoute le plus petit a la liste et on incrémente l'indice de sa liste d'origine
		if a[i] <= b[j] {
			fusion = append(fusion, a[i])
			i++
		} else {
			fusion = append(fusion, b[j])
			j++
		}
	}

	// On ajoute les éléments restants
	fusion = append(fusion, a[i:]...)
	fusion = append(fusion, b[j:]...)
	return fusion
}

func main() {
	liste1 := []int{10, 20, 30}
	liste2 := []int{1, 2, 3, 4, 25}

	fmt.Printf("Liste 1: %v\n", liste1)
    fmt.Printf("Liste 2: %v\n\n", liste2)
    
    fmt.Printf("Fusion: %v\n", fusion(liste1, liste2))
}