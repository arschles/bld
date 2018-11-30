#!/bin/sh -l

set -e
set -o pipefail

cd /github/workspace
go test ./...
