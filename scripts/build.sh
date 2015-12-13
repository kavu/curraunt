#!/usr/bin/env bash

set -e

REPO="github.com/kavu/curraunt"
NAME="curraunt"

REV=$( git rev-parse --short HEAD 2> /dev/null || echo "unknown" )
DATE=$( TZ=UTC date "+%Y%m%d-%H:%M:%S UTC" )
GO_VERSION=$( go version | sed -e "s/^[^0-9.]*\([0-9.]*\).*/\1/" )

LDFLAGS="
  -X '${REPO}/version.Revision=${REV}'
  -X '${REPO}/version.BuildDate=${DATE}'
  -X '${REPO}/version.GoVersion=${GO_VERSION}'"

go build -ldflags "${LDFLAGS}" -o "${NAME}" "${REPO}/cmd/${NAME}"
