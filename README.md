# gssh

CLI para configurar chaves SSH e preparar o ambiente de desenvolvimento para uso com GitHub.

---

## Sobre o projeto

O **gssh** simplifica a configuração de chaves SSH para desenvolvedores que precisam autenticar no GitHub. Com um único comando, você gera a chave, adiciona ao ssh-agent e recebe instruções para cadastrá-la no GitHub.

**Características:**

- Fluxo guiado com steps e feedback visual
- Cópia automática da chave para o clipboard
- Suporte a Linux e macOS (amd64 e arm64)
- Interface com cores e animações no terminal

---

## Instalação

### Script automático (recomendado)

Detecta automaticamente Linux/macOS e arquitetura (amd64/arm64):

```bash
curl -fsSL https://raw.githubusercontent.com/ovitorvalente/gssh/main/install.sh | bash
```

**Opções do script:**

| Opção       | Descrição                                      |
| ----------- | ---------------------------------------------- |
| `--local`   | Instala em `~/.local/bin` (padrão, sem sudo)   |
| `--system`  | Instala em `/usr/local/bin` (requer sudo)      |
| `--version` | Versão específica (ex: `--version v1.0.0`)     |
| `--help`    | Exibe ajuda                                    |

Exemplos:

```bash
curl -fsSL ... | bash                  # instala em ~/.local/bin
curl -fsSL ... | bash -s -- --system   # instala em /usr/local/bin
```

---

### Binário manual

Baixe o binário para sua plataforma na [página de releases](https://github.com/ovitorvalente/gssh/releases):

**Linux (amd64):**

```bash
curl -fsSL https://github.com/ovitorvalente/gssh/releases/latest/download/gssh-linux-amd64 -o gssh
chmod +x gssh
sudo mv gssh /usr/local/bin/
```

**macOS (Apple Silicon):**

```bash
curl -fsSL https://github.com/ovitorvalente/gssh/releases/latest/download/gssh-darwin-arm64 -o gssh
chmod +x gssh
sudo mv gssh /usr/local/bin/
```

---

### Via Go

Requer Go instalado:

```bash
go install github.com/ovitorvalente/gssh/cmd/gssh@latest
```

> Certifique-se de que `$HOME/go/bin` esteja no seu `PATH`.

---

## Uso

```bash
gssh run       # configura SSH e exibe instruções
gssh help      # documentação
gssh version   # versão, commit e data de build
```

O comando `gssh run` irá:

1. Verificar se existe chave SSH em `~/.ssh/id_ed25519`
2. Gerar nova chave (ed25519) caso não exista
3. Adicionar a chave ao ssh-agent (quando disponível)
4. Copiar a chave pública para o clipboard (opcional)
5. Exibir a chave e instruções para adicioná-la no GitHub

**Flag `--no-copy`:** evita copiar a chave para o clipboard.

---

## Requisitos

- Linux ou macOS
- `ssh-keygen` e `ssh-add` (incluídos no OpenSSH)

---

## Desenvolvimento

### Build local

```bash
git clone https://github.com/ovitorvalente/gssh.git
cd gssh
make build    # build com versão dev
make test     # executa testes
make install  # instala em $GOPATH/bin
```

### Estrutura do projeto

```
gssh/
├── cmd/gssh/           # ponto de entrada
├── internal/           # lógica interna
│   ├── adapter/        # implementações (arquivo, ssh)
│   ├── config/         # configurações
│   ├── domain/         # entidades
│   ├── port/           # interfaces
│   ├── usecase/        # casos de uso
│   └── version/        # versão injetada no build
├── pkg/
    ├── clipboard/      # cópia para clipboard
    └── ui/             # interface de linha de comando
```

### Versionamento e releases

1. Crie uma tag: `git tag v1.0.0`
2. Faça push: `git push origin v1.0.0`
3. O GitHub Actions builda e publica automaticamente binários para Linux, macOS e Windows

---

## Contribuindo

Contribuições são bem-vindas.

### Como contribuir

1. Faça um **fork** do repositório
2. Crie uma **branch** para sua feature: `git checkout -b feat/minha-feature`
3. Faça **commit** das mudanças seguindo [Conventional Commits](https://www.conventionalcommits.org/):
   - `feat:` nova funcionalidade
   - `fix:` correção de bug
   - `docs:` documentação
   - `refactor:` refatoração
4. Faça **push** para a branch: `git push origin feat/minha-feature`
5. Abra um **Pull Request** descrevendo as alterações

### Padrões de código

- Código em **inglês** (variáveis, funções, comentários)
- Mensagens ao usuário em **português**
- Máximo **4 arquivos** por commit (divida em commits atômicos)
- Execute os testes antes de abrir o PR: `go test ./...`

### Reportar bugs

Abra uma [issue](https://github.com/ovitorvalente/gssh/issues) com:

- Descrição do problema
- Passos para reproduzir
- Sistema operacional e versão
- Saída de `gssh version`

---

## Licença

MIT
