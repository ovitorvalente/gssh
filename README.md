# gssh

CLI para configurar chave SSH e preparar o ambiente para desenvolvimento com GitHub.

## Instalação

```bash
go install github.com/ovitorvalente/gssh/cmd/gssh@latest
```

Ou baixe o binário da [última release](https://github.com/ovitorvalente/gssh/releases).

## Uso

```bash
gssh run
gssh version   # exibe versão, commit e data de build
gssh help      # documentação
```

O comando irá:

1. Verificar se existe chave SSH em `~/.ssh/id_ed25519`
2. Gerar nova chave (ed25519) se não existir
3. Adicionar a chave ao ssh-agent (quando disponível)
4. Exibir a chave pública e instruções para adicionar no GitHub

## Requisitos

- `ssh-keygen` e `ssh-add` instalados (geralmente incluídos no OpenSSH)

## Desenvolvimento

### Build local

```bash
make build    # build com versão dev
make test     # executa testes
make install  # instala em $GOPATH/bin
```

### Versionamento e releases

1. Crie uma tag no formato `vX.Y.Z` (ex: `v1.0.0`)
2. Faça push da tag: `git push origin v1.0.0`
3. O GitHub Actions builda automaticamente e publica a release com binários para Linux, macOS e Windows
