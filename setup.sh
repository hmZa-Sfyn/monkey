#!/usr/bin/env bash
set -euo pipefail

interpreter_name="monkey"
binary_name="monkey"
go_file="./main.go"

# Detect OS
OS="$(uname -s)"
case "$OS" in
  Linux*)   PLATFORM="linux" ;;
  Darwin*)  PLATFORM="darwin" ;;
  MINGW*|MSYS*|CYGWIN*) PLATFORM="windows" ;;
  *)        PLATFORM="unknown" ;;
esac

# Determine user home and target bin path
if [[ "$PLATFORM" == "windows" ]]; then
  USER_HOME="$(cmd.exe /c "echo %USERPROFILE%" 2>/dev/null | tr -d '\r')"
  BIN_DIR="$USER_HOME/.mk/bin"
  binary_name="monkey.exe"
else
  USER_HOME="$HOME"
  BIN_DIR="$HOME/.mk/bin"
fi

mkdir -p "$BIN_DIR"

# 1) Build binary
echo "Building $interpreter_name..."

if [[ "$PLATFORM" == "windows" ]]; then
  env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o "$BIN_DIR/$binary_name" "$go_file"
else
  env GOOS="$PLATFORM" GOARCH=amd64 go build -ldflags "-s -w" -o "$BIN_DIR/$binary_name" "$go_file"
fi

echo "Binary installed to: $BIN_DIR/$binary_name"

# 2) Create PATH.yaml (different for windows and linux)
CFG_DIR="$USER_HOME/.mk/cfg"
mkdir -p "$CFG_DIR"

if [[ "$PLATFORM" == "windows" ]]; then
  STD_PATH="$USER_HOME/.mk/lib/std"
else
  STD_PATH="$USER_HOME/.mk/lib/std"
fi

cat > "$CFG_DIR/PATH.yaml" <<EOF
std: $STD_PATH
usr: $USER_HOME/.mk/lib/usr
tmp: /tmp
EOF

echo "PATH.yaml created at: $CFG_DIR/PATH.yaml"

# 3) Add bin to PATH
if [[ "$PLATFORM" == "windows" ]]; then
  # Windows: just show message (can't modify system env safely)
  echo "Please add this to your PATH manually (Windows):"
  echo "  $BIN_DIR"
else
  # Linux: modify bashrc/zshrc safely (remove old duplicates)
  add_to_shell_config() {
    local rc_file="$1"
    local marker_start="# >>> monkey config start >>>"
    local marker_end="# <<< monkey config end <<<"

    # Remove existing block
    if grep -q "$marker_start" "$rc_file"; then
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

  # Update bashrc and zshrc
  if [[ -f "$HOME/.bashrc" ]]; then
    add_to_shell_config "$HOME/.bashrc"
  fi
  if [[ -f "$HOME/.zshrc" ]]; then
    add_to_shell_config "$HOME/.zshrc"
  fi

  echo "Updated ~/.bashrc and ~/.zshrc (if they exist)."
  echo "Reload shell or run: source ~/.bashrc"
fi

echo "Done!"
