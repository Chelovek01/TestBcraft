package handlers

import (
	"TestBcraft/internal/adapters/db/postgresadapter"
	"TestBcraft/internal/controllers/dto"
	"TestBcraft/internal/domain/service"
	"TestBcraft/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func updateRecipe(c *gin.Context, p *pgxpool.Pool) {

	var updateObject dto.UpdateRecipeDTO

	err := c.ShouldBindJSON(&updateObject)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	recipeStorage := postgresadapter.NewRecipeStorage(p)

	recipeService := service.RecipeStorage(recipeStorage)

	err = recipeService.Update()

}
