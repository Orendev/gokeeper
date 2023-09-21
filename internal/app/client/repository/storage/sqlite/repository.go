package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/pkg/store/sqlite"
	"github.com/pressly/goose"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	db      *sql.DB
	genSQL  squirrel.StatementBuilderType
	options sqlite.Options
}

func New(db *sql.DB, o sqlite.Options) (*Repository, error) {
	if err := migrations(db, o.MigrationsDir); err != nil {
		return nil, err
	}
	var r = &Repository{
		genSQL: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		db:     db,
	}

	r.SetOptions(o)
	return r, nil
}

func (r *Repository) SetOptions(options sqlite.Options) {
	if r.options != options {
		r.options = options
	}
}

func migrations(db *sql.DB, dir string) (err error) {
	err = goose.SetDialect("sqlite3")
	if err != nil {
		return err
	}
	goose.SetTableName("keeper_version")
	if err = goose.Run("up", db, dir); err != nil {
		return fmt.Errorf("goose %s error : %w", "up", err)
	}
	return
}
