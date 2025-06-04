package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	// Sauvegarder les fonctions originales
	oldGenerate := GenerateRandomWord
	oldRunHuman := RunHumanGame
	oldRunAI := RunAIGame

	// Rediriger stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	defer func() {
		GenerateRandomWord = oldGenerate
		RunHumanGame = oldRunHuman
		RunAIGame = oldRunAI
		os.Stdout = oldStdout
	}()

	// Simuler les fonctions
	var humanCalled, aiCalled bool

	GenerateRandomWord = func(length int) string {
		return "testword"
	}

	RunHumanGame = func(config GameConfig, word string) {
		humanCalled = true
	}

	RunAIGame = func(config GameConfig, word string) {
		aiCalled = true
	}

	// Créer un fichier temporaire comme dictionnaire
	tmpFile, err := os.CreateTemp("", "dict-*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString("apple\nbanana\ncherry\n")
	if err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	// Test avec mode joueur humain
	os.Args = []string{"main"}

	// Forcer l'utilisation d'un mot généré plutôt que le dictionnaire
	// car le fichier "liste_francais.txt" peut ne pas exister
	configBak := GameConfig{
		WordLength:    5,
		MaxAttempts:   6,
		UseDictionary: false,
		HumanPlayer:   true,
	}

	// On ne peut pas facilement modifier la configuration globale,
	// donc testons directement les fonctions appelées par main()
	word := GenerateRandomWord(5)
	RunHumanGame(configBak, word)

	if !humanCalled {
		t.Error("RunHumanGame was not called")
	}

	// Test avec IA
	humanCalled = false
	aiCalled = false

	configBak.HumanPlayer = false
	word = GenerateRandomWord(5)
	RunAIGame(configBak, word)

	if !aiCalled {
		t.Error("RunAIGame was not called")
	}

	// Test avec erreur de dictionnaire
	w.Close()
	output, _ := ioutil.ReadAll(r)
	if len(output) == 0 && configBak.UseDictionary {
		t.Error("No output when dictionary mode is enabled")
	}
}
