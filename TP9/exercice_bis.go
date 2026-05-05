package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    )
    
//////////////////////////
// DEFINITION DES TYPES //
//////////////////////////

type adresse struct{
    rue, codePostal, ville string
}
type etudiant struct{
    num int
    nom, prenom string
    adresseEtu adresse
    notes []float64
}

///////////////////////
// FONCTIONS ANNEXES //
///////////////////////

func input(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(message)
	scanner.Scan()
	return scanner.Text()
}

/////////////////////////
// FONCTIONS DEMANDEES //
/////////////////////////

func inputEtu() etudiant {
    num_str := input("Entrez un numero étudiant:")
    num, err := strconv.Atoi(num_str)
    for err!=nil {
        num_str = input("Numero invalide! Entrez un numero étudiant:")
        num, err = strconv.Atoi(num_str)
    }
    
    nom := input("Entrez le nom: ")
    prenom := input("Entrer le prenom:")
    
    rue := input("Entrez la rue:")
    codePostal := input("Entrez le code postal:")
    ville := input("Entrez la ville:")
    
    adresseEtu := adresse{rue, codePostal, ville} 
    e := etudiant{num, nom, prenom, adresseEtu, []float64{}}
    
    return e
}

func etuToString(e etudiant) string {
    var resultat = ""
    var num_str = strconv.Itoa(e.num)
    var n_str string
    
    resultat += "Numéro: " + num_str + "\n"
    resultat += "Nom: " + e.nom + "\n"
    resultat += "Prénom: " + e.prenom + "\n"
    
    resultat += "Adresse: " + e.adresseEtu.rue + ", " + e.adresseEtu.codePostal + " " + e.adresseEtu.ville
    
    resultat += "\nNotes: "
    
    for _, n := range(e.notes) {
        n_str = strconv.FormatFloat(n, 'f', -1, 64)
        resultat += n_str + " "
    }
    
    return resultat
}

func findEtu(num int, classe []etudiant) (int, etudiant) {
    for i, e := range(classe) {
        // Si je trouve l'étudiant, je le renvoie avec son index
        if e.num == num {
            return i, e
        }
    }
    
    // Si j'ai parcourru toute la classe sans le trouver, je renvoie un étudiant vide, avec l'index -1
    return -1, etudiant{}
}

func ajoutNote(num int, classe []etudiant) error {
    
    // Je cherche l'index de l'étudiant dans la classe
    i, e := findEtu(num, classe)
    
    // Si je l'ai trouvé, je demande de saisir une note et je l'ajoute à celles de l'étudiant
    if i != -1 {
        n_str := input("Entrer une note pour "+e.nom+" : ")
        n, err2 := strconv.ParseFloat(n_str, 64)
        for err2 != nil || n<0.0 || n>20.0 {
            n_str = input("Note invalide !Entrer une note: ")
            n, err2 = strconv.ParseFloat(n_str, 64)
        }
        
        classe[i].notes = append(classe[i].notes, n)
        return nil // on oublie pas de renvoyer quelque chose!
        
    // Sinon, je renvoie juste un message d'erreur
    } else {
        return fmt.Errorf("L'étudiant numéro %d n'existe pas", num)
    }
}

/////////////////////////
// FONCTION PRINCIPALE //
/////////////////////////

func main() {
    var err error
    classe := []etudiant{
        inputEtu(),
        inputEtu(),
    }
    
    for _, e := range(classe) {
        for j:=0; j<3; j++ {
            err = ajoutNote(e.num, classe)
            for err != nil {
                fmt.Println(err)
                err = ajoutNote(e.num, classe)  
            }
        }
    }
    
    fmt.Println("\n")
    fmt.Println(etuToString(classe[0]))
    fmt.Println(etuToString(classe[0]))
}