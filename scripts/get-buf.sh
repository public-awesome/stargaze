# Substitute BIN for your bin directory.
# Substitute VERSION for the current released version.
BIN="bin/"
VERSION="1.28.1"
URL="https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)"
echo "$URL"
curl -sSL "$URL" -o "${BIN}/buf" 
chmod +x "${BIN}/buf"
echo "buf installed to ${BIN}/buf"
