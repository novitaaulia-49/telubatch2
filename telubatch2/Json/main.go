package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

type Employees struct {
	Address struct {
		City       string `json:"city"`
		PostalCode string `json:"postalCode"`
		Region     string `json:"region"`
		Street     string `json:"street"`
	} `json:"address"`
	Birthdate string `json:"birthdate"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// type Customer struct {
// 	Address struct {
// 		City   string `json:"city"`
// 		State  string `json:"state"`
// 		Street string `json:"street"`
// 		Zip    string `json:"zip"`
// 	} `json:"address"`
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// }

func getEmployees(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	var request Employees

	if err = json.Unmarshal(body, &request); err != nil {
		fmt.Println("Failed decoding json message")
	} else {
		if r.Method == "POST" {

			FirstName := request.FirstName
			LastName := request.LastName

			stmt, err := db.Prepare("insert into employees (LastName,FirstName) VALUES (?,?)")

			_, err = stmt.Exec(LastName, FirstName)

			if err != nil {
				fmt.Fprintf(w, "Data Duplicate")
			} else {
				fmt.Fprintf(w, "Data Created")
			}

		}
	}

}

func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/northwind")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	fmt.Println("Server on :8181")

	// Route handles & endpoints
	r.HandleFunc("/employees", getEmployees).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8181", r))

}
