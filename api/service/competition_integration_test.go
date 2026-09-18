//go:build integration
// +build integration

package service

import (
	"context"
	"testing"

	"github.com/bradley-adams/gainline/http/api"
	"github.com/stretchr/testify/require"
)

func TestCompetitionService_Integration(t *testing.T) {
	dbHandler := setupTestDB(t)
	svc := NewCompetitionService(dbHandler)
	ctx := context.Background()

	t.Run("create, get, update and soft-delete a competition", func(t *testing.T) {
		created, err := svc.Create(ctx, &api.CompetitionRequest{Name: "Integration Test League"})
		require.NoError(t, err)
		require.Equal(t, "Integration Test League", created.Name)

		fetched, err := svc.Get(ctx, created.ID)
		require.NoError(t, err)
		require.Equal(t, created.ID, fetched.ID)

		updated, err := svc.Update(ctx, created.ID, &api.CompetitionRequest{Name: "Renamed League"})
		require.NoError(t, err)
		require.Equal(t, "Renamed League", updated.Name)

		require.NoError(t, svc.Delete(ctx, created.ID))

		_, err = svc.Get(ctx, created.ID)
		require.Error(t, err, "GetCompetition should exclude soft-deleted rows")
	})

	t.Run("Create returns ErrCompetitionNameTaken on a case-insensitive duplicate name", func(t *testing.T) {
		_, err := svc.Create(ctx, &api.CompetitionRequest{Name: "Duplicate League"})
		require.NoError(t, err)

		_, err = svc.Create(ctx, &api.CompetitionRequest{Name: "duplicate league"})
		require.ErrorIs(t, err, ErrCompetitionNameTaken)
	})

	t.Run("Update returns ErrCompetitionNameTaken when renaming onto an existing name", func(t *testing.T) {
		_, err := svc.Create(ctx, &api.CompetitionRequest{Name: "Existing League"})
		require.NoError(t, err)

		toRename, err := svc.Create(ctx, &api.CompetitionRequest{Name: "League To Rename"})
		require.NoError(t, err)

		_, err = svc.Update(ctx, toRename.ID, &api.CompetitionRequest{Name: "existing league"})
		require.ErrorIs(t, err, ErrCompetitionNameTaken)
	})

	t.Run("a soft-deleted name can be reused", func(t *testing.T) {
		first, err := svc.Create(ctx, &api.CompetitionRequest{Name: "Reused Name League"})
		require.NoError(t, err)
		require.NoError(t, svc.Delete(ctx, first.ID))

		_, err = svc.Create(ctx, &api.CompetitionRequest{Name: "Reused Name League"})
		require.NoError(t, err, "the unique index is partial on deleted_at IS NULL, so this should succeed")
	})
}
