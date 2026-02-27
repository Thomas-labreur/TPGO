package main
import "fmt"

func decomposer(liste []int) ([]int, []int) {
    
    // Déclarer/initialiser les deux listes
    pair, impair := []int{}, []int{}
    
    // Parcourir la liste initiale
    for _, elem := range liste {
        if elem%2==0 {
            pair = append(pair, elem)
        } else {
            impair = append(impair, elem)
        }
    }
    
    return pair, impair
}

func main() {
    liste1 := []int{-3, 4, 2, 9, 8}
    liste2 := []int{6, -4, 8}
    p1, i1 := decomposer(liste1)
    p2, i2 := decomposer(liste2)
    fmt.Printf("Liste 1: %v\n  nombres pairs: %v\n  nombres impairs: %v\n", liste1, p1, i1)
    fmt.Printf("Liste 2: %v\n  nombres pairs: %v\n  nombres impairs: %v\n", liste2, p2, i2)
}