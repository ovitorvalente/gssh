package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ovitorvalente/gssh/internal/adapter"
	"github.com/ovitorvalente/gssh/internal/usecase"
	"github.com/ovitorvalente/gssh/internal/version"
	"github.com/ovitorvalente/gssh/pkg/clipboard"
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
	root.AddCommand(runCommand())
	root.AddCommand(helpCommand())
	root.AddCommand(versionCommand())
	return root
}

func versionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Exibe informações sobre a versão do projeto",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("gssh %s\n", version.Version)
			fmt.Printf("  commit: %s\n", version.Commit)
			fmt.Printf("  built:  %s\n", version.BuildDate)
			return nil
		},
	}
}

func helpCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "help",
		Short: "Exibe documentação e instruções de uso",
		RunE: func(cmd *cobra.Command, args []string) error {
			out := ui.NewPrinter(cmd.OutOrStdout())
			out.PrintHelp()
			return nil
		},
	}
}

func runCommand() *cobra.Command {
	var noCopy bool

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Configura chave SSH e exibe instruções para o GitHub",
		RunE: func(c *cobra.Command, args []string) error {
			uc := newSetupUseCase()
			printer := ui.NewPrinter(c.OutOrStdout())
			stepper := ui.NewStepper(c.OutOrStdout(), 3)

			printer.PrintHeader()

			// Step 1: Verificar/Gerar chave
			stepper.StartStep(ui.StepCheckingKey)
			existed, err := uc.EnsureKeyExists()
			if err != nil {
				stepper.Fail("Erro ao gerar chave")
				printer.PrintError("Erro ao gerar chave: %v", err)
				return err
			}
			if existed {
				stepper.Skip(ui.StepKeyFound)
			} else {
				stepper.Ok(ui.StepKeyGenerated, 0)
			}

			// Step 2: Adicionar ao ssh-agent
			stepper.StartStep(ui.StepAddingAgent)
			start := time.Now()
			addKeyErr := uc.AddKeyToAgent()
			if addKeyErr != nil {
				stepper.Warn(ui.StepAgentWarn)
			} else {
				stepper.Ok(ui.StepAgentOk, time.Since(start))
			}

			// Step 3: Ler chave pública
			stepper.StartStep(ui.StepReadingKey)
			start = time.Now()
			pub, err := uc.GetPublicKey()
			if err != nil {
				stepper.Fail("Erro ao ler chave")
				printer.PrintError("Erro ao ler chave: %v", err)
				return err
			}
			stepper.Ok(ui.StepKeyRead, time.Since(start))

			// Copiar para clipboard
			copied := false
			if !noCopy {
				copied = clipboard.Copy(pub)
			}

			printer.PrintConfigurationCompleted(pub, copied, noCopy)
			return nil
		},
	}

	cmd.Flags().BoolVar(&noCopy, "no-copy", false, "não copia a chave para área de transferência")
	return cmd
}

func newSetupUseCase() *usecase.SetupGitHubUseCase {
	repo := adapter.NewKeyRepositoryFile()
	ssh := adapter.NewSSHLocalSystem()
	return usecase.NewSetupGitHubUseCase(repo, ssh)
}
