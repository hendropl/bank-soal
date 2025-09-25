package constant

import "os"

const (
	UsersBucket = `users`
)

var JWT_SECRET = os.Getenv("JWT_SECRET")
var JWT_EXPIRED = "7d"

var SERVER_PORT = os.Getenv("SERVER_PORT")
var DB_USERNAME = os.Getenv("DB_USERNAME")
var DB_PASSWORD = os.Getenv("DB_PASSWORD")
var DB_NAME = os.Getenv("DB_NAME")
var DB_HOST = os.Getenv("DB_HOST")
var DB_PORT = os.Getenv("DB_PORT")
var DB_MAX_IDLE_CONNS = os.Getenv("DB_MAX_IDLE_CONNS")
var DB_MAX_OPEN_CONNS = os.Getenv("DB_MAX_OPEN_CONNS")
var DB_CONN_MAX_LIFETIME = os.Getenv("DB_CONN_MAX_LIFETIME")
