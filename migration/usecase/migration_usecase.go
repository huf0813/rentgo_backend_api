package usecase

import (
	"context"
	"github.com/huf0813/rentgo_backend_api/domain"
	"time"
)

type MigrationUseCase struct {
	migrationRepoMysql domain.MigrationRepository
	timeOut            time.Duration
}

func NewMigrationUseCase(m domain.MigrationRepository, timeOut time.Duration) domain.MigrationUseCase {
	return &MigrationUseCase{
		migrationRepoMysql: m,
		timeOut:            timeOut,
	}
}

func (m *MigrationUseCase) Migrate(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, m.timeOut)
	defer cancel()

	if err := m.migrationRepoMysql.Migrate(ctx); err != nil {
		return err
	}

	return nil
}

func (m *MigrationUseCase) Drop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, m.timeOut)
	defer cancel()

	if err := m.migrationRepoMysql.Drop(ctx); err != nil {
		return err
	}

	return nil
}

func (m *MigrationUseCase) Seed(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, m.timeOut)
	defer cancel()

	if err := m.migrationRepoMysql.Seed(ctx); err != nil {
		return err
	}

	return nil
}

func (m *MigrationUseCase) Faker(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, m.timeOut)
	defer cancel()

	if err := m.migrationRepoMysql.Faker(ctx); err != nil {
		return err
	}

	return nil
}
