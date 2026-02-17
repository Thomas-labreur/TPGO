package main
import ("fmt"; "strconv")

func input(message string) string {

	// Déclaration
	var resultat string

	// Lecture de la console
	fmt.Println(message)
	fmt.Scan(&resultat)

	// Renvoi du résultat
	return resultat
}

func main() {
    
    // demander age
    age := input("Bonjour, quel âge as-tu?") 
    //age, _ := strconv.Atoi(age_str)
    
    // Vérifier 
    if age>64 { // si age est un string, erreur !
        fmt.Println("Tu peux partir à la retraite ! Whouhou !")
    } else if age<18 {
        fmt.Println("Bon courage pour le BAC.")
    } else {
        fmt.Println("Retournes bosser et plus vite que ça !")
    }
}