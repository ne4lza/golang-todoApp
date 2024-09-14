package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type PostgresDB struct {
	DB *sqlx.DB
}

func GetDb() *PostgresDB {
	db := createDbConn()
	return &PostgresDB{
		DB: db,
	}
}
func createDbConn() *sqlx.DB {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatalf(".ENV DOSYASI YÜKLENİRKEN BİR HATA OLUŞTU: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		log.Fatalf("VT HATASI : %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("VT PING HATASI : %v", err)
	}

	fmt.Println("DB BAĞLANTISI BAŞARILI!")
	
	return db
}
