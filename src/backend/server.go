package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
)

type Penyakit struct {
	NamaPenyakit string
	DNA          string
}

func getEnv(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", getEnv("DATABASE_USERNAME")+":"+getEnv("DATABASE_PASSWORD")+"@tcp("+getEnv("DATABASE_PORT")+")/"+getEnv("DATABASE_NAME"))

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM jenispenyakit")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var penyakit Penyakit
		// for each row, scan the result into our tag composite object
		err = results.Scan(&penyakit.NamaPenyakit, &penyakit.DNA)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Println(penyakit.DNA)
	}

	// Server
	http.HandleFunc(getEnv("BASE_PORT")+"/", helloWorld)

	fmt.Println("Starting server at port " + getEnv("BACKEND_PORT"))
	if err := http.ListenAndServe(":"+getEnv("BACKEND_PORT"), nil); err != nil {
		log.Fatal(err)
	}
}