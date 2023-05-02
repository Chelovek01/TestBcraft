package app

import (
	"TestBcraft/internal/controllers/http/handlers/v1"
	"TestBcraft/internal/controllers/http/handlers/v2"
	"TestBcraft/internal/middlewares"
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

	public := r.Group("/v1")

	public.GET("get_one_recipe", func(c *gin.Context) {
		v1.GetOneRecipe(c, pgclient)
	})

	public.GET("/get_all_recipe", func(c *gin.Context) {
		v1.GetAllRecipe(c, pgclient)
	})

	public.POST("/register", func(c *gin.Context) {
		v1.Register(c, pgclient)
	})

	public.POST("/login", func(c *gin.Context) {
		err := v1.LoginUser(c, pgclient)
		if err != nil {
			logger.ErrorLogger.Println(err)
		}
	})

	protected := r.Group("v2")

	protected.Use(middlewares.JwtAuthMiddleware())

	protected.POST("/create_recipe", func(c *gin.Context) {
		v2.CreateRecipe(c, pgclient)
	})

	protected.POST("/delete_recipe", func(c *gin.Context) {
		v2.DeleteRecipe(c, pgclient)
	})

	protected.POST("/update_recipe", func(c *gin.Context) {
		v2.UpdateRecipe(c, pgclient)
	})

	protected.GET("/user", func(c *gin.Context) {
		v2.CurrentUser(c, pgclient)
	})

	r.Run()

}
