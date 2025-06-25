#!/bin/bash

# Script to vendor golang.org/x/tools/internal/mcp

set -e

TARGET_DIR="internal/imported/x-tools"

# Get the latest release tag from GitHub API
echo "Fetching latest version of golang.org/x/tools..."
TOOLS_VERSION=$(curl -s https://api.github.com/repos/golang/tools/releases/latest | grep '"tag_name":' | cut -d'"' -f4)

if [ -z "$TOOLS_VERSION" ]; then
    echo "Failed to fetch latest version, falling back to master branch"
    TOOLS_VERSION="master"
else
    echo "Found latest version: ${TOOLS_VERSION}"
fi

echo "Vendoring golang.org/x/tools/internal/mcp@${TOOLS_VERSION} to ${TARGET_DIR}..."

# Create temporary directory
TEMP_DIR=$(mktemp -d)
trap "rm -rf $TEMP_DIR" EXIT

# Clone golang.org/x/tools
cd $TEMP_DIR
git clone --depth 1 --branch $TOOLS_VERSION https://github.com/golang/tools.git

# Copy internal packages to x-tools structure
cp -r tools/internal/mcp "${OLDPWD}/${TARGET_DIR}/mcp"
cp -r tools/internal/jsonrpc2_v2 "${OLDPWD}/${TARGET_DIR}/jsonrpc2_v2"
cp -r tools/internal/xcontext "${OLDPWD}/${TARGET_DIR}/xcontext"
cp -r tools/internal/testenv "${OLDPWD}/${TARGET_DIR}/testenv"

# Update package name and import paths
cd "${OLDPWD}/${TARGET_DIR}"
find . -name "*.go" -exec sed -i.bak 's|golang.org/x/tools/internal/mcp|github.com/takak2166/scrapbox-mcp/internal/imported/x-tools/mcp|g' {} \;
find . -name "*.go" -exec sed -i.bak 's|golang.org/x/tools/internal/jsonrpc2_v2|github.com/takak2166/scrapbox-mcp/internal/imported/x-tools/jsonrpc2_v2|g' {} \;
find . -name "*.go" -exec sed -i.bak 's|golang.org/x/tools/internal/xcontext|github.com/takak2166/scrapbox-mcp/internal/imported/x-tools/xcontext|g' {} \;
find . -name "*.go" -exec sed -i.bak 's|golang.org/x/tools/internal/testenv|github.com/takak2166/scrapbox-mcp/internal/imported/x-tools/testenv|g' {} \;
find . -name "*.go.bak" -delete

echo "Vendoring completed: ${TARGET_DIR}"
echo "To add new implementation: create cmd/x-tools-mcp/ and internal/x-tools-mcp/" 
