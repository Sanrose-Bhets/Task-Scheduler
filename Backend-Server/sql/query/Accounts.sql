-- name: CreateUser :one
INSERT INTO accounts( Name, email, Role, APIkey)
VALUES($1, $2, $3, 
encode(sha256(random()::text::bytea),'hex'))
RETURNING *;

-- name: GetUserbyAPI :one
SELECT * from accounts where APIKey = $1;

-- name: GetAllUsers :many
SELECT * from accounts;