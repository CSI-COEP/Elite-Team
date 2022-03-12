package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"sync"
)

type PostgresDatabase struct {
	mutex sync.RWMutex
	Db    *sqlx.DB
}

func GetDBType() string {
	dbType := os.Getenv("DBType")
	return dbType
}

func GetPostgresConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DBHost"),
		os.Getenv("DBPort"),
		os.Getenv("DBUser"),
		os.Getenv("DBName"),
		os.Getenv("DBPassword"),
	)
	return dataBase
}

func InitPostgresDataBase() *PostgresDatabase {
	db, err := sqlx.Connect(GetDBType(), GetPostgresConnectionString())
	if err != nil {
		panic(err)
	}

	return &PostgresDatabase{
		Db: db,
	}
}
