package handlers

import (
	"TestBcraft/internal/adapters/db/postgresadapter"
	"TestBcraft/internal/controllers/dto"
	"TestBcraft/internal/domain/service"
	"TestBcraft/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func UpdateRecipe(c *gin.Context, p *pgxpool.Pool) {

	var updateObject dto.UpdateRecipeDTO

	err := c.ShouldBindJSON(&updateObject)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	recipeStorage := postgresadapter.NewRecipeStorage(p)

	recipeService := service.NewRecipeService(recipeStorage)

	err = recipeService.Update(updateObject.ID, updateObject.Column, updateObject.Value)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

}
