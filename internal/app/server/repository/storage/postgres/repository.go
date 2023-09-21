package postgres

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/pkg/store/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pressly/goose"
)

type Repository struct {
	db      *pgxpool.Pool
	genSQL  squirrel.StatementBuilderType
	options postgres.Options
}

func New(db *pgxpool.Pool, o postgres.Options) (*Repository, error) {
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

func (r *Repository) SetOptions(options postgres.Options) {
	if options.DefaultLimit == 0 {
		options.DefaultLimit = 10
	}
	if r.options != options {
		r.options = options
	}
}

func migrations(pool *pgxpool.Pool, dir string) (err error) {
	db, err := goose.OpenDBWithDriver("postgres", pool.Config().ConnConfig.ConnString())
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
