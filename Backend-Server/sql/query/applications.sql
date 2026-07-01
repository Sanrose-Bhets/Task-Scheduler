-- name: CreateApplications :one
INSERT INTO applications( department, status, company, position, interviewId)
VALUES($1, $2, $3,$4, $5)
RETURNING *;