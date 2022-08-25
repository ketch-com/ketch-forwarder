#!/usr/bin/env sh

# Code generated by shipbuilder init 1.18.0. DO NOT EDIT.

if [ ! -f "./scripts/check.sh" ]; then
  cd $(command dirname -- "$(command readlink -f "$(command -v -- "$0")")")/..
fi

. ./scripts/check.sh docker

./scripts/build.sh linux

set -e

$docker compose --profile service up --build --force-recreate --always-recreate-deps --quiet-pull --detach --wait $*
