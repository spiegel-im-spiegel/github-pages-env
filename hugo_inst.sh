#!/usr/bin/env bash
set -euo pipefail

# Install latest Hugo Extended .deb from GitHub Releases using apt.
# Usage: ./hugo_inst.sh

REPO="gohugoio/hugo"
API_URL="https://api.github.com/repos/${REPO}/releases/latest"
WORK_DIR="${TMPDIR:-/tmp}/hugo-inst.$$"

cleanup() {
  rm -rf "$WORK_DIR"
}
trap cleanup EXIT

mkdir -p "$WORK_DIR"
cd "$WORK_DIR"

arch="$(dpkg --print-architecture)"
case "$arch" in
  amd64|arm64)
    ;;
  *)
    echo "Unsupported architecture: $arch" >&2
    echo "This script supports: amd64, arm64" >&2
    exit 1
    ;;
esac

echo "Current Hugo:"
if command -v hugo >/dev/null 2>&1; then
  hugo version || true
else
  echo "hugo command not found"
fi

echo "Fetching latest release metadata from ${REPO} ..."
release_json="$(curl -fsSL "$API_URL")"

version="$(awk -F '"' '/"tag_name"/ {print $4; exit}' <<<"$release_json")"
if [[ -z "$version" ]]; then
  echo "Failed to read latest release version." >&2
  exit 1
fi

asset_url="$(awk -F '"' '/"browser_download_url"/ && /hugo_extended_.*_linux-'"$arch"'\.deb/ {print $4; exit}' <<<"$release_json")"
if [[ -z "$asset_url" ]]; then
  echo "Failed to find hugo_extended .deb for architecture: $arch" >&2
  exit 1
fi

deb_file="${asset_url##*/}"

echo "Latest release: ${version}"
echo "Downloading: ${deb_file}"
curl -fL --retry 3 --retry-delay 2 -o "$deb_file" "$asset_url"

# Optional checksum verification when checksums.txt exists in latest release.
checksums_url="$(awk -F '"' '/"browser_download_url"/ && /checksums\.txt/ {print $4; exit}' <<<"$release_json")"
if [[ -n "$checksums_url" ]]; then
  echo "Downloading checksums.txt for verification ..."
  curl -fL --retry 3 --retry-delay 2 -o checksums.txt "$checksums_url"

  if awk '{print $2}' checksums.txt | grep -qx "$deb_file"; then
    echo "Verifying checksum ..."
    grep "  $deb_file$" checksums.txt | sha256sum -c -
  else
    echo "checksums.txt does not include ${deb_file}; skipping verification."
  fi
else
  echo "checksums.txt not found in latest release; skipping verification."
fi

echo "Installing ${deb_file} with apt ..."
sudo apt install -y "./$deb_file"

echo "Installed Hugo:"
hugo version

echo "Done."
