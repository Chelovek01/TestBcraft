package handlers

import (
	"TestBcraft/internal/adapters/db/postgresadapter"
	"TestBcraft/internal/domain/entity"
	"TestBcraft/internal/domain/service"
	"TestBcraft/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateRecipe(c *gin.Context, p *pgxpool.Pool) {

	var recipe entity.Recipe

	err := c.ShouldBindJSON(&recipe)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	recipeStorage := postgresadapter.NewRecipeStorage(p)

	recipeService := service.NewRecipeService(recipeStorage)

	err = recipeService.Create(&recipe)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
}
