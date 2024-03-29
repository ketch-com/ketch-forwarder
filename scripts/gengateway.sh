#!/usr/bin/env sh

# Code generated by shipbuilder init 1.21.7. DO NOT EDIT.

if [ ! -f "./scripts/check.sh" ]; then
  cd $(command dirname -- "$(command readlink -f "$(command -v -- "$0")")")/..
fi

. ./scripts/check.sh

if [ -d openapi ]; then
  check yq swagger_cli

  set -e

  echo "Merging all servers into openapi/servers/index.yaml..."
  $yq ea '. as $item ireduce ({}; . * $item )' openapi/servers/*.yaml openapi/servers/index.yaml > openapi/servers/index.yaml.bak
  mv openapi/servers/index.yaml.bak openapi/servers/index.yaml

  echo "Generating gateway indexes..."
  for i in parameters responses requestBodies schemas; do
    if [ -d "openapi/components/$i" ]; then
      echo "* openapi/components/$i/_index.yaml"
      find "openapi/components/$i" -type f -name '*.yaml' -not -name '_index.yaml' | xargs -I {} basename {} .yaml | xargs -I {} printf "{}:\n  \$ref: \"./{}.yaml\"\n" > "openapi/components/$i/_index.yaml"
    fi
  done

  servers=$(find openapi/servers -type f -name '*.yaml' -not -name 'index.yaml' | xargs -I {} basename {} .yaml)

  echo "Generating openapi definitions..."
  for i in $servers; do
    $swagger_cli bundle --outfile "openapi/${i}_gen.yaml" --type yaml "openapi/servers/$i.yaml"
  done

  $swagger_cli bundle --outfile openapi/index_gen.yaml --type yaml openapi/servers/index.yaml
fi
