package db

const (
	DriverMSSQL  = "mssql"
	DriverMySQL  = "mysql"
	DriverSQLite = "sqlite3"
)

type Config struct {
	Driver string
	User string
	Password string
	Host string
	Database string
	Port int
}

