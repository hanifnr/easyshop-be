package utils

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"

	_ "github.com/jackc/pgx/v4/stdlib"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func OpenDBConnection() *gorm.DB {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Warning: %s environment variable not set.\n", k)
		}
		return v
	}
	var (
		dbUser         = mustGetenv("DB_USER")              // e.g. 'my-db-user'
		dbPwd          = mustGetenv("DB_PASS")              // e.g. 'my-db-password'
		dbName         = mustGetenv("DB_NAME")              // e.g. 'my-database'
		unixSocketPath = mustGetenv("INSTANCE_UNIX_SOCKET") // e.g. '/cloudsql/project:region:instance'
	)
	fmt.Printf("pg info %s %s %s %s", dbUser, dbPwd, dbName, unixSocketPath)
	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s",
		dbUser, dbPwd, dbName, unixSocketPath)

	dbPool, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		fmt.Printf("sql.Open: %v", err)
	}

	// [START_EXCLUDE]
	db, _ := dbPool.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(10 * time.Minute)
	// [END_EXCLUDE]

	return dbPool
}

func GetDB() *gorm.DB {
	once.Do(func() {
		db = OpenDBConnection()
	})
	return db
}
