package adapter

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/ovitorvalente/gssh/internal/config"
)

type SSHLocalSystem struct {
	pathKey string
}

func NewSSHLocalSystem() *SSHLocalSystem {
	pathKey, _ := config.SSHPaths()
	return &SSHLocalSystem{
		pathKey: pathKey,
	}
}

func (s *SSHLocalSystem) AddKey() error {
	if _, err := os.Stat(s.pathKey); os.IsNotExist(err) {
		return fmt.Errorf("chave SSH não encontrada")
	}
	cmd := exec.Command("ssh-add", s.pathKey)
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	return cmd.Run()
}
