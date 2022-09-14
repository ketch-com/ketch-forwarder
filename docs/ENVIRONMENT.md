# Environment

The environment variables available for this repository are documented here, by the container name.

## ketch-forwarder

| Variable                                   | Description                                                                                       |
|--------------------------------------------|---------------------------------------------------------------------------------------------------|
| `KETCH_FORWARDER_SERVER_*`    | [HTTP server variables](https://github.com/ketch-com/orlop-http/blob/main/docs/ENVIRONMENT.md)    |
| `KETCH_FORWARDER_VAULT_*`     | [Vault variables](https://github.com/ketch-com/orlop-vault/blob/main/docs/ENVIRONMENT.md)         |
| `KETCH_FORWARDER_EVENTS_*`    | [NATS variables](https://github.com/ketch-com/orlop-nats/blob/main/docs/ENVIRONMENT.md)           |
| `KETCH_FORWARDER_{CLIENT}_*`  | [GRPC client variables](https://github.com/ketch-com/orlop-grpc/blob/main/docs/ENVIRONMENT.md)    |
| `KETCH_FORWARDER_{CLIENT}_*`  | [HTTP client variables](https://github.com/ketch-com/orlop-http/blob/main/docs/ENVIRONMENT.md)    |

For more environment variables, look at the `README.md` files inside the `features` folder.
