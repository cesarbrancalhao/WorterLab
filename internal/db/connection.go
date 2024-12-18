package db

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	RD *redis.Client
)

func ConnectPG() {

	dsn := os.Getenv("PSQL")
	if dsn == "" {
		log.Fatalln("Failed to retrieve connection data. {1}")
	}

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}
}

func ConnectRedis() {
	db, err := strconv.Atoi(os.Getenv("REDISDB"))
	addr := os.Getenv("REDISURL")
	pw := os.Getenv("REDISPW")

	if err != nil || addr == "" {
		log.Fatalln("Failed to retrieve connection data. {2}")
	}

	RD = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pw,
		DB:       db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := RD.Ping(ctx).Err(); err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}
}

func ConnectDB() {
	ConnectPG()
	ConnectRedis()
}
