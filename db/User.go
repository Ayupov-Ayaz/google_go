package db

type User struct {
	ID int `db: id`
	FirstName string `db:first_name`
	LastName string `db:last_name`
	Age int `db:age`
	Email string `db:email`
}