package ui

const (
	GithubSSHURL  = "https://github.com/settings/ssh/new"
	Header        = "  🔐  GitHub SSH Setup"
	Separator     = "  ───────────────────"
	KeyExists     = "  ✔ Chave SSH encontrada"
	GeneratingKey = "  ➜ Gerando nova chave SSH..."
	KeyGenerated  = "  ✔ Chave gerada com sucesso"
	ConfigDone    = "  ✔ Tudo pronto! Adicione a chave no GitHub:"
	NextSteps     = ""
	Step1         = "  1. Copie a chave pública (linha abaixo):"
	Step2         = "  2. Abra este link no navegador:"
	Step3         = "  3. Cole a chave e clique em \"Add SSH key\""
	Step4         = "  4. Teste a conexão:"
	TestCommand   = "     $ ssh -T git@github.com"
	WarningAddKey = "  ⚠ Não foi possível adicionar ao ssh-agent (pode ignorar)"
	ErrorPrefix   = "  ❌"
)
