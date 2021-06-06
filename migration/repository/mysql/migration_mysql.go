package mysql

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/huf0813/rentgo_backend_api/domain"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type MigrationRepoMysql struct {
	DB *gorm.DB
}

func NewMigrationRepoMysql(db *gorm.DB) domain.MigrationRepository {
	return &MigrationRepoMysql{DB: db}
}

func (m *MigrationRepoMysql) Migrate(ctx context.Context) error {
	/* start drop table */
	// layer four
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		DropTable(&domain.InvoiceProduct{}); err != nil {
		return err
	}

	// layer three
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		DropTable(&domain.ProductImage{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		DropTable(&domain.EventProduct{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		DropTable(&domain.Invoice{}); err != nil {
		return err
	}

	// layer two
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		DropTable(&domain.Product{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		DropTable(&domain.Event{}); err != nil {
		return err
	}
	// layer one
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		DropTable(&domain.User{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		DropTable(&domain.ProductCategory{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		DropTable(&domain.EventCategory{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		DropTable(&domain.InvoiceCategory{}); err != nil {
		return err
	}
	/* end drop table */

	/* start migrate table */
	// layer one
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
		CreateTable(&domain.EventCategory{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.InvoiceCategory{}); err != nil {
		return err
	}

	// layer two
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
		CreateTable(&domain.Event{}); err != nil {
		return err
	}

	// layer three
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
		CreateTable(&domain.EventProduct{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.Invoice{}); err != nil {
		return err
	}

	// layer four
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.InvoiceProduct{}); err != nil {
		return err
	}
	/* end migrate table */
	return nil
}

func (m *MigrationRepoMysql) Seed(ctx context.Context) error {
	if err := m.DB.
		WithContext(ctx).
		Exec("INSERT INTO product_categories(name,created_at,updated_at) VALUES "+
			"('furniture',?,?), "+
			"('household',?,?), "+
			"('automotive',?,?);",
			time.Now(), time.Now(),
			time.Now(), time.Now(),
			time.Now(), time.Now()).Error; err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Exec("INSERT INTO event_categories(name,created_at,updated_at) VALUES "+
			"('indoor',?,?), "+
			"('birthday',?,?), "+
			"('outdoor',?,?);",
			time.Now(), time.Now(),
			time.Now(), time.Now(),
			time.Now(), time.Now()).Error; err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Exec("INSERT INTO invoice_categories(name,created_at,updated_at) VALUES "+
			"('cart',?,?), "+
			"('on_going',?,?), "+
			"('completed',?,?), "+
			"('to_pay',?,?);",
			time.Now(), time.Now(),
			time.Now(), time.Now(),
			time.Now(), time.Now(),
			time.Now(), time.Now()).Error; err != nil {
		return err
	}
	return nil
}

func (m *MigrationRepoMysql) Faker(ctx context.Context) error {
	for i := 0; i < 100; i++ {
		newProduct := domain.Product{
			Name:              faker.FirstName(),
			Price:             uint(rand.Intn(500000-10000) + 10000),
			Stock:             uint(rand.Intn(50-10) + 10),
			Star:              uint(rand.Intn(5-1) + 1),
			ProductCategoryID: uint(rand.Intn(3-1) + 1),
		}
		if err := m.DB.
			WithContext(ctx).
			Create(&newProduct).Error; err != nil {
			return err
		}
	}
	return nil
}
