#!/usr/bin/env bash
set -euo pipefail

interpreter_name="monkey"
binary_name="monkey"
go_file="./main.go"
STD_LIBS_SOURCE_DIR="./lib"          # ← change to "./lib" if that's your real folder name

# Detect OS
OS="$(uname -s)"
case "$OS" in
  Linux*)   PLATFORM="linux" ;;
  Darwin*)  PLATFORM="darwin" ;;
  MINGW*|MSYS*|CYGWIN*) PLATFORM="windows" ;;
  *)        PLATFORM="unknown" ;;
esac

# Determine user home and target paths
if [[ "$PLATFORM" == "windows" ]]; then
  USER_HOME="$(cmd.exe /c "echo %USERPROFILE%" 2>/dev/null | tr -d '\r')"
  BIN_DIR="$USER_HOME/.mk/bin"
  LIB_DIR="$USER_HOME/.mk/lib"
  binary_name="monkey.exe"
else
  USER_HOME="$HOME"
  BIN_DIR="$HOME/.mk/bin"
  LIB_DIR="$HOME/.mk/lib"
fi

STD_LIB_TARGET="$LIB_DIR/std"

mkdir -p "$BIN_DIR"
mkdir -p "$LIB_DIR"
mkdir -p "$STD_LIB_TARGET"
mkdir -p "$LIB_DIR/usr"

# 1) Build binary
echo "Building $interpreter_name..."

if [[ "$PLATFORM" == "windows" ]]; then
  env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o "$BIN_DIR/$binary_name" "$go_file"
else
  env GOOS="$PLATFORM" GOARCH=amd64 go build -ldflags "-s -w" -o "$BIN_DIR/$binary_name" "$go_file"
fi

echo "Binary installed to: $BIN_DIR/$binary_name"

# 2) Copy standard library files
if [[ -d "$STD_LIBS_SOURCE_DIR" ]]; then
  echo "Copying standard library files from $STD_LIBS_SOURCE_DIR → $STD_LIB_TARGET ..."
  
  # Copy everything from ./libs/ to ~/.mk/lib/std/
  # -a = archive mode (preserves permissions, times, etc)
  # -v = verbose
  cp -av "$STD_LIBS_SOURCE_DIR"/* "$STD_LIB_TARGET"/ 2>/dev/null || true
  
  # Alternative (more explicit / safer in some cases):
  # find "$STD_LIBS_SOURCE_DIR" -maxdepth 1 -type f -name "*.mk" -exec cp -v {} "$STD_LIB_TARGET"/ \;
  
  echo "Standard library files copied to: $STD_LIB_TARGET"
else
  echo "Warning: $STD_LIBS_SOURCE_DIR not found — skipping std lib copy"
fi

# 3) Create PATH.yaml
CFG_DIR="$USER_HOME/.mk/cfg"
mkdir -p "$CFG_DIR"

cat > "$CFG_DIR/PATH.yaml" <<EOF
std: $STD_LIB_TARGET
usr: $LIB_DIR/usr
monkey: $LIB_DIR/std/monkey
tmp: /tmp
EOF

echo "PATH.yaml created at: $CFG_DIR/PATH.yaml"

# 4) Add bin to PATH
if [[ "$PLATFORM" == "windows" ]]; then
  echo "Please add this to your PATH manually (Windows):"
  echo "  $BIN_DIR"
  echo ""
  echo "Standard library location:"
  echo "  $STD_LIB_TARGET"
  echo "  (usually: %USERPROFILE%\\.mk\\lib\\std)"
else
  add_to_shell_config() {
    local rc_file="$1"
    local marker_start="# >>> monkey config start >>>"
    local marker_end="# <<< monkey config end <<<"

    # Remove existing block
    if grep -q "$marker_start" "$rc_file" 2>/dev/null; then
      sed -i "/$marker_start/,/$marker_end/d" "$rc_file"
    fi

    cat >> "$rc_file" <<EOF
$marker_start
export PATH="$BIN_DIR:\$PATH"
alias monkey="$BIN_DIR/$binary_name"
export MK_BIN="$BIN_DIR/$binary_name"
$marker_end
EOF
  }

  # Update bashrc and zshrc if they exist
  for rc in "$HOME/.bashrc" "$HOME/.zshrc"; do
    if [[ -f "$rc" ]]; then
      add_to_shell_config "$rc"
    fi
  done

  echo "Updated ~/.bashrc and/or ~/.zshrc (if they exist)."
  echo "Reload shell or run: source ~/.bashrc (or ~/.zshrc)"
fi

echo ""
echo "Installation complete!"
echo "Standard library location: $STD_LIB_TARGET"
echo "Example: ./libs/abc.mk  →  $STD_LIB_TARGET/abc.mk"
echo "Done! 🚀"