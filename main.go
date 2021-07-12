package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	connStatus := connect()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("starting server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if connStatus {
			_, _ = fmt.Fprintln(w, `Hello, visitor! DB Connection successful`)
		} else {
			_, _ = fmt.Fprintln(w, `Hello, visitor! DB Connection Failed`)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func connect() bool {

	var (
		host     = goDotEnvVariable("HOST")
		port     = 5432
		user     = goDotEnvVariable("DB_USER")
		password = goDotEnvVariable("DB_PASSWORD")
		dbname   = goDotEnvVariable("DB_NAME")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		//panic(err)
		log.Fatal("Error Connecting to DB")
		return false
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		//panic(err)
		log.Fatal("Cannot Access DB")
		return false
	}

	fmt.Println("Successfully Connected!")
	return true

}
