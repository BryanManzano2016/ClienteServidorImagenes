package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

const (
	dbUsuario  = "postgres"
	dbPassword = "sqlsql"
	dbNombre   = "Usuarios"
)

func main() {
	http.HandleFunc("/", imagenesPasar)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func imagenesPasar(w http.ResponseWriter, r *http.Request) {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUsuario, dbPassword, dbNombre)
	db, err := sql.Open("postgres", dbinfo)

	check(err)

	rows, err := db.Query("SELECT nombre, datos FROM imagenes")
	db.Close()

	for rows.Next() {
		var datos []byte
		var nombre string
		err = rows.Scan(&nombre, &datos)
		check(err)

		w.Write(datos)
		fmt.Println(time.Now())
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
