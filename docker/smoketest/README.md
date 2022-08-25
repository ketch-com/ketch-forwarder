# smoketest

This folder produces the smoketest container used for smoketesting a service.

## Building

Build this container using the following command from the repository root:

```shell
./scripts/build.sh
docker build -f docker/smoketest/Dockerfile .
```
