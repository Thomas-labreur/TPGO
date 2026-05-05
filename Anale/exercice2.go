package main 
import ( 
    "fmt" 
) 

// Renvoie la valeur a+10
// Affiche la valeur de a et a+10
func fonction1(a int) int { 
    var b = a + 10 
    fmt.Printf("a : %v \n", a) 
    fmt.Printf("b : %v \n", b) 
    return b 
} 

// Affiche les valeurs 0, 1 et a+1
func fonction2(a int) { 
    for a := 0; a < 2; a++ { 
        fmt.Printf("a : %v \n", a) 
    } 
    a = a + 1 
    fmt.Printf("a : %v \n", a) 
} 
func main() { 
    a := 6 
    b := 2 
    c := fonction1(b) 
    // affiche "a: 2" , "b: 12"
    fonction2(a) 
    // affiche "a: 0" "a: 1" et "a: 7"

    fmt.Printf("a : %v \n", a) // affiche "a: 6"
    fmt.Printf("b : %v \n", b) // affiche "b: 2"
    fmt.Printf("c : %v \n", c) // affiche "c: 12"
}