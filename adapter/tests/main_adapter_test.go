package adapter_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	onlinetest, _ = strconv.ParseBool(os.Getenv("ONLINE_TEST"))
	db            *sql.DB
	gdb           *gorm.DB
)

func setup() error {
	var err error
	if onlinetest {
		dsn := "host=testpghost user=test password=test123 dbname=woong_auth port=5432 sslmode=disable TimeZone=Asia/Seoul"
		if db, err = sql.Open("pgx", dsn); err != nil {
			log.Println(err)
			return err
		}
		db.SetMaxIdleConns(3)
		db.SetMaxOpenConns(3)
		db.SetConnMaxLifetime(3 * time.Minute)

		gdb, err = gorm.Open(postgres.New(postgres.Config{
			Conn: db,
		}), &gorm.Config{})
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func shutdown() {
	if db != nil {
		db.Close()
	}
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Println(err)
		shutdown()
		os.Exit(1)
	}

	exitCode := m.Run()

	shutdown()
	os.Exit(exitCode)
}
