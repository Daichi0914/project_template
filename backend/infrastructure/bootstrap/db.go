package bootstrap

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"project_template/backend/infrastructure/config"
)

const (
	maxRetries    = 5
	retryInterval = 3 * time.Second
)

// InitDB はデータベース接続を初期化します
func InitDB(cfg *config.Config) *sql.DB {
	var db *sql.DB
	var err error

	// 接続の再試行
	for i := 0; i < maxRetries; i++ {
		log.Printf("Attempting to connect to database (attempt %d/%d)...", i+1, maxRetries)
		
		db, err = sql.Open("mysql", cfg.GetDSN())
		if err != nil {
			log.Printf("Failed to open database connection: %v", err)
			time.Sleep(retryInterval)
			continue
		}

		// 接続確認
		if err := db.Ping(); err != nil {
			log.Printf("Failed to ping database: %v", err)
			time.Sleep(retryInterval)
			continue
		}

		// 接続成功
		log.Println("Connected to database successfully")
		break
	}

	// 全ての再試行が失敗した場合
	if err != nil {
		log.Fatalf("Failed to connect to database after %d attempts: %v", maxRetries, err)
	}

	// 接続設定
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}
