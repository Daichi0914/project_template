package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
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
	requiredEnvs := []string{"MYSQL_HOST", "MYSQL_USER", "MYSQL_PASSWORD", "MYSQL_DATABASE"}
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
		DBHost:     os.Getenv("MYSQL_HOST"),
		DBPort:     getEnv("MYSQL_PORT", "3306"),
		DBUser:     os.Getenv("MYSQL_USER"),
		DBPassword: os.Getenv("MYSQL_PASSWORD"),
		DBName:     os.Getenv("MYSQL_DATABASE"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}

	return config, nil
}

// GetDSN はデータベース接続用のDSNを返します
func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}

// LoadEnvFile は指定された環境変数ファイルをロードします
func LoadEnvFile(fileName string) {
	// まず、現在のディレクトリを試す
	err := godotenv.Load(fileName)
	if err == nil {
		fmt.Printf("Loaded environment from %s\n", fileName)
		fmt.Printf("MYSQL_HOST=%s, MYSQL_PORT=%s, MYSQL_USER=%s, MYSQL_DATABASE=%s\n", 
			os.Getenv("MYSQL_HOST"), 
			os.Getenv("MYSQL_PORT"), 
			os.Getenv("MYSQL_USER"), 
			os.Getenv("MYSQL_DATABASE"))
		return
	}
	
	// 親ディレクトリを試す（最大3階層）
	for i := 0; i < 3; i++ {
		envFile := filepath.Join("..", fileName)
		err = godotenv.Load(envFile)
		if err == nil {
			fmt.Printf("Loaded environment from %s\n", envFile)
			fmt.Printf("MYSQL_HOST=%s, MYSQL_PORT=%s, MYSQL_USER=%s, MYSQL_DATABASE=%s\n", 
				os.Getenv("MYSQL_HOST"), 
				os.Getenv("MYSQL_PORT"), 
				os.Getenv("MYSQL_USER"), 
				os.Getenv("MYSQL_DATABASE"))
			return
		}
	}
	
	fmt.Printf("Warning: %s file not found. Using OS environment variables.\n", fileName)
}

// getEnv は環境変数を取得し、未設定の場合はデフォルト値を返します
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
