package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// DB is the database connection
var DB *sql.DB

// InitDB initializes the database connection
func InitDB() error {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Get database connection parameters from environment variables
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "schools")
	sslmode := getEnv("DB_SSLMODE", "disable")

	// Create connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	// Open database connection
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Connected to database successfully")
	return nil
}

// CreateTables creates the necessary tables in the database
func CreateTables() error {
	// Create schools table
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS schools (
		id SERIAL PRIMARY KEY,
		objectid INTEGER UNIQUE,
		name TEXT NOT NULL,
		address TEXT,
		city TEXT,
		state TEXT,
		zip TEXT,
		country TEXT,
		county TEXT,
		countyfips TEXT,
		latitude FLOAT,
		longitude FLOAT,
		level TEXT,
		st_grade TEXT,
		end_grade TEXT,
		enrollment INTEGER,
		ft_teacher INTEGER,
		type INTEGER,
		status INTEGER,
		population INTEGER,
		ncesid TEXT,
		districtid TEXT,
		naics_code TEXT,
		naics_desc TEXT,
		website TEXT,
		telephone TEXT,
		sourcedate TIMESTAMP,
		val_date TIMESTAMP,
		val_method TEXT,
		source TEXT,
		shelter_id TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	
	-- Add PostGIS extension if not exists
	CREATE EXTENSION IF NOT EXISTS postgis;
	
	-- Add geometry column for location
	DO $$
	BEGIN
		IF NOT EXISTS (
			SELECT 1 FROM information_schema.columns 
			WHERE table_name = 'schools' AND column_name = 'location'
		) THEN
			PERFORM AddGeometryColumn('schools', 'location', 4326, 'POINT', 2);
		END IF;
	END $$;
	
	-- Create index on geometry column
	CREATE INDEX IF NOT EXISTS schools_location_idx ON schools USING GIST(location);
	
	-- Create index on objectid
	CREATE INDEX IF NOT EXISTS schools_objectid_idx ON schools(objectid);
	`)
	
	if err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	log.Println("Tables created successfully")
	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}

// Helper function to get environment variable with fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}