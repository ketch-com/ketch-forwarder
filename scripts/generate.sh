#!/usr/bin/env sh

# Code generated by shipbuilder init 1.18.0. DO NOT EDIT.

if [ ! -f "./scripts/check.sh" ]; then
  cd $(command dirname -- "$(command readlink -f "$(command -v -- "$0")")")/..
fi

if [ -f "./features/shipbuilder/.env" ]; then
  . ./features/shipbuilder/.env
fi

if [ -z "$*" ]; then
  what="$shipbuilder_generate"
else
  what="$*"
fi

for i in $what; do
  if [ -f "./scripts/gen$i.sh" ]; then
    . "./scripts/gen$i.sh"
  fi
done
