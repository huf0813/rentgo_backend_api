package mysql

import (
	"context"
	"github.com/huf0813/rentgo_backend_api/domain"
	"gorm.io/gorm"
)

type MigrationRepoMysql struct {
	DB *gorm.DB
}

func NewMigrationRepoMysql(db *gorm.DB) domain.MigrationRepository {
	return &MigrationRepoMysql{DB: db}
}

func (m *MigrationRepoMysql) Migrate(ctx context.Context) error {
	if err := m.DB.
		WithContext(ctx).Error; err != nil {
		return err
	}
	return nil
}

func (m *MigrationRepoMysql) Seed(ctx context.Context) error {
	if err := m.DB.
		WithContext(ctx).Error; err != nil {
		return err
	}
	return nil
}
