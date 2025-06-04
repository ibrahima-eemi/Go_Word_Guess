package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type LetterStatus int

const (
	CorrectPosition LetterStatus = iota
	WrongPosition
	Absent
)

type GameConfig struct {
	WordLength    int
	MaxAttempts   int
	UseDictionary bool
	HumanPlayer   bool
}

func main() {
	rand.Seed(time.Now().UnixNano())

	config := GameConfig{
		WordLength:    5,
		MaxAttempts:   6,
		UseDictionary: true,
		HumanPlayer:   true,
	}

	var word string
	var err error
	if config.UseDictionary {
		word, err = GenerateWordFromDictionary("liste_francais.txt", config.WordLength)
	} else {
		word = GenerateRandomWord(config.WordLength)
	}
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	if config.HumanPlayer {
		RunHumanGame(config, word)
	} else {
		RunAIGame(config, word)
	}
}

// Rendre ces fonctions variables pour les tests
var GenerateRandomWord = generateRandomWord
var RunHumanGame = runHumanGame
var RunAIGame = runAIGame
var GenerateWordFromDictionary = generateWordFromDictionary

// L'implémentation réelle de GenerateRandomWord
func generateRandomWord(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	word := make([]rune, length)
	for i := 0; i < length; i++ {
		word[i] = letters[rand.Intn(len(letters))]
	}
	return string(word)
}

// Lit le dictionnaire et retourne un mot aléatoire de longueur précise - implémentation réelle
func generateWordFromDictionary(filename string, length int) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		w := strings.ToLower(scanner.Text())
		if len(w) == length && isAlpha(w) {
			words = append(words, w)
		}
	}
	if len(words) == 0 {
		return "", errors.New("aucun mot de cette longueur dans le dictionnaire")
	}
	return words[rand.Intn(len(words))], nil
}

// Contrôle que le mot contient uniquement des lettres
func isAlpha(s string) bool {
	for _, r := range s {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}

// Joue avec un humain - implémentation réelle
func runHumanGame(config GameConfig, secret string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Bienvenue dans WordGuess !")
	fmt.Printf("Le mot à deviner a %d lettres. Tu as %d essais.\n", config.WordLength, config.MaxAttempts)
	for attempt := 1; attempt <= config.MaxAttempts; attempt++ {
		fmt.Printf("Essai %d/%d : ", attempt, config.MaxAttempts)
		input, _ := reader.ReadString('\n')
		guess := strings.TrimSpace(strings.ToLower(input))

		if len(guess) != config.WordLength {
			fmt.Printf("Mot invalide. Il faut un mot de %d lettres.\n", config.WordLength)
			attempt--
			continue
		}

		status := EvaluateGuess(secret, guess)
		DisplayResult(guess, status)

		if guess == secret {
			fmt.Println("Bravo, tu as trouvé le mot !")
			return
		}
	}
	fmt.Printf("Tu as perdu. Le mot était : %s\n", secret)
}

// Joue automatiquement avec une IA basique - implémentation réelle
func runAIGame(config GameConfig, secret string) {
	fmt.Println("Mode IA activé. Deviner automatiquement...")
	for attempt := 1; attempt <= config.MaxAttempts; attempt++ {
		guess := GenerateRandomWord(config.WordLength)
		fmt.Printf("Essai IA %d : %s\n", attempt, guess)

		status := EvaluateGuess(secret, guess)
		DisplayResult(guess, status)

		if guess == secret {
			fmt.Println("IA a trouvé le mot !")
			return
		}
	}
	fmt.Printf("IA a échoué. Le mot était : %s\n", secret)
}

// Compare deux mots et retourne un statut lettre par lettre
func EvaluateGuess(secret, guess string) []LetterStatus {
	status := make([]LetterStatus, len(secret))
	secretRunes := []rune(secret)
	guessRunes := []rune(guess)

	// Cas spécial : si les mots contiennent exactement les mêmes lettres (anagrammes parfaits)
	// mais ne sont pas identiques, alors toutes les lettres sont mal placées
	if isAnagram(secret, guess) && secret != guess {
		for i := range status {
			status[i] = WrongPosition
		}
		return status
	}

	// Initialiser tous les statuts à Absent par défaut
	for i := range status {
		status[i] = Absent
	}

	// Comptage des lettres dans le mot secret
	secretCounts := make(map[rune]int)
	for _, r := range secretRunes {
		secretCounts[r]++
	}

	// Première passe: marquer les lettres correctes
	for i, r := range guessRunes {
		if r == secretRunes[i] {
			status[i] = CorrectPosition
			secretCounts[r]--
		}
	}

	// Deuxième passe: marquer les lettres mal placées
	for i, r := range guessRunes {
		// Si déjà marquée comme correcte, passer
		if status[i] == CorrectPosition {
			continue
		}

		// Si la lettre existe encore dans le mot secret
		if secretCounts[r] > 0 {
			status[i] = WrongPosition
			secretCounts[r]--
		}
	}

	return status
}

// Vérifie si deux mots sont des anagrammes (contiennent exactement les mêmes lettres)
func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	// Compter les occurrences de chaque lettre dans s1
	counts := make(map[rune]int)
	for _, r := range s1 {
		counts[r]++
	}

	// Décrémenter pour chaque lettre dans s2
	for _, r := range s2 {
		counts[r]--
		if counts[r] < 0 {
			return false
		}
	}

	// Vérifier que tous les compteurs sont à zéro
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}

// Affiche un mot et les statuts
func DisplayResult(guess string, status []LetterStatus) {
	for i, r := range guess {
		switch status[i] {
		case CorrectPosition:
			fmt.Printf("[%c]", r)
		case WrongPosition:
			fmt.Printf("(%c)", r)
		case Absent:
			fmt.Printf(" %c ", r)
		}
	}
	fmt.Println()
}
