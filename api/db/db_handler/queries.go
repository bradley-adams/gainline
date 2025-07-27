package db_handler

import (
	"context"

	"github.com/bradley-adams/gainline/db/db"
)

func CreateCompetition(
	ctx context.Context,
	q Queries,
	params db.CreateCompetitionParams,
) error {
	err := q.CreateCompetition(ctx, params)
	if err != nil {
		return err
	}
	return nil
}
