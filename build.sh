#!/bin/bash
set -e

VERSION=$(git describe --tags --always 2>/dev/null || echo "0.1.0-dev")
COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

echo "Building jot-api..."
echo "  Version:    $VERSION"
echo "  Commit:     $COMMIT"
echo "  Build Time: $BUILD_TIME"
echo ""

docker build \
  --build-arg VERSION="$VERSION" \
  --build-arg GIT_COMMIT="$COMMIT" \
  --build-arg BUILD_TIME="$BUILD_TIME" \
  -t jot-api:"$VERSION" \
  -t jot-api:latest \
  .

echo ""
echo "✅ Built jot-api:$VERSION and jot-api:latest"
