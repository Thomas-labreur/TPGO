package main
import "fmt"

func procedure_inverser(liste []int) {
    n := len(liste)
	for i := 0; i < n/2; i++ {
		liste[i], liste[n-1-i] = liste[n-1-i], liste[i]
	}
}

func fonction_inverser(liste []int) []int {
    n := len(liste)
    liste_inv := make([]int, n)
    for i, _ := range liste {
        liste_inv[i] = liste[n-1-i]
    }
    return liste_inv
}

func main() {
    liste1 := []int{-3, 4, 2, 9, 8}
    liste2 := []int{6, -4, 8}
    fmt.Printf("Liste1: %v\n", liste1)
    fmt.Printf("Liste2: %v\n\n", liste2)

    liste2_inv := fonction_inverser(liste2)
    fmt.Printf("Liste2 inversée: %v\n", liste2_inv)
    fmt.Printf("L'objet Liste2 n'a pas changé: %v\n\n", liste2)
    
    procedure_inverser(liste1)
    fmt.Printf("Mais l'objet Liste1 a changé: %v\n", liste1)
    fmt.Println("L'ancien Liste1 n'existe plus.\n")
}