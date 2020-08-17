package middleware

import (
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"log"
	"net/http" // used to access the request and response object of the api
	"strconv"  // package used to covert string into int type

	"fmt"
	"github.com/gorilla/mux" // used to get the params from the route
)

type response struct {
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contexxt-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	fmt.Print(id)

	/*	user, err := getUser(int64(id))

		if err != nil {
			log.Fatalf("Unable to get user. %v", err)
		}
	*/

	// send the response
	json.NewEncoder(w).Encode("")
}
