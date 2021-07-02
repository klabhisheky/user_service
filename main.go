package main

import (
	"log"
	"os"

	"developer.zopsmart.com/go/backend/zs"
	userHTTPServer "github.com/klabhisheky/user_service/http/user"
	userService "github.com/klabhisheky/user_service/services/user"
	userStore "github.com/klabhisheky/user_service/store/user"
)

func main() {
	//database connection, maybe move it to init()
	logger := log.New(os.Stdout, "[LOG:]", 0)
	db := zs.NewMYSQL(logger, zs.MySQLConfig{
		HostName: os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		Database: "client",
	})
	defer db.Close()

	//userHTTP server
	if db != nil {
		userHTTPServer := userHTTPServer.New(userService.New(userStore.New(db)))
		router := zs.NewRouter()
		router.REST("user", userHTTPServer)
		zs.StartServers(logger, &zs.HTTPServer{Router: router, Port: 9100})
	}
}
