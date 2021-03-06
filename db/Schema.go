package db

import (
	"fmt"
)

type Schema struct {
	tableName string
	create string
	drop string
}

/**
 	Создание всех таблиц
 */
func CreateTables(db *DB) error {
	for _, schema := range allSchemas {
		_, err := db.Query(schema.create)
		if err != nil {
			fmt.Println("Не удалось создать таблицу:", schema.tableName, "\n", err.Error())
		}
	}
	return nil
}

/**
	Удаление всех таблиц
 */
func DropTables(db *DB) error {
	for _, schema := range allSchemas {
		_, err := db.Query(schema.drop)
		if err != nil {
			fmt.Println("Не удалось удалить таблицу:", schema.tableName, "\n", err.Error())
		}
	}
	return nil
}

var allSchemas = []Schema{
	// table users
	{
		tableName: "users",
		create: `CREATE TABLE IF NOT EXISTS users (
  				 id SERIAL NOT NULL PRIMARY KEY,
  				 first_name VARCHAR(150) NOT NULL,
  				 last_name VARCHAR(150) NOT NULL,
  				 email VARCHAR(255) NOT NULL,
  				 age int(3) NOT NULL);`,

		drop:   `DROP TABLE users`,
	},
	//
}