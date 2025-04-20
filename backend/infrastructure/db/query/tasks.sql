-- name: CreateTask :one
INSERT INTO tasks (name, estimate_minutes, started_at)
VALUES ($1, $2, $3)
    RETURNING *;

-- name: StopTask :one
UPDATE tasks
SET stopped_at = $2
WHERE id = $1
    RETURNING *;

-- name: ListTasks :many
SELECT * FROM tasks ORDER BY started_at DESC;
