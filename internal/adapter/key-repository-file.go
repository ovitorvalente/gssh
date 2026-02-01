package adapter

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/ovitorvalente/gssh/internal/config"
)

type KeyRepositoryFile struct {
	pathKey       string
	pathPublicKey string
}

func NewKeyRepositoryFile() *KeyRepositoryFile {
	pathKey, pathPublicKey := config.SSHPaths()
	return &KeyRepositoryFile{
		pathKey:       pathKey,
		pathPublicKey: pathPublicKey,
	}
}

func (r *KeyRepositoryFile) Exists() bool {
	_, err := os.Stat(r.pathKey)
	return err == nil
}

func (r *KeyRepositoryFile) Generate() error {
	cmd := exec.Command(
		"ssh-keygen",
		"-t", "ed25519",
		"-f", r.pathKey,
		"-N", "",
		"-q",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (r *KeyRepositoryFile) ReadPublicKey() (string, error) {
	data, err := os.ReadFile(r.pathPublicKey)
	if err != nil {
		return "", fmt.Errorf("ler chave pública: %w", err)
	}
	return strings.TrimSpace(string(data)), nil
}
