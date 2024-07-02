package pkg

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func Environment() error {
	// Lecture du fichier .env
	// Reading the .env file
	content, err := os.ReadFile(".env")
	if err != nil {
		return errors.New("cannot read.env file")
	}
	// Diviser le contenu en lignes
	// Splitting the file into lines
	lines := bufio.NewScanner(strings.NewReader(string(content)))
	for lines.Scan() {
		line := lines.Text()

		// Interprétation de chaque ligne
		// Interpreting each line
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			// Définir la variable d'environnement
			// Set the environment variable
			err := os.Setenv(key, value)
			if err != nil {
				return errors.New("Cannot set environment variable: " + key + " = " + value)
			}
		} else {
			return errors.New("Invalid line in .env file: " + line)
		}
	}
	if err := lines.Err(); err != nil {
		return errors.New("cannot read .env file")
	}
	return nil
}
