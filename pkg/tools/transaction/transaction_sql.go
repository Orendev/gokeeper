package transaction

import (
	"context"
	"database/sql"
)

func FinishSQL(ctx context.Context, tx *sql.Tx, err error) error {
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	} else {
		if commitErr := tx.Commit(); commitErr != nil {
			return commitErr
		}
		return nil
	}
}
