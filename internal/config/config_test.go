package config

import (
	"path/filepath"
	"testing"
)

func TestSSHPaths_RetornaCaminhosValidos(t *testing.T) {
	keyPath, pubKeyPath := SSHPaths()

	if keyPath == "" {
		t.Error("keyPath não deve ser vazio")
	}
	if pubKeyPath == "" {
		t.Error("pubKeyPath não deve ser vazio")
	}
	if filepath.Ext(keyPath) != "" {
		t.Error("keyPath não deve ter extensão")
	}
	if filepath.Ext(pubKeyPath) != ".pub" {
		t.Error("pubKeyPath deve terminar em .pub")
	}
	if pubKeyPath != keyPath+".pub" {
		t.Errorf("pubKeyPath (%s) deve ser keyPath+.pub (%s)", pubKeyPath, keyPath+".pub")
	}
}
