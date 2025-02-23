DB_FILE := "db/gofer.db"

sync:
  go mod tidy

gen:
  sqlc generate

init: sync gen
  # Copy env example without overriding existing env
  cp -n .env.example .env; exit 0
  # Create tables
  sqlite3 {{DB_FILE}} < db/schema.sql
  # Create a sample user
  sqlite3 {{DB_FILE}} "INSERT INTO user (username) VALUES ('user')"
  # Generate and print API key for user
  just gen-api-key 1


[group('dev')]
serve:
  go run cmd/server/main.go

[group('dev')]
run-client command:
  go run cmd/client/main.go {{command}}

[group('dev')]
gen-api-key user_id:
  go run cmd/gen_api_key/main.go {{user_id}}


[group('build')]
build-server: sync gen
  go build -o bin/server cmd/server/main.go

[group('build')]
build-client: sync gen
  go build -o bin/client cmd/client/main.go

[group('build')]
build: build-server build-client
