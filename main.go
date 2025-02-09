package main

import (
	"context"
	"log"
	"simplebank/api"

	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/techschool/simplebank/db/sqlc"
)

const (
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	ctx := context.Background()
	conn, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the database:", err)
	}
	defer conn.Close()

	store := db.NewStore(conn)

	server := api.NewServer(&store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
