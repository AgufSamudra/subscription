package databases

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const DefaultMigrationPath = "src/internal/databases/migrations"

func NewMigrator(databaseURL string) (*migrate.Migrate, error) {
	return migrate.New("file://"+DefaultMigrationPath, databaseURL)
}

func MigrateUp(databaseURL string) error {
	m, err := NewMigrator(databaseURL)
	if err != nil {
		return err
	}

	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}

	return err
}

func MigrateDown(databaseURL string) error {
	m, err := NewMigrator(databaseURL)
	if err != nil {
		return err
	}

	err = m.Down()
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}

	return err
}

func MigrateSteps(databaseURL string, steps int) error {
	m, err := NewMigrator(databaseURL)
	if err != nil {
		return err
	}

	err = m.Steps(steps)
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}

	return err
}
