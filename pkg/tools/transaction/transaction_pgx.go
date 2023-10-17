package transaction

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func FinishPGX(ctx context.Context, tx pgx.Tx, err error) error {
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return rollbackErr
		}
		return err
	} else {
		if commitErr := tx.Commit(ctx); commitErr != nil {
			return commitErr
		}
		return nil
	}
}
