# MCP Kubernetes Client

This repository contains a Go project for interacting with Kubernetes resources via the MCP (Model Context Protocol) server.

## Overview

The project provides command-line utilities that use a Kubernetes client implementation to list and fetch logs from different resources. It is structured around two main packages:

- `resources/`: helpers for the Kubernetes client, resource listing, and log retrieval
- `structvalues/`: package holding shared data structures

The `main.go` file builds an executable which ties everything together.

## Building

Ensure you have Go installed (>=1.18 recommended) and run:

```bash
go build -o <binary-name>
```

This will produce an executable in the workspace root (replace `<binary-name>` with your preferred name).

## Usage

Typical commands depend on the implementation details, but the executable is designed to communicate with a Kubernetes cluster via the MCP server. Modify source code as needed to configure context or resource types.

## Project Structure

```
<binary-name>        # built binary output (rename after building)
main.go             # entrypoint
resources/          # kubernetes_client.go, list_resources.go, logs.go
structvalues/       # structs.go
```

## Contributing

Pull requests are welcome. Follow standard Go project conventions, keep formatting with `go fmt`, and write tests when adding features.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
