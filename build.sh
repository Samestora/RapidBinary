#!/bin/bash

# Configuration
BINARY_NAME="rbserver"
BUILD_DIR="build"
MAIN_PATH="./cmd/rbserver"

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m'

show_help() {
    echo "Usage: ./build.sh [command]"
    echo ""
    echo "Commands:"
    echo "  build    Compile for current OS"
    echo "  release  Cross-compile for Windows, Linux, and Mac"
    echo "  clean    Remove the build directory"
}

build_current() {
    echo -e "${BLUE}Building ${BINARY_NAME}...${NC}"
    mkdir -p $BUILD_DIR
    go build -o "${BUILD_DIR}/${BINARY_NAME}" $MAIN_PATH
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}Done! Binary located at: ${BUILD_DIR}/${BINARY_NAME}${NC}"
    else
        echo -e "${RED}❌ Build failed.${NC}"
        exit 1
    fi
}

build_release() {
    echo -e "${BLUE}Cross-compiling for all platforms...${NC}"
    mkdir -p $BUILD_DIR

    echo "Building Windows..."
    GOOS=windows GOARCH=amd64 go build -o "${BUILD_DIR}/rbserver_windows_amd64.exe" $MAIN_PATH

    echo "Building Linux..."
    GOOS=linux GOARCH=amd64 go build -o "${BUILD_DIR}/rbserver_linux_amd64" $MAIN_PATH

    echo "Building MacOS..."
    GOOS=darwin GOARCH=arm64 go build -o "${BUILD_DIR}/rbserver_macos_arm64" $MAIN_PATH

    echo -e "${GREEN}Release binaries ready in ${BUILD_DIR}${NC}"
}

case "$1" in
    build) build_current ;;
    release) build_release ;;
    clean)
        rm -rf $BUILD_DIR
        echo "Cleaned."
        ;;
    *) show_help ;;
esac
