DB_FILE := "gofer.db"

sync:
  go mod tidy

gen:
  sqlc generate

init: sync gen
  # Create tables
  sqlite3 {{DB_FILE}} < db/schema.sql
  # Create a sample user
  sqlite3 {{DB_FILE}} "INSERT INTO user (username) VALUES ('user')"
  # Generate and print API key for user
  just gen-api-key 1


[group('code')]
serve:
  go run cmd/server/main.go

[group('code')]
run-client:
  go run cmd/client/main.go

[group('code')]
gen-api-key user_id:
  go run cmd/gen_api_key/main.go {{user_id}}
