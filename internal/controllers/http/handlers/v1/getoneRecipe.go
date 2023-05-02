package v1

import (
	"TestBcraft/internal/adapters/db/postgresadapter"
	"TestBcraft/internal/controllers/dto"
	"TestBcraft/internal/domain/service"
	"TestBcraft/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func GetOneRecipe(c *gin.Context, p *pgxpool.Pool) {

	var getOne dto.GetOneRecipe

	err := c.ShouldBindJSON(&getOne)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	storageRecipe := postgresadapter.NewRecipeStorage(p)

	serviceRecipe := service.NewRecipeService(storageRecipe)

	resultRecipe, errGetOne := serviceRecipe.GetOne(getOne.ID)
	if errGetOne != nil {
		logger.ErrorLogger.Println(errGetOne)
	}

	c.JSON(http.StatusOK, resultRecipe)

}
