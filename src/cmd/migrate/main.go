package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/AgufSamudra/subscription/src/internal/databases"
)

const envFile = ".env"

func main() {
	if err := loadEnvFile(envFile); err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	databaseURL := os.Getenv("DATABASE_URL")

	switch command {
	case "create":
		if len(os.Args) < 3 {
			log.Fatal("migration name is required")
		}
		if err := createMigrationFiles(os.Args[2]); err != nil {
			log.Fatal(err)
		}
	case "up":
		requireDatabaseURL(databaseURL)
		if err := databases.MigrateUp(databaseURL); err != nil {
			log.Fatal(err)
		}
		log.Println("migration up completed")
	case "down":
		requireDatabaseURL(databaseURL)
		if err := databases.MigrateDown(databaseURL); err != nil {
			log.Fatal(err)
		}
		log.Println("migration down completed")
	case "steps":
		requireDatabaseURL(databaseURL)
		if len(os.Args) < 3 {
			log.Fatal("steps value is required")
		}
		steps, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		if err := databases.MigrateSteps(databaseURL, steps); err != nil {
			log.Fatal(err)
		}
		log.Println("migration steps completed")
	case "version":
		requireDatabaseURL(databaseURL)
		m, err := databases.NewMigrator(databaseURL)
		if err != nil {
			log.Fatal(err)
		}
		version, dirty, err := m.Version()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("migration version: %d, dirty: %t\n", version, dirty)
	default:
		printUsage()
		os.Exit(1)
	}
}

func requireDatabaseURL(databaseURL string) {
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
}

func createMigrationFiles(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("migration name cannot be empty")
	}

	safeName := strings.ToLower(name)
	safeName = strings.ReplaceAll(safeName, " ", "_")
	version := time.Now().Format("20060102150405")
	baseName := fmt.Sprintf("%s_%s", version, safeName)

	migrationPath := databases.DefaultMigrationPath
	if err := os.MkdirAll(migrationPath, 0755); err != nil {
		return err
	}

	upFile := filepath.Join(migrationPath, baseName+".up.sql")
	downFile := filepath.Join(migrationPath, baseName+".down.sql")

	if err := os.WriteFile(upFile, []byte("-- Write your migrate up SQL here.\n"), 0644); err != nil {
		return err
	}
	if err := os.WriteFile(downFile, []byte("-- Write your migrate down SQL here.\n"), 0644); err != nil {
		return err
	}

	log.Printf("created %s\n", upFile)
	log.Printf("created %s\n", downFile)
	return nil
}

func loadEnvFile(path string) error {
	file, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, found := strings.Cut(line, "=")
		if !found {
			continue
		}

		key = strings.TrimSpace(key)
		value = strings.Trim(strings.TrimSpace(value), `"'`)
		if key != "" {
			_ = os.Setenv(key, value)
		}
	}

	return scanner.Err()
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  go run ./src/cmd/migrate create <migration_name>")
	fmt.Println("  go run ./src/cmd/migrate up")
	fmt.Println("  go run ./src/cmd/migrate down")
	fmt.Println("  go run ./src/cmd/migrate steps <number>")
	fmt.Println("  go run ./src/cmd/migrate version")
}
