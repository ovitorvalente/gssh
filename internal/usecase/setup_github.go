package usecase

import (
	"github.com/ovitorvalente/gssh/internal/port"
)

type SetupGitHubUseCase struct {
	keyRepository port.KeyRepository
	sshSystem     port.SSHSystem
}

func NewSetupGitHubUseCase(keyRepository port.KeyRepository, sshSystem port.SSHSystem) *SetupGitHubUseCase {
	return &SetupGitHubUseCase{
		keyRepository: keyRepository,
		sshSystem:     sshSystem,
	}
}

func (s *SetupGitHubUseCase) EnsureKeyExists() (existed bool, err error) {
	if s.keyRepository.Exists() {
		return true, nil
	}
	if err := s.keyRepository.Generate(); err != nil {
		return false, err
	}
	return false, nil
}

func (s *SetupGitHubUseCase) AddKeyToAgent() error {
	return s.sshSystem.AddKey()
}

func (s *SetupGitHubUseCase) GetPublicKey() (string, error) {
	return s.keyRepository.ReadPublicKey()
}
