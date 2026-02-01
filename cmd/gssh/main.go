package main

import (
	"os"

	"github.com/ovitorvalente/gssh/internal/adapter"
	"github.com/ovitorvalente/gssh/internal/usecase"
	"github.com/ovitorvalente/gssh/pkg/ui"
	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}

func rootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "gssh",
		Short: "Configura SSH para desenvolvimento com GitHub",
	}
	root.AddCommand(setupCommand())
	return root
}

func setupCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "setup",
		Short: "Configura chave SSH e exibe instruções para o GitHub",
		RunE: func(cmd *cobra.Command, args []string) error {
			repo := adapter.NewKeyRepositoryFile()
			ssh := adapter.NewSSHLocalSystem()
			uc := usecase.NewSetupGitHubUseCase(repo, ssh)
			out := ui.NewPrinter(cmd.OutOrStdout())

			out.PrintHeader()

			existed := repo.Exists()
			if existed {
				out.PrintKeyExists()
			} else {
				out.PrintGenerating()
			}

			result, err := uc.Execute()
			if err != nil {
				out.PrintError("Erro ao configurar SSH: %v", err)
				return err
			}

			if !existed {
				out.PrintKeyGenerated()
			}

			if result.AddKeyWarn != nil {
				out.PrintAddKeyWarning(result.AddKeyWarn)
			}

			out.PrintConfigurationCompleted(result.PublicKey.Content)
			return nil
		},
	}
}
