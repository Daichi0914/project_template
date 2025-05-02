package config

import (
	"fmt"
	"os"
)

// Config はアプリケーション設定を管理します
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

// NewConfig は環境変数から設定を読み込みます
func NewConfig() (*Config, error) {
	// 必須環境変数の確認
	requiredEnvs := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	var missingEnvs []string

	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			missingEnvs = append(missingEnvs, env)
		}
	}

	if len(missingEnvs) > 0 {
		return nil, fmt.Errorf("required environment variables are not set: %v", missingEnvs)
	}

	// 環境変数から設定を取得
	config := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}

	return config, nil
}

// GetDSN はデータベース接続用のDSNを返します
func (c *Config) GetDSN() string {
	return c.DBUser + ":" + c.DBPassword + "@tcp(" + c.DBHost + ":" + c.DBPort + ")/" + c.DBName + "?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci"
}

// getEnv は環境変数を取得し、未設定の場合はデフォルト値を返します
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
