# Running

## Running in Docker Compose

To run the dependencies, use the following:

```shell
docker compose up
```

When testing the service, ensure you use the following to ensure the latest binary is executed:
```shell
./scripts/run.sh
```

Additional profiles are available:

| Profile   | Description                                             |
|-----------|---------------------------------------------------------|
| service   | Run the actual service                                  |
| smoketest | Run the smoketest after bringing up the service/gateway |
