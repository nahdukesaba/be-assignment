package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/nahdukesaba/be-assignment/docs"
	"github.com/nahdukesaba/be-assignment/repo"
	"github.com/nahdukesaba/be-assignment/server"
)

func main() {
	db := repo.NewDB()
	defer db.Close()

	err := repo.InitSchema(db)
	if err != nil {
		log.Fatalf("Error initializing schema: %v\n", err)
	}

	r := gin.Default()
	srv := server.NewServer(r, db)
	srv.RegisterRoutes()

	srv.Run()
}
