package ui

import (
	"fmt"
	"io"
	"strings"
)

// Output define a interface para saída do CLI (facilita testes).
type Output interface {
	PrintHeader()
	PrintKeyExists()
	PrintGenerating()
	PrintKeyGenerated()
	PrintConfigurationCompleted(publicKey string)
	PrintError(format string, args ...interface{})
	PrintAddKeyWarning(err error)
}

type Printer struct {
	out io.Writer
}

func NewPrinter(out io.Writer) *Printer {
	return &Printer{out: out}
}

func (p *Printer) PrintHeader() {
	fmt.Fprintln(p.out)
	fmt.Fprintln(p.out, Header)
	fmt.Fprintln(p.out, Separator)
	fmt.Fprintln(p.out)
}

func (p *Printer) PrintKeyExists() {
	fmt.Fprintln(p.out, KeyExists)
}

func (p *Printer) PrintGenerating() {
	fmt.Fprintln(p.out, GeneratingKey)
}

func (p *Printer) PrintKeyGenerated() {
	fmt.Fprintln(p.out, KeyGenerated)
	fmt.Fprintln(p.out)
}

func (p *Printer) PrintConfigurationCompleted(publicKey string) {
	chave := strings.TrimSpace(publicKey)
	fmt.Fprintln(p.out)
	fmt.Fprintln(p.out, ConfigDone)
	fmt.Fprintln(p.out, NextSteps)
	fmt.Fprintln(p.out)
	fmt.Fprintln(p.out, Step1)
	fmt.Fprintf(p.out, "     %s\n", chave)
	fmt.Fprintln(p.out)
	fmt.Fprintln(p.out, Step2)
	fmt.Fprintf(p.out, "     %s\n", GithubSSHURL)
	fmt.Fprintln(p.out)
	fmt.Fprintln(p.out, Step3)
	fmt.Fprintln(p.out)
	fmt.Fprintln(p.out, Step4)
	fmt.Fprintln(p.out, TestCommand)
	fmt.Fprintln(p.out)
}

func (p *Printer) PrintError(format string, args ...interface{}) {
	fmt.Fprintln(p.out, ErrorPrefix, fmt.Sprintf(format, args...))
}

func (p *Printer) PrintAddKeyWarning(err error) {
	fmt.Fprintln(p.out, WarningAddKey)
	fmt.Fprintln(p.out, "    ", err)
}
