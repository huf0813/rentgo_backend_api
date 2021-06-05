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
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.User{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.ProductCategory{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.Product{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.ProductImage{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.EventCategory{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.Event{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.Invoice{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.InvoiceProduct{}); err != nil {
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
