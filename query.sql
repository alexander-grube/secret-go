-- name: CreateSecret :one
INSERT INTO secret_message (message, user_id) VALUES ($1, $2) RETURNING *;

-- name: GetSecret :one
SELECT * FROM secret_message WHERE id = $1;

-- name: DeleteSecret :exec
DELETE FROM secret_message WHERE id = $1;

-- name: CreateUser :one
INSERT INTO public.user (username) VALUES ($1) RETURNING *;

-- name: GetUser :one
SELECT * FROM public.user WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM public.user WHERE id = $1;

-- name: GetMessagesOfUser :many
SELECT id, message FROM secret_message WHERE user_id = $1;