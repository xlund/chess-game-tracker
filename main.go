package main

import (
	"embed"

	"github.com/joho/godotenv"
	"github.com/xlund/chess-games-tracker/api"
	"github.com/xlund/chess-games-tracker/repository"
)

//go:embed web/*
var fs embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	auth, err := repository.New()
	if err != nil {
		panic(err)
	}
	app := api.New(&auth, fs)

	app.Logger.Fatal(app.Start("localhost:4000"))

}
