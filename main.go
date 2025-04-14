package main

import (
	"database/sql"
	"github.com/rs/zerolog/log"
	"simplebanks/api"
	db "simplebanks/db/sqlc"
	"simplebanks/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading config: ")
	}
	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("Error opening database connection")
	}
	store := db.NewStore(conn)
	runGinServer(config, store)
}
func runGinServer(config util.Config, store db.Store) {
	//tạo máy chủ mới
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating server: ")
	}
	err = server.Start(config.HttpServerDriver)
	if err != nil {
		log.Fatal().Err(err).Msg("could not start GIN server")
	}
}
