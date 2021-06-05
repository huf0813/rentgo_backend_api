package mysql_driver

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type envMysqlDriver struct {
	mysqlHost     string
	mysqlPort     string
	mysqlUsername string
	mysqlPassword string
	mysqlDatabase string
}

func readEnvMysql() (envMysqlDriver, error) {
	if err := godotenv.Load(".env"); err != nil {
		return envMysqlDriver{}, err
	}
	mysqlHost := os.Getenv("MYSQL_DB_HOST")
	mysqlPort := os.Getenv("MYSQL_DB_PORT")
	mysqlUsername := os.Getenv("MYSQL_DB_USERNAME")
	mysqlPassword := os.Getenv("MYSQL_DB_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_DB_DATABASE")
	return envMysqlDriver{
		mysqlHost:     mysqlHost,
		mysqlPort:     mysqlPort,
		mysqlUsername: mysqlUsername,
		mysqlPassword: mysqlPassword,
		mysqlDatabase: mysqlDatabase,
	}, nil
}

func NewMysqlDriver() (*gorm.DB, error) {
	env, err := readEnvMysql()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.mysqlUsername,
		env.mysqlPassword,
		env.mysqlHost,
		env.mysqlPort,
		env.mysqlDatabase,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
