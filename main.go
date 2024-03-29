package main

import (
	"os"

	"github.com/joho/godotenv"
)

var app App

var dbHost, dbUser, dbName, dbPass, dbType, port string

func init() {
	godotenv.Load()
	dbHost = os.Getenv("db_host")
	dbUser = os.Getenv("db_user")
	dbName = os.Getenv("db_name")
	dbPass = os.Getenv("db_pass")
	dbType = os.Getenv("db_type")
	port = os.Getenv("PORT")

	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbName == "" {
		dbName = "rest_db"
	}
	if port == "" {
		port = "8000"
	}
}

func main() {
	app = App{}
	app.Init(dbHost, dbUser, dbName, dbPass, dbType)
	app.Run(port)
}
