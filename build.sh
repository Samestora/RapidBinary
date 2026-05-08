#!/bin/bash

BUILD_DIR="build"
APPS=("rbserver" "rbhash")

case "$1" in
    build)
        echo "Building for current OS..."
        mkdir -p $BUILD_DIR
        for APP in "${APPS[@]}"; do
            go build -ldflags="-s -w" -o "$BUILD_DIR/$APP" "./cmd/$APP"
            echo "  - $APP ready"
        done
        ;;
    release)
        echo "Cross-compiling RapidBinary Toolsuite..."
        mkdir -p $BUILD_DIR
        # Target: GOOS/GOARCH
        PLATFORMS=("windows/amd64" "linux/amd64" "darwin/arm64")

        for PLATFORM in "${PLATFORMS[@]}"; do
            GOOS=${PLATFORM%/*}
            GOARCH=${PLATFORM#*/}
            EXT=""
            if [ "$GOOS" == "windows" ]; then EXT=".exe"; fi

            for APP in "${APPS[@]}"; do
                echo "  - Building $APP for $GOOS/$GOARCH"
                GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o "$BUILD_DIR/${APP}_${GOOS}_${GOARCH}${EXT}" "./cmd/$APP"
            done
        done
        echo -e "\nRelease binaries ready in /$BUILD_DIR"
        ;;
    clean)
        if [ -d "$BUILD_DIR" ]; then
            rm -rf "$BUILD_DIR"
            echo "Cleaned."
        else
            echo "Nothing to clean."
        fi
        ;;
    *)
        echo "Usage: ./build.sh {build|release|clean}"
        exit 1
        ;;
esac
