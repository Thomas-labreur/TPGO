**Données**: chaine (string), n (int)

**Sortie**: la même chaîne avec décalage de n des caractères minuscule (string)

**Etapes**:
- DECLARER chaineR ET resultat DE TYPE []rune
- AFFECTER []rune(chaine) A chaineR
- POUR i DE 0 A len(chaineR)-1
    - SI chaineR[i] >= 'a' OU chaineR[i]<='z' ALORS AFFECTER 'a'+ (chaineR[i]+n)%26 A resultat[i]
    - SINON AFFECTER chaineR[i] A resultat[i]
    - FIN SI 
- FIN POUR
- CONVERTIR resultat EN TYPE string
- RENVOYER resultat