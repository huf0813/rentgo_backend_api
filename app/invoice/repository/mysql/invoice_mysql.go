package mysql

import (
	"github.com/huf0813/rentgo_backend_api/domain"
	"gorm.io/gorm"
)

type InvoiceRepoMysql struct {
	DB *gorm.DB
}

func NewInvoiceRepoMysql(db *gorm.DB) domain.InvoiceRepository {
	return &InvoiceRepoMysql{DB: db}
}
