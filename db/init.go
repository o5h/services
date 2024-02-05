package db

import (
	"context"
	"database/sql"
	"embed"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/o5h/services/config"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrationsDir embed.FS

var db *sql.DB

func Init(ctx context.Context) {
	cfg := ctx.Value(config.ContextKey).(*config.Config)
	var err error
	db, err = sql.Open("pgx", cfg.DB.URL)
	must(err)
	updateSchema()
}

func updateSchema() {
	goose.SetBaseFS(migrationsDir)
	must(goose.SetDialect("postgres"))
	must(goose.Up(db, "migrations"))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
