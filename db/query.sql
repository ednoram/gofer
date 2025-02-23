-- name: GetTask :one
SELECT * FROM task
WHERE task_id = ? LIMIT 1;

-- name: ListTasks :many
SELECT * FROM task
WHERE completed = COALESCE(?, completed);

-- name: CreateTask :one
INSERT INTO task (
  title, description, completed, created_by
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateTask :exec
UPDATE task
set updated_at = sqlc.arg('updated_at'),
title = COALESCE(sqlc.narg('title'),  title),
description = COALESCE(sqlc.narg('description'), description),
completed = COALESCE(sqlc.narg('completed'),  completed)

WHERE task_id = sqlc.arg('task_id');

-- name: DeleteTask :exec
DELETE FROM task
WHERE task_id = ?;


-- name: GetApiKey :one
SELECT * FROM api_key
WHERE api_key = ? LIMIT 1;
