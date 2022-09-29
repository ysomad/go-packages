package migrate

import (
	"errors"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	defaultAttempts = 20
	defaultTimeout  = time.Second
)

func connect(migrationsDir, connString string) *migrate.Migrate {
	if migrationsDir == "" {
		log.Fatalf("migrate: migrationsDir is empty")
	}
	if connString == "" {
		log.Fatalf("migrate: connString is empty")
	}

	connString += "?sslmode=disable"

	var (
		attempts = defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://"+migrationsDir, connString)
		if err == nil {
			break
		}

		log.Printf("migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Fatalf("migrate: postgres connect error: %s", err)
	}

	return m
}

func Up(migrationsDir, connString string) {
	m := connect(migrationsDir, connString)

	err := m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("migrate: up error: %s", err)
	}
	defer m.Close()

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("migrate: no change")
		return
	}

	log.Printf("migrate: up success")
}

func Down(migrationsDir, connString string) {
	m := connect(migrationsDir, connString)

	err := m.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("migrate: down error: %s", err)
	}
	defer m.Close()

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("migrate: no change")
		return
	}

	log.Printf("migrate: down success")
}
