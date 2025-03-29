package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v9"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB
var RedisClient *redis.Client

func ConnectDB() {
	dbURL := os.Getenv("DATABASE_URL")

	var err error
	DB, err = sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connected to PostgreSQL!")
}

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
		DB:   0,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	fmt.Println("Connected to Redis!")
}
