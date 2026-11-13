# gator

A minimal CLI RSS blog aggregator in Go using PostgreSQL and `sqlc`, built for the [Boot.dev](https://www.boot.dev) guided project.

This is a learning project only—neither finished nor polished, and its chaotic structure reflects learning along the way.

## Prerequisites

- [Go](https://go.dev/) (`1.27.3` in `go.mod`, uses `encoding/json/v2`)
- [PostgreSQL](https://www.postgresql.org/) (schema uses `uuidv7()`, available in Postgres 18+)
- [goose](https://github.com/pressly/goose) (for running database migrations)

## Install

```bash
go install github.com/ppieniazek/gator@latest
```

## Setup

1. Create a PostgreSQL database (for example, `gator`).

2. Run the schema migrations in `sql/schema`:

```bash
goose -dir sql/schema postgres "postgres://<user>:<password>@localhost:5432/gator?sslmode=disable" up
```

3. Create `~/.gatorconfig.json` (this file must exist before running `gator`):

```json
{
  "db_url": "postgres://<user>:<password>@localhost:5432/gator?sslmode=disable"
}
```

## Run

Run commands with `gator <command> [args...]`:

- `gator register <name>` — register a user and switch to them
- `gator login <name>` — switch to an existing user
- `gator addfeed <name> <url>` — add and follow an RSS feed
- `gator agg <duration>` — continuously fetch feeds on an interval (e.g. `1m`, minimum `5s`)
- `gator browse [limit]` — print recent posts from followed feeds (defaults to `2`)
- `gator feeds` / `gator follow <url>` / `gator following` / `gator unfollow <url>` — list and manage feeds
