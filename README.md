# Projet Go - Exercices de Programmation

Ce projet contient trois exercices en langage Go qui démontrent différentes compétences de programmation.

## Structure du Projet

Le projet est organisé en trois dossiers principaux :

```txt
/sujet/
├── exo1/     # Fonctions de base
├── exo2/     # Gestion d'une liste d'étudiants
└── exo3/     # Jeu de devinette de mots (Word Guess)
```

## Exercice 1 - Basics (4 points)

Implémentation de fonctions de base en Go :

- Somme de nombres
- Vérification de nombres pairs
- Recherche de maximum
- Calcul de factorielle
- Comptage d'occurrences de caractères
- Filtrage de nombres pairs
- Inversion de chaînes de caractères

### Tester l'Exercice 1

```bash
cd exo1
go test -v -cover
```

## Exercice 2 - Gestion d'une liste d'étudiants (5 points)

Système de gestion d'étudiants avec :

- Structure Student (nom, âge, note)
- Structure StudentList (liste d'étudiants)
- Fonctions pour ajouter, supprimer, trier et afficher des étudiants
- Validations des données

### Tester l'Exercice 2

```bash
cd exo2
go test -v -cover
```

## Exercice 3 - Word Guess (11 points)

Un jeu de devinette de mots, similaire au concept de "Wordle" :

- Génération de mots secrets (aléatoire ou à partir d'un dictionnaire)
- Mode joueur humain (entrée à partir du terminal)
- Mode IA (l'ordinateur devine le mot)
- Évaluation des lettres (position correcte, position incorrecte, absente)
- Tests unitaires complets

### Tester l'Exercice 3

```bash
cd exo3
go test -v -cover
```

### Exécuter le jeu Word Guess

```bash
cd exo3
go build
./exo3        # Mode humain par défaut avec dictionnaire
./exo3 -ai    # Mode IA
./exo3 -nodict # Sans utiliser de dictionnaire
```

## Tests et Couverture de Code

Tous les exercices sont accompagnés de tests unitaires complets avec une couverture de code supérieure à 90%. Les tests peuvent être exécutés avec :

```bash
go test -v -cover
```

Pour générer un rapport de couverture plus détaillé :

```bash
go test -cover -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Technologies Utilisées

- Go 1.17+ (uniquement les bibliothèques standards)
- Tests unitaires natifs de Go

## Contraintes de Développement

- Utilisation exclusive des bibliothèques standards de Go
- Tous les tests unitaires doivent réussir
- Couverture de code proche de 100%
- Respect des structures d'interfaces définies dans les README.md des exercices
