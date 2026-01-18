#!/usr/bin/env bash
# =============================================================================
# vscode_setup.sh
# Copies & registers the Monkey/Magpie VS Code extension locally
# =============================================================================

set -euo pipefail

# ── Configuration ────────────────────────────────────────────────────────────

EXTENSION_SRC_DIR="$(pwd)/misc/vscode/mk"
PUBLISHER="hmza"
EXTENSION_NAME="monkey-magpie"
VERSION="0.1.0"                      # ← change when you update package.json

TARGET_DIR="$HOME/.vscode/extensions/${PUBLISHER}.${EXTENSION_NAME}-${VERSION}"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# ── Helper functions ─────────────────────────────────────────────────────────

print_step() {
    printf "\n${GREEN}▶ %s${NC}\n" "$1"
}

print_warning() {
    printf "${YELLOW}⚠  %s${NC}\n" "$1"
}

print_error() {
    printf "${RED}✖  %s${NC}\n" "$1" >&2
    exit 1
}

# ── Main logic ───────────────────────────────────────────────────────────────

print_step "VS Code Monkey/Magpie extension local setup"

# 1. Check prerequisites
if [[ ! -d "$EXTENSION_SRC_DIR" ]]; then
    print_error "Source directory not found: $EXTENSION_SRC_DIR"
fi

if [[ ! -f "$EXTENSION_SRC_DIR/package.json" ]]; then
    print_error "package.json not found in $EXTENSION_SRC_DIR"
fi

# 2. Create parent directory if needed
mkdir -p "$HOME/.vscode/extensions" 2>/dev/null || true

# 3. Remove old version if exists (prevents conflicts)
if [[ -d "$TARGET_DIR" ]]; then
    print_step "Removing previous version..."
    rm -rf "$TARGET_DIR"
fi

# 4. Copy the extension files
print_step "Copying extension to:"
printf "   %s\n" "$TARGET_DIR"

cp -r "$EXTENSION_SRC_DIR" "$TARGET_DIR"

# 5. Final information
printf "\n${GREEN}Installation finished!${NC}\n\n"

cat <<EOF
Extension location:
  → $TARGET_DIR

Next steps:
  1. Reload VS Code window                 (Ctrl+Shift+P → Developer: Reload Window)
  2. Open any .magpie file (or configured extension)
  3. Syntax highlighting should be active

For development (recommended while working on the extension):
  • Open the folder in VS Code:
      code "$EXTENSION_SRC_DIR"
  • Press  F5   → starts Extension Development Host window
  • Changes → Ctrl+R (reload) in dev window

To update later:
  • Just re-run this script (it removes old version automatically)

Happy coding with Monkey/Magpie! 🐒
EOF

exit 0