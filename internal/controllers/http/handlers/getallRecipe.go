package handlers

import (
	"TestBcraft/internal/adapters/db/postgresadapter"
	"TestBcraft/internal/controllers/dto"
	"TestBcraft/internal/domain/service"
	"TestBcraft/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func GetAllRecipe(c *gin.Context, p *pgxpool.Pool) {

	var dtoData dto.GetAllRecipe

	err := c.ShouldBindJSON(dtoData)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	storageRecipe := postgresadapter.NewRecipeStorage(p)

	serviceRecipe := service.NewRecipeService(storageRecipe)

	result, errGetAll := serviceRecipe.GetAll()
	if errGetAll != nil {
		logger.ErrorLogger.Println(errGetAll)
	}

	c.JSON(http.StatusOK, result)

}
