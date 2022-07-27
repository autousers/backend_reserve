package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/autousers/backend_reserve/models"
	"github.com/joho/godotenv"
)

func createConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")

	// check error load environment
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	// Open connection to database
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// Check connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database successfully")

	//return the connection
	return db
}

// GetAllRoom will return all the rooms in the database
func GetAllRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get all the rooms in database
	rooms, err := getAllRooms()

	if err != nil {
		log.Fatalf("Failed to get all the rooms in database: %v", err)
	}

	// send all the rooms as response
	json.NewEncoder(w).Encode(rooms)
}

//-------------------------handlers functions--------------------------------
// get all rooms from the database
func getAllRooms() ([]models.Rooms, error) {
	// create the postgres connection
	db := createConnection()

	// close the postgres connection
	defer db.Close()

	var rooms []models.Rooms

	// create the select sql query
	sqlStatement := "SELECT * FROM rooms"

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Failed to get all the query rows in database: %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var room models.Rooms

		err = rows.Scan(&room.ID, &room.Room_name, &room.Room_desc)

		if err != nil {
			log.Fatalf("Failed to scan the rows in database: %v", err)
		}

		rooms = append(rooms, room)
	}

	return rooms, err
}
