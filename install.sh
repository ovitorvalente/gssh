#!/usr/bin/env bash

set -e

BIN_NAME="gssh"
REPO="ovitorvalente/gssh"
INSTALL_DIR=""
USE_LOCAL=true
VERSION="latest"

usage() {
  echo "Uso: $0 [OPÇÕES]"
  echo ""
  echo "Opções:"
  echo "  --local    Instala em ~/.local/bin (padrão, sem sudo)"
  echo "  --system  Instala em /usr/local/bin (requer sudo)"
  echo "  --version  Instala versão específica (ex: v1.0.0, padrão: latest)"
  echo "  -h, --help Exibe esta ajuda"
}

while [[ $# -gt 0 ]]; do
  case $1 in
    --local)
      USE_LOCAL=true
      shift
      ;;
    --system)
      USE_LOCAL=false
      shift
      ;;
    --version)
      if [[ -z "${2:-}" ]] || [[ "$2" == --* ]]; then
        echo "❌ --version requer um valor (ex: v1.0.0)"
        exit 1
      fi
      VERSION="$2"
      shift 2
      ;;
    -h|--help)
      usage
      exit 0
      ;;
    *)
      echo "Opção desconhecida: $1"
      usage
      exit 1
      ;;
  esac
done

detect_platform() {
  local os
  local arch

  case "$(uname -s)" in
    Linux)
      os="linux"
      ;;
    Darwin)
      os="darwin"
      ;;
    *)
      echo "❌ Sistema operacional não suportado: $(uname -s)"
      exit 1
      ;;
  esac

  case "$(uname -m)" in
    x86_64|amd64)
      arch="amd64"
      ;;
    aarch64|arm64)
      arch="arm64"
      ;;
    *)
      echo "❌ Arquitetura não suportada: $(uname -m)"
      exit 1
      ;;
  esac

  echo "${os}-${arch}"
}

set_install_dir() {
  if [[ "$USE_LOCAL" == true ]]; then
    INSTALL_DIR="${HOME}/.local/bin"
    mkdir -p "$INSTALL_DIR"
  else
    INSTALL_DIR="/usr/local/bin"
  fi
}

resolve_url() {
  local platform
  platform=$(detect_platform)

  if [[ "$VERSION" == "latest" ]]; then
    echo "https://github.com/${REPO}/releases/latest/download/gssh-${platform}"
  else
    echo "https://github.com/${REPO}/releases/download/${VERSION}/gssh-${platform}"
  fi
}

check_existing() {
  local existing
  existing=$(command -v gssh 2>/dev/null || true)
  if [[ -n "$existing" ]] && [[ -t 0 ]]; then
    echo "⚠️  gssh já instalado em: $existing"
    read -r -p "Sobrescrever? [s/N] " response
    if [[ ! "$response" =~ ^[sSyY]$ ]]; then
      echo "Instalação cancelada."
      exit 0
    fi
  fi
}

main() {
  set_install_dir
  check_existing

  local url
  url=$(resolve_url)

  echo "⬇️  Baixando gssh..."
  if ! curl -fsSL "$url" -o "$BIN_NAME"; then
    echo "❌ Download falhou. Verifique se a release existe: $url"
    exit 1
  fi

  if [[ ! -s "$BIN_NAME" ]]; then
    echo "❌ Arquivo baixado está vazio."
    rm -f "$BIN_NAME"
    exit 1
  fi

  chmod +x "$BIN_NAME"

  if [[ "$USE_LOCAL" == true ]]; then
    mv "$BIN_NAME" "$INSTALL_DIR"
  else
    echo "📦 Instalando em $INSTALL_DIR (sudo necessário)"
    sudo mv "$BIN_NAME" "$INSTALL_DIR"
  fi

  echo "✅ gssh installed successfully!"
  echo ""
  echo "  Run:     gssh run"
  echo "  Help:    gssh help"
  echo "  Version: gssh version"
  echo ""

  if [[ "$USE_LOCAL" == true ]]; then
    if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
      echo "💡 Adicione ao seu PATH (adicione ao ~/.bashrc ou ~/.zshrc):"
      echo "   export PATH=\"\$HOME/.local/bin:\$PATH\""
      echo ""
    fi
  fi
}

main
