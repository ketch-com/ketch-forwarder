# Testing

## HTTP/GRPC requests

The `test/requests` folder and any subfolders required for organization store `.http` files that can be used to test
the services in this repository.

## Test Fixtures

All test fixtures are stored in the `test/fixtures` folder and any subfolders required for organization. These files are
made available using the following import and code:

```go
import "github.com/ketch-com/ketch-forwarder/test"

func TestSomething(t *testing.T) {
	// This opens the test/fixtures/foo/bar.json file from assets
    f, err := test.Fixtures.Open("foo/bar.json")
}

```

## Unit testing

All unit tests are in files sitting in the same package folder as the units under test.

To unit test this repository, run the dependencies as described in [RUNNING](RUNNING.md).

All unit tests should be created with the following build tag:

```go
//go:build unit && !integration && !smoke

```

Then you can run unit tests using Go Test:

```shell
go test -v --tags unit ./...
```

You can also set the `unit` build tag in your IDE.

## Integration testing

All integration tests are in folders in the `test/integration` folder.

To integration test this repository, run the service as described in [RUNNING](RUNNING.md).

All integration tests should be created with the following build tag:

```go
//go:build !unit && integration && !smoke

```

Then you can run integration tests using Go Test:

```shell
go test -v --tags integration ./...
```

## Smoke testing

The smoke tests should be created in the `test/smoketest` folder. Environment variables should be loaded using the `SMOKETEST` prefix.

All smoke tests should be created with the following build tag:

```go
//go:build !unit && !integration && smoke

```

To smoketest this repository, as follows:

```shell
./scripts/build.sh
docker compose --profile smoketest up --build --attach smoketest
```
