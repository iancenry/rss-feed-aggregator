# rss-feed-aggregator

- Keeps track of rss feeds and periodically downloads them.
- go mod init github.com/iancenry/rss-feed-aggregator

## Sample package installation flow

- go get github.com/joho/godotenv
- go mod vendor
- go mod tidy

## Install sqlc and goose into your command line

- This projects uses sqlc to handle queries and goose to handle migrations:

  - go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest / brew install sqlc
  - execute `sqlc version`
  - go install github.com/pressly/goose/v3/cmd/goose@latest / brew install goose
  - `goose -version`

## Migrations

- Run in the schema folder.
- Replace `<username>` and `<password>`.
- goose postgres postgres://<username>:<password>@localhost:5432/rssagg up
- To check for users as to get username - query:

```sql
SELECT *
FROM pg_catalog.pg_user;
```

- After creating queries and schemas - Run `sqlc generate` in base so that sqlc can generate the go code for the sql files under the sql/queries and sql/schema folders.

## Marshaling

- The JSON tags are not part of the Go language syntax but are a feature provided by the encoding/json package. This package handles JSON encoding and decoding and respects these tags when marshaling (converting Go values to JSON) and unmarshaling (converting JSON to Go values) data

```go
type User struct{
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
}
```

## env vars

```
PORT=5000
DB_URL=postgres://<username>:<password>@localhost:5432/rssagg?sslmode=disable
```
