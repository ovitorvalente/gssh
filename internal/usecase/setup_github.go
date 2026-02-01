package usecase

import (
	"github.com/ovitorvalente/gssh/internal/domain"
	"github.com/ovitorvalente/gssh/internal/port"
)

type SetupResult struct {
	PublicKey  *domain.PublicKey
	AddKeyWarn error
}

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

func (s *SetupGitHubUseCase) Execute() (*SetupResult, error) {
	if !s.keyRepository.Exists() {
		if err := s.keyRepository.Generate(); err != nil {
			return nil, err
		}
	}

	addKeyWarn := s.sshSystem.AddKey()

	pub, err := s.keyRepository.ReadPublicKey()
	if err != nil {
		return nil, err
	}

	return &SetupResult{
		PublicKey:  &domain.PublicKey{Content: pub},
		AddKeyWarn: addKeyWarn,
	}, nil
}
