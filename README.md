# gssh

CLI para configurar chave SSH e preparar o ambiente para desenvolvimento com GitHub.

## Instalação

```bash
go install github.com/ovitorvalente/gssh/cmd/gssh@latest
```

Ou clone e construa localmente:

```bash
git clone https://github.com/ovitorvalente/gssh.git
cd gssh
go build -o gssh ./cmd/gssh
```

## Uso

```bash
gssh setup
```

O comando irá:

1. Verificar se existe chave SSH em `~/.ssh/id_ed25519`
2. Gerar nova chave (ed25519) se não existir
3. Adicionar a chave ao ssh-agent (quando disponível)
4. Exibir a chave pública e instruções para adicionar no GitHub

## Requisitos

- `ssh-keygen` e `ssh-add` instalados (geralmente incluídos no OpenSSH)
- Git configurado com seu nome e email
