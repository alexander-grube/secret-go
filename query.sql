-- name: CreateSecret :one
INSERT INTO secret_message (message) VALUES ($1) RETURNING *;

-- name: GetSecret :one
SELECT * FROM secret_message WHERE id = $1;

-- name: DeleteSecret :exec
DELETE FROM secret_message WHERE id = $1;