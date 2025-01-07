# Gofer

Gofer is a simple tool for task management.

## Usage

- Create a `gofer.db` sqlite database file in the root directory of the project.
- Initialize database tables using `db/schema.sql` file.
- Run `sqlc generate` to generate database helpers.

## Authentication

API key authentication is used for calling the API.
You can generate an API key using `cmd/gen_api_key/main.go` script. This will print the API key and store the hash in the database.
Set `GOFER_API_KEY` environment variable before running the client.

