-- name: CreateInterviews :one
INSERT INTO interviews( status, interviewDate)
VALUES ($1, $2)
RETURNING *;