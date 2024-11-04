package main

import (
	"context"
	"log"

	"github.com/EleyOliveira/leilao_go/configuration/database/mongodb"
	"github.com/joho/godotenv"
)

func main() {

	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
		return
	}

	databaseClient, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	databaseClient.Client().Disconnect(ctx)

}
