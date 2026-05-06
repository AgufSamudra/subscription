package databases

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const (
	defaultMaxOpenConns    = 25
	defaultMaxIdleConns    = 25
	defaultConnMaxLifetime = 5 * time.Minute
	defaultConnMaxIdleTime = 5 * time.Minute
	defaultConnectTimeout  = 5 * time.Second
)

func PostgreSQLConnection(ctx context.Context) (*sql.DB, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL wajib diisi")
	}

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("membuka koneksi postgresql: %w", err)
	}

	db.SetMaxOpenConns(envInt("DB_MAX_OPEN_CONNS", defaultMaxOpenConns))
	db.SetMaxIdleConns(envInt("DB_MAX_IDLE_CONNS", defaultMaxIdleConns))
	db.SetConnMaxLifetime(envDuration("DB_CONN_MAX_LIFETIME", defaultConnMaxLifetime))
	db.SetConnMaxIdleTime(envDuration("DB_CONN_MAX_IDLE_TIME", defaultConnMaxIdleTime))

	pingCtx, cancel := context.WithTimeout(ctx, envDuration("DB_CONNECT_TIMEOUT", defaultConnectTimeout))
	defer cancel()

	if err := db.PingContext(pingCtx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping postgresql gagal: %w", err)
	}

	return db, nil
}

func envInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil || parsed <= 0 {
		return fallback
	}

	return parsed
}

func envDuration(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	duration, err := time.ParseDuration(value)
	if err == nil && duration > 0 {
		return duration
	}

	seconds, err := strconv.Atoi(value)
	if err != nil || seconds <= 0 {
		return fallback
	}

	return time.Duration(seconds) * time.Second
}
