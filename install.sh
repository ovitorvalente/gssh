#!/usr/bin/env bash

set -e

BIN_NAME="gssh"
INSTALL_DIR="/usr/local/bin"
URL="https://github.com/ovitorvalente/gssh/releases/latest/download/gssh-linux-amd64"

echo "⬇️ Baixando o binário do gssh..."
curl -fsSl "$URL" -o "$BIN_NAME"

chmod +x "$BIN_NAME"

echo "📦 Instalando o gssh no diretório $INSTALL_DIR (sudo requer senha)"
sudo mv "$BIN_NAME" "$INSTALL_DIR"

echo "✅ gssh instalado com sucesso!"
echo "🔗 Para usar, execute: gssh setup"
echo "💡 Para ver a documentação: gssh help"