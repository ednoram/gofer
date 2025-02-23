# Gofer

Gofer is a simple tool for managing your errands and tasks.

One of the hosts in the network functions as the central HTTP server. All other hosts interact with the service using the Gofer CLI client.

## Setup

Install the required tools:

- go (v1.23.6) - <https://go.dev/doc/install>
- sqlc (v1.28.0) - <https://docs.sqlc.dev/en/stable/overview/install.html>
- just - <https://github.com/casey/just>

Run `just init` to initialize the environment.

See `justfile` for other actions.

## Database

After completing the initialization step, sqlite database file will be created at `db/gofer.db`.
User management is handled manually through the database.

## Authentication

API key authentication is used for calling the API.
You can generate an API key by running `just gen-api-key <user_id>`.
This will print the API key and store the hash in the database.
Set `GOFER_API_KEY` environment variable before running the client.

## Configuration

Application configuration is at `config/config.toml`. You can modify it if needed.

## Build

Run `just build-server` to build the server and `just build-client` to build the client.
Build output will be saved at `./bin/`.
Run the binaries from project root for using correct relative paths to configuration and database files.
