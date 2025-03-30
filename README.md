# ezeX Notification Service

A microservice for sending notifications through various mediums such as email and SMS.
Currently, it supports email notifications only.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- **[Go](https://go.dev/doc/install/)**: The Go programming language.
- **[Protocol Buffer Compiler](https://protobuf.dev/installation/)**: Needed to generate code from `.proto` files.
- **Development Tools**: Run `make devtools` to install necessary tools for development.

### Build

To build the project, use:

```bash
make build
```

## Test

To run the tests, use:

```bash
make test
```

## Code Quality and Formatting

To automatically format the code, run:

```bash
make fmt
```

Run the linter to catch common mistakes and improve code quality:

```bash
make check
```

## Contributing

Contributions are most welcome!
Whether it's code, documentation, or ideas, every contribution makes a difference.
Please read the [Contributing](CONTRIBUTING.md) guide to get started.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
