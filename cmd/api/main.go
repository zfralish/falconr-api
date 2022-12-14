package main

import (
	"context"
	"database/sql"
	"falconr-api/internal/data"
	firebase "firebase.google.com/go"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
	"time"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
	fb firebase.App
}

type application struct {
	config config
	logger *log.Logger
	models data.Models
}

const version = "1.0.0"

func main() {

	opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_APP")))
	fbApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	var cfg config
	cfg.fb = *fbApp

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	if cfg.env == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	cfg.db.dsn = os.Getenv("DATABASE_URL")

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}(db)

	logger.Printf("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Printf("Starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()

	logger.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
