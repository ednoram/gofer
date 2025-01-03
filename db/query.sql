-- name: GetTask :one
SELECT * FROM task
WHERE task_id = ? LIMIT 1;

-- name: ListTasks :many
SELECT * FROM task;

-- name: CreateTask :one
INSERT INTO task (
  title, description, completed, created_by
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateTask :exec
UPDATE task
set title = ?,
description = ?,
completed = ?,
updated_at = ?
WHERE task_id = ?
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM task
WHERE task_id = ?;


-- name: GetApiKey :one
SELECT * FROM api_key
WHERE api_key = ? LIMIT 1;
