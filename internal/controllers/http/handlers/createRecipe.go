package handlers

import (
	"TestBcraft/internal/adapters/db/postgresadapter"
	"TestBcraft/internal/controllers/dto"
	"TestBcraft/internal/domain/entity"
	"TestBcraft/internal/domain/service"
	"TestBcraft/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func CreateRecipe(c *gin.Context, p *pgxpool.Pool) {

	var insertData dto.CreateRecipeDTO

	var recipe entity.Recipe

	err := c.ShouldBindJSON(&insertData)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	recipe.ID = insertData.Recipe.ID
	recipe.Name = insertData.Recipe.Name
	recipe.Description = insertData.Recipe.Description
	recipe.Ingredients = insertData.Recipe.Ingredients
	recipe.PrepareSteps = insertData.Recipe.PrepareSteps

	recipeStorage := postgresadapter.NewRecipeStorage(p)

	recipeService := service.NewRecipeService(recipeStorage)

	err = recipeService.Create(&recipe)

	if err != nil {
		logger.ErrorLogger.Println(err)
		fmt.Println(err)

	}
	c.JSON(http.StatusOK, gin.H{"message": "created"})
}
