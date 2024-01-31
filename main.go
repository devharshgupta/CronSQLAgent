package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron"
)

const (
	dbHost     = "your_db_host"
	dbPort     = "your_db_port"
	dbUser     = "your_db_user"
	dbPassword = "your_db_password"
	dbName     = "your_db_name"
)

func main() {
	// Create a cron job scheduler
	c := cron.New()

	// Add a cron job that runs the updateQuery function every minute
	c.AddFunc("@every 1m", updateQuery)

	// Start the cron job scheduler
	c.Start()

	// Run indefinitely, or until interrupted
	select {}
}

func updateQuery() {
	// Establish a connection to the PostgreSQL database
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName))
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Prepare the update query
	query := "UPDATE your_table SET your_column = 'new_value' WHERE your_condition"

	// Execute the update query
	_, err = db.Exec(query)
	if err != nil {
		log.Println("Error executing update query:", err)
		return
	}

	log.Println("Update query executed successfully at", time.Now())
}
