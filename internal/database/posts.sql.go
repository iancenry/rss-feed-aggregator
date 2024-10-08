// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: posts.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (id, feed_id, url, title, description, content, published_at, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, feed_id, url, title, description, content, published_at, created_at, updated_at
`

type CreatePostParams struct {
	ID          uuid.UUID
	FeedID      uuid.UUID
	Url         string
	Title       string
	Description sql.NullString
	Content     string
	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.FeedID,
		arg.Url,
		arg.Title,
		arg.Description,
		arg.Content,
		arg.PublishedAt,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.FeedID,
		&i.Url,
		&i.Title,
		&i.Description,
		&i.Content,
		&i.PublishedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPostsForUser = `-- name: GetPostsForUser :many
SELECT posts.id, posts.feed_id, posts.url, posts.title, posts.description, posts.content, posts.published_at, posts.created_at, posts.updated_at FROM posts
JOIN feed_follows ON feed_follows.feed_id = posts.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2
`

type GetPostsForUserParams struct {
	UserID uuid.UUID
	Limit  int32
}

func (q *Queries) GetPostsForUser(ctx context.Context, arg GetPostsForUserParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostsForUser, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.FeedID,
			&i.Url,
			&i.Title,
			&i.Description,
			&i.Content,
			&i.PublishedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
