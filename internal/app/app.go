package app

import (
	"TestBcraft/internal/controllers/http/handlers"
	"TestBcraft/pkg/logger"
	"TestBcraft/pkg/postgresbd"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func RunApp() {

	logger.Init()
	logger.InfoLogger.Println("Starting the application...")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	pgclient, err := postgresbd.NewClientPostgres()
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	r := gin.Default()

	r.POST("/create_recipe", func(c *gin.Context) {
		handlers.CreateRecipe(c, pgclient)
	})

	r.POST("/delete_recipe", func(c *gin.Context) {
		handlers.DeleteRecipe(c, pgclient)
	})

	r.POST("/update_recipe", func(c *gin.Context) {
	})
	r.Run()

}
