package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const (
	DesktopPath = "C:\\Users\\windows\\Desktop"
	Key         = 1
	Code        = "1234"
)

func EncryptFile(filename string, key byte) error {
	// Lire le contenu du fichier
	fichier, err := OpenFile(filename)
	if err != nil {
		return err
	}

	fichierChiffre := []byte{}
	for _, bit := range fichier {
		fichierChiffre = append(fichierChiffre, bit+key)
	}

	// Écrire le contenu chiffré dans un nouveau fichier
	if err := os.WriteFile(filename, fichierChiffre, 0644); err != nil {
		return fmt.Errorf("os.WriteFile(filename, fichierChiffre, 0644): %v", err)
	}
	return nil
}

func OpenFile(filename string) ([]byte, error) {
	result, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile(filename): %v", err)
	}
	return result, nil
}

func Decrypt(filename string, key byte) error {
	return EncryptFile(filename, -key)
}

func ClearConsole() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
