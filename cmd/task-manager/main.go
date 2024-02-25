package main

import (
	"log"

	db "github.com/abhinavthapa1998/task-manager/internal/DB"
	"github.com/abhinavthapa1998/task-manager/internal/config"
	"github.com/abhinavthapa1998/task-manager/internal/routes"
	"github.com/abhinavthapa1998/task-manager/internal/store"
	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlStorage := db.NewMySQLStorage(cfg)

	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := store.NewStore(db)

	server := routes.NewAPIServer(":8080", store)
	server.Serve()
}
