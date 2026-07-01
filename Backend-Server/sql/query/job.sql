-- name: CreateJobs :one
INSERT INTO jobs( positionName, companyId, experienceRequired, expectedCompensation, employmentType, jobLocationType)
VALUES($1, $2, $3, $4, $5, $6)
RETURNING *;