# rss-feed-aggregator

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

- Replace `<username>` and `<password>`.
- goose postgres postgres://<username>:<password>@localhost:5432/rssagg up
- To check for users sp as to get username - query:

```sql
SELECT *
FROM pg_catalog.pg_user;
```

- After creating queries and schemas - Run `sqlc generate` so that sqlc can generate the go code for the sql files under the sql/queries and sql/schema folders.
