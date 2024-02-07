package db

import (
	"context"
	"database/sql"
	"embed"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/o5h/services/config"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrationsDir embed.FS

var (
	ctx context.Context
	db  *sql.DB
)

func Init(c context.Context) {
	ctx = c
	cfg := c.Value(config.ContextKey).(*config.Config).DB
	var err error
	db, err = sql.Open("pgx", cfg.URL)
	must(err)

	db.SetMaxOpenConns(cfg.MaxOpen)
	db.SetMaxIdleConns(cfg.MaxIdle)

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

func mustRow[T any](t *T, err error) *T {
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		} else {
			must(err)
		}
	}
	return t
}

func Close() { db.Close() }

func BeginTx() *sql.Tx {
	log.Println("Begin Tx")
	tx, err := db.BeginTx(ctx, nil)
	must(err)
	return tx
}

func BeginReadOnlyTx() *sql.Tx {
	log.Println("Begin Tx")
	tx, err := db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	must(err)
	return tx
}

func Commit(tx *sql.Tx) {
	if r := recover(); r != nil {
		log.Println("Rollback Tx", r)
		tx.Rollback()
	} else {
		if err := tx.Commit(); err != nil {
			log.Println("Rollback Tx")
			tx.Rollback()
		} else {
			log.Println("Commited Tx")
		}
	}
}
