-- name: CreatePost :one
INSERT INTO posts (id, feed_id, url, title, description, content, published_at, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetPostsForUser :many
SELECT * FROM posts
WHERE feed_id = $1
ORDER BY published_at DESC
LIMIT $2 OFFSET $3