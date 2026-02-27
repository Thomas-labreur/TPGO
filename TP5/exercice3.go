package main
import "fmt"

func est_dans(elt int, liste []int) bool {

    // Déclarer et affecter la sortie
    dans_liste := false

    // Cas où la liste est vide
    if len(liste)==0 {
        return dans_liste
    }

    // On parcours la liste tant que dans_liste est false
    for i:=0; i<len(liste) && !dans_liste; i++ {
        // Si on trouve elt dans la liste, dans_liste devient true
        dans_liste = liste[i]==elt
    }

    // On renvoie dans_liste
    return dans_liste
}

func intersection(a, b []int) []int {
    // On créée une liste vide
    intersection := []int{}
    
    // On ajoute les éléments de a dans cette liste s'ils sont dans b aussi
    for _, elem_a := range(a) {
        if est_dans(elem_a, b) {
            intersection = append(intersection, elem_a)
        }
    }
    return intersection
}

func union(a, b []int) []int {
    
    // On créée une copie de l'ensemble a
    union := make([]int, len(a))
    copy(union, a)
    
    // On ajoute les éléments de b dans cette copie s'ils ne sont pas deja dans a
    for _, elem_b := range(b) {
        if !est_dans(elem_b, a) {
            union = append(union, elem_b)
        }
    }
    return union
}

func difference(a, b []int) []int {

    // On créé une lite vide
    difference := []int{}

    // On ajoute les éléments de a dans cette liste s'ils ne sont pas dans b
    for _, elem_a := range(a) {
        if !est_dans(elem_a, b){
            difference = append(difference, elem_a)
        }
    }
    return difference
}

func main() {
    liste1 := []int{-3, 4, 2, 9, 8}
    liste2 := []int{6, 4, -4, 8}
    
    fmt.Printf("Liste 1: %v\n", liste1)
    fmt.Printf("Liste 2: %v\n\n", liste2)
    
    fmt.Printf("Union: %v\n", union(liste1, liste2))
    fmt.Printf("Intersection: %v\n", intersection(liste1, liste2))
    
    fmt.Printf("Difference entre Liste 1 et 2: %v\n", difference(liste1, liste2))
    fmt.Printf("Difference entre Liste 2 et 1: %v\n", difference(liste2, liste1))
}