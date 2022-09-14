#!/usr/bin/env sh

# Code generated by shipbuilder init 1.21.2. DO NOT EDIT.

if [ ! -f "./scripts/check.sh" ]; then
  cd $(command dirname -- "$(command readlink -f "$(command -v -- "$0")")")/..
fi

. ./scripts/check.sh

check go

set -e

if [ -z "$*" ]; then
  args="./..."
else
  args="$*"
fi

$go test -v --tags unit $args
