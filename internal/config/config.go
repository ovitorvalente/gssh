package config

import (
	"os"
	"path/filepath"
)

// SSHPaths retorna os caminhos padrão para chave SSH.
func SSHPaths() (keyPath, pubKeyPath string) {
	home, err := os.UserHomeDir()
	if err != nil {
		home = os.Getenv("HOME")
		if home == "" {
			home = "."
		}
	}
	keyPath = filepath.Join(home, ".ssh", "id_ed25519")
	pubKeyPath = keyPath + ".pub"
	return keyPath, pubKeyPath
}
