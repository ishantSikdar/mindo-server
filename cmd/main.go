package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ishantSikdar/mindo-server/internal/handlers"
	"github.com/ishantSikdar/mindo-server/pkg/db"
	"github.com/ishantSikdar/mindo-server/pkg/logger"
)

func main() {

	err := db.InitDB("postgresql://postgres:6515@localhost:5432/mindo?sslmode=disable")
	if err != nil {
		logger.Log.Error("Failed to init DB: ", err)
	}

	router := gin.Default()
	handlers.RegisterRoutes(router)

	routerErr := router.Run(":8080")

	if routerErr != nil {
		panic(routerErr)
	}
}
