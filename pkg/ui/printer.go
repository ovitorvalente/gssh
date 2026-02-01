package ui

import (
	"fmt"
	"io"
	"strings"
)

type Output interface {
	PrintHeader()
	PrintKeyExists()
	PrintGenerating()
	PrintKeyGenerated()
	PrintConfigurationCompleted(publicKey string, copied bool, noCopy bool)
	PrintError(format string, args ...interface{})
	PrintAddKeyWarning(err error)
	PrintKeyCopiedClipboard()
	PrintClipboardUnavailable()
	PrintOpenBrowserMessage()
	PrintBrowserOpenFailed()
}

type Printer struct {
	out io.Writer
}

func NewPrinter(out io.Writer) *Printer {
	return &Printer{out: out}
}

func (p *Printer) PrintHeader() {
	fmt.Fprintln(p.out)
	fmt.Fprintln(p.out, styleHeader.Render("  🔐  GitHub SSH Setup  "))
	fmt.Fprintln(p.out, styleSeparator.Render("  ───────────────────────────────────"))
	fmt.Fprintln(p.out)
}

func (p *Printer) PrintKeyExists() {
	fmt.Fprintln(p.out, styleOk.Render("  ✔ ")+styleMessage.Render("Chave SSH encontrada"))
}

func (p *Printer) PrintKeyGenerated() {
	fmt.Fprintln(p.out, styleOk.Render("  ✔ ")+styleMessage.Render("Chave gerada com sucesso"))
	fmt.Fprintln(p.out)
}

func (p *Printer) PrintGenerating() {
	fmt.Fprintln(p.out, styleMessage.Render("  ➜ Gerando nova chave SSH..."))
}

func (p *Printer) PrintConfigurationCompleted(publicKey string, copied bool, noCopy bool) {
	key := strings.TrimSpace(publicKey)
	fmt.Fprintln(p.out)
	fmt.Fprintln(p.out, styleOk.Render("  ✔ Tudo pronto! ")+styleMessage.Render("Adicione a chave no GitHub:"))
	fmt.Fprintln(p.out)
	if copied {
		fmt.Fprintln(p.out, styleOk.Render("  ✔ Chave copiada para área de transferência"))
	} else if !noCopy {
		fmt.Fprintln(p.out, styleMessage.Render("  1. Copie a chave pública (linha abaixo):"))
		fmt.Fprintln(p.out, styleWarn.Render("  ⚠ Não foi possível copiar (copie manualmente)"))
	} else {
		fmt.Fprintln(p.out, styleMessage.Render("  1. Copie a chave pública (linha abaixo):"))
	}
	fmt.Fprintf(p.out, "     %s\n", styleKey.Render(key))
	fmt.Fprintln(p.out)
	fmt.Fprintln(p.out, styleMessage.Render("  2. Abra este link no navegador:"))
	fmt.Fprintln(p.out, "     "+styleLink.Render(GithubSSHURL))
	fmt.Fprintln(p.out)
	fmt.Fprintln(p.out, styleMessage.Render("  3. Cole a chave e clique em \"Add SSH key\""))
	fmt.Fprintln(p.out)
	fmt.Fprintln(p.out, styleMessage.Render("  4. Teste a conexão:"))
	fmt.Fprintln(p.out, "     "+styleCommand.Render("$ ssh -T git@github.com"))
	fmt.Fprintln(p.out)
}

func (p *Printer) PrintError(format string, args ...interface{}) {
	fmt.Fprintln(p.out, styleFail.Render("  ❌")+" "+styleMessage.Render(fmt.Sprintf(format, args...)))
}

func (p *Printer) PrintAddKeyWarning(err error) {
	fmt.Fprintln(p.out, styleWarn.Render("  ⚠")+" "+styleMessage.Render("Não foi possível adicionar ao ssh-agent"))
	fmt.Fprintln(p.out, "    ", err)
}

func (p *Printer) PrintKeyCopiedClipboard() {
	fmt.Fprintln(p.out, styleOk.Render("  ✔ Chave copiada para área de transferência"))
}

func (p *Printer) PrintClipboardUnavailable() {
	fmt.Fprintln(p.out, styleWarn.Render("  ⚠ Não foi possível copiar para área de transferência (copie manualmente)"))
}

func (p *Printer) PrintOpenBrowserMessage() {
	fmt.Fprintln(p.out, styleMessage.Render("  Abrindo navegador para adicionar chave..."))
}

func (p *Printer) PrintBrowserOpenFailed() {
	fmt.Fprintln(p.out, styleWarn.Render("  ⚠ Não foi possível abrir o navegador (copie manualmente)"))
}

func (p *Printer) PrintHelp() {
	fmt.Fprintln(p.out)
	fmt.Fprintln(p.out, styleHeader.Render("  📖  gssh — Ajuda  "))
	fmt.Fprintln(p.out, styleSeparator.Render("  ───────────────────────────────────"))
	fmt.Fprintln(p.out)
	fmt.Fprintln(p.out, styleMessage.Render(HelpDescription))
	fmt.Fprintln(p.out)
	fmt.Fprintln(p.out, styleMessage.Render(HelpUsage))
	fmt.Fprintln(p.out)
	fmt.Fprintln(p.out, styleMessage.Render(HelpUsageExample))
	fmt.Fprintln(p.out)
	fmt.Fprintln(p.out, styleMessage.Render(HelpCommands))
	fmt.Fprintln(p.out)
	fmt.Fprintln(p.out, styleMessage.Render(HelpRun))
	fmt.Fprintln(p.out)
	fmt.Fprintln(p.out, styleMessage.Render(HelpMoreInfo))
	fmt.Fprintln(p.out, styleLink.Render(HelpRepo))
	fmt.Fprintln(p.out)
}
