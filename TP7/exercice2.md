**Données**: chaine (string), c1, c2 (runes)

**Sortie**: la même chaîne avec c2 à la place de c1 (string)

**Etapes**:
    DECLARER chaineR ET resultat DE TYPE []rune
    AFFECTER []rune(chaine) A chaineR
    POUR i DE 0 A len(chaineR)-1
        SI chaineR[i] == c1 ALORS AFFECTER c2 A resultat[i]
        SINON AFFECTER chaineR[i] A resultat[i]
        FIN SI
    FIN POUR
    CONVERTIR resultat EN TYPE string
    RENVOYER resultat