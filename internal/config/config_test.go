package config

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestSSHPaths_ReturnsValidPaths(t *testing.T) {
	keyPath, pubKeyPath := SSHPaths()

	if keyPath == "" {
		t.Error("keyPath should not be empty")
	}
	if pubKeyPath == "" {
		t.Error("pubKeyPath should not be empty")
	}
	if filepath.Ext(keyPath) != "" {
		t.Error("keyPath should not have extension")
	}
	if filepath.Ext(pubKeyPath) != ".pub" {
		t.Error("pubKeyPath must end with .pub extension")
	}
	if pubKeyPath != keyPath+".pub" {
		t.Errorf("pubKeyPath (%s) must equal keyPath+.pub (%s)", pubKeyPath, keyPath+".pub")
	}
}

func TestSSHPaths_ContainsExpectedSegments(t *testing.T) {
	keyPath, pubKeyPath := SSHPaths()

	expectedSegments := []string{".ssh", "id_ed25519"}
	for _, seg := range expectedSegments {
		if !strings.Contains(keyPath, seg) {
			t.Errorf("keyPath %q should contain %q", keyPath, seg)
		}
	}

	if !strings.HasSuffix(pubKeyPath, "id_ed25519.pub") {
		t.Errorf("pubKeyPath %q should end with id_ed25519.pub", pubKeyPath)
	}
}

func TestSSHPaths_KeysAreSiblings(t *testing.T) {
	keyPath, pubKeyPath := SSHPaths()

	keyDir := filepath.Dir(keyPath)
	pubDir := filepath.Dir(pubKeyPath)

	if keyDir != pubDir {
		t.Errorf("key and pub key should be in same directory: %q vs %q", keyDir, pubDir)
	}
}

func TestSSHPaths_UsesCorrectSeparator(t *testing.T) {
	keyPath, _ := SSHPaths()

	// On Windows we expect backslashes, on Unix forward slashes
	separator := string(filepath.Separator)
	if !strings.Contains(keyPath, separator) {
		t.Errorf("keyPath should use OS path separator %q: got %q", separator, keyPath)
	}
}
