package postgres

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/internal/app/server/configs"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type Repository struct {
	db      *pgxpool.Pool
	genSQL  squirrel.StatementBuilderType
	options configs.DB
}

func New(db *pgxpool.Pool, o configs.DB) (*Repository, error) {
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

func (r *Repository) SetOptions(options configs.DB) {
	if options.DefaultLimit == 0 {
		options.DefaultLimit = 10
	}
	if r.options != options {
		r.options = options
	}
}

func migrations(pool *pgxpool.Pool, dir string) (err error) {
	db, err := goose.OpenDBWithDriver("pgx", pool.Config().ConnConfig.ConnString())
	if err != nil {
		return err
	}
	defer func() {
		if errClose := db.Close(); errClose != nil {
			err = errClose
			return
		}
	}()

	goose.SetTableName("keeper_version")
	if err = goose.Run("up", db, dir); err != nil {
		return fmt.Errorf("goose %s error : %w", "up", err)
	}
	return
}
