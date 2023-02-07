package main

import (
	"database/sql"
	"log"

	"github.com/fxfrancky/simplebank/api"
	db "github.com/fxfrancky/simplebank/db/sqlc"
	"github.com/fxfrancky/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".") // "." the path is the current folder
	if err != nil {
		log.Fatal("Cannot Log configurations")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot Create Server :", err)
	}
	err = server.Start(config.ServerAdress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
