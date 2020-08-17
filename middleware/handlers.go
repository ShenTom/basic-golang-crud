package middleware

import (
	"database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"log"
	"net/http" // used to access the request and response object of the api
	"os"
	"strconv" // package used to covert string into int type

	"github.com/gorilla/mux" // used to get the params from the route
	"github.com/joho/godotenv"

	"basic-golang-crud/models"

	_ "github.com/lib/pq" // postgres golang driver
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	// return the connection
	return db
}

// GetUser fetches a user row in the db with #id
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contexxt-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	user, err := getUser(id)

	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(user)
}

// CreateUser inserts a user object into the db
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	insertID, err := insertUser(user)

	// format a response object
	res := response{
		ID:      insertID,
		Message: "User created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// UpdateUser updates a user row in the db
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRowID, err := updateUser(user)

	res := response{
		ID:      updatedRowID,
		Message: "User updated successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// --------------------------- Handler functions -------------------------

func getUser(id int) (models.User, error) {
	db := createConnection()
	defer db.Close()

	var user models.User

	sqlStatement := `SELECT * FROM users WHERE userid=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return user, err
}

func insertUser(user models.User) (int64, error) {
	db := createConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO users (name, location, age) VALUES ($1, $2, $3)`

	var id int64
	err := db.QueryRow(sqlStatement, user.Name, user.Location, user.Age).Scan(&id)

	return id, err
}

func updateUser(user models.User) (int64, error) {
	db := createConnection()
	defer db.Close()

	sqlStatement := `UPDATE users SET name=$1, location=$2, age=$3 WHERE userid = $4`

	res, err := db.Exec(sqlStatement, user.Name, user.Location, user.Age, user.ID)
	rowsAffected, err := res.RowsAffected()
	return rowsAffected, err
}
