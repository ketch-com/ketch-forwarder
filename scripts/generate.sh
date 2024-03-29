#!/usr/bin/env sh

# Code generated by shipbuilder init 1.21.7. DO NOT EDIT.

if [ ! -f "./scripts/check.sh" ]; then
  cd $(command dirname -- "$(command readlink -f "$(command -v -- "$0")")")/..
fi

if [ -z "$*" ]; then
  what="gateway"
else
  what="$*"
fi

for i in $what; do
  if [ -f "./scripts/gen$i.sh" ]; then
    . "./scripts/gen$i.sh"
  fi
done
