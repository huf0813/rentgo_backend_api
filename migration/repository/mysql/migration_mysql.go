package mysql

import (
	"context"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/huf0813/rentgo_backend_api/domain"
	"github.com/huf0813/rentgo_backend_api/utils/custom_security"
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
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.Cart{}); err != nil {
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
	return nil
}

func (m *MigrationRepoMysql) Drop(ctx context.Context) error {
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
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		DropTable(&domain.Cart{}); err != nil {
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
	/* user */
	for i := 1; i <= 10; i++ {
		password, err := custom_security.NewHashingValue("1234567890")
		if err != nil {
			return err
		}
		newUser := domain.User{
			Name:     faker.FirstName(),
			Email:    fmt.Sprintf("user%d@gmail.com", i),
			Password: password,
		}
		if err := m.DB.
			WithContext(ctx).
			Create(&newUser).Error; err != nil {
			return err
		}
	}
	/* user */

	/* product */
	for i := 1; i <= 100; i++ {
		// product
		newProduct := domain.Product{
			Name:              faker.FirstName(),
			Price:             uint(rand.Intn(500000-10000) + 10000),
			Stock:             uint(rand.Intn(50-10) + 10),
			ProductCategoryID: uint(rand.Intn(3-1) + 1),
			UserID:            uint(rand.Intn(10-1) + 1),
		}
		if err := m.DB.
			WithContext(ctx).
			Create(&newProduct).Error; err != nil {
			return err
		}

		// product image
		newProductImage := domain.ProductImage{
			ProductID: uint(i),
			Path:      "default.jpg",
		}
		if err := m.DB.
			WithContext(ctx).
			Create(&newProductImage).Error; err != nil {
			return err
		}
	}
	/* product */
	return nil
}
