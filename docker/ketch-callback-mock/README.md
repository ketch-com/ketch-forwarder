# ketch-forwarder

This folder produces the ketch-forwarder container for the service.

## Building

Build this container using the following command from the repository root:

```shell
./scripts/build.sh linux
docker build -f docker/ketch-forwarder/Dockerfile .
```

## Contents

This folder contains a `Dockerfile` and a `docker-entrypoint.sh`. The `docker-entrypoint.sh` is copied into the docker
container and set as the `ENTRYPOINT`.
