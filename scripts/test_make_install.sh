#!/usr/bin/env sh
set -eu

tmpdir="$(mktemp -d)"
trap 'rm -rf "$tmpdir"' EXIT INT TERM

install_dir="$tmpdir/bin"

make install INSTALL_DIR="$install_dir" SUDO=

if [ ! -x "$install_dir/tmdb" ]; then
  echo "expected executable at $install_dir/tmdb" >&2
  exit 1
fi

"$install_dir/tmdb" --version | grep 'tmdb version 1.0.2' >/dev/null
