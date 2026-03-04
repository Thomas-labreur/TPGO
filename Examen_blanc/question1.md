## 1 Algorithme

Données: un code numérique (string)

Sortie: booléen indiquant si le code est correct ou non (bool)

Remarque: Il est indiqué que la position 0 est celle du caractère le plus **à droite**, on va donc lire "à l'envers" le code, de droite à gauche

Etapes:

    - DECLARER codeR DE TYPE []rune, n, somme DE TYPE int

    - AFFECTER []rune(code) à codeR, 0 A somme, len(codeR) A n

    - POUR i allant de 0 à len(codeR)-1

        * SI i%2==0 ALORS AFFECTER somme + int(codeR[n-1-i] - '0') A somme

        * SINON AFFECTER somme + (int(codeR[n-1-i] - '0')*2)%9 A somme
    
    - RENVOYER somme%10==0

