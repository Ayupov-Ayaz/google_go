package web

import "google_go/db"

func DbStart() {
	//
	config := db.Config{
		Driver: "mysql",
		User: "tommy",
		Password: "43",
		Host: "127.0.0.1",
		Database: "go_lang",
		Port: 3306,
	}

	connection, err := db.DbConnection(&config)
	if err != nil {
		panic(err)
	}

	db.CreateTables(connection)
}