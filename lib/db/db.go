package db

import (
	"fmt"

	"github.com/JustKato/example-golang-api/lib/conf"

	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Get the database cofiguration from the environment
func getDatbaseConfig() conf.DatabaseConfig {
	// Get the main configuration of the app without forcefully refreshing the cache
	conf := conf.GetConfig(false)

	// Return the database configuration
	return conf.Database
}

func getDatabaseConnectionString() string {
	// Get the database's configuration
	dbc := getDatbaseConfig()

	return fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", dbc.User, dbc.Pass, dbc.Host, dbc.Port, dbc.Db)
}

// get a database connection
func GetDatabaseConnection() *sql.DB {

	// Using the database configuration connect to the database
	db, err := sql.Open("mysql", getDatabaseConnectionString())
	if err != nil {
		panic(err)
	}

	dbc := getDatbaseConfig()

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(dbc.MaxConn)
	db.SetMaxIdleConns(dbc.MaxIdleConn)

	return db
}
