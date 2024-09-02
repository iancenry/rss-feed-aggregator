-- name: CreateFeed :one
INSERT INTO feeds (id, url, user_id, name, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;