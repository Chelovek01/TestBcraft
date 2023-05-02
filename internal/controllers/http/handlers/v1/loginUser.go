package v1

import (
	"TestBcraft/internal/adapters/db/postgresadapter"
	"TestBcraft/internal/controllers/dto"
	"TestBcraft/internal/domain/entity"
	"TestBcraft/internal/domain/service"
	"TestBcraft/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func LoginUser(c *gin.Context, p *pgxpool.Pool) error {

	var inputData dto.InputLogin

	var userLogin entity.User

	loginStorage := postgresadapter.NewUserStorage(p)

	loginService := service.NewUserService(loginStorage)

	err := c.ShouldBindJSON(&inputData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.ErrorLogger.Println()
		return err
	}

	userLogin.Username = inputData.Username
	userLogin.Password = inputData.Password

	token, errToken := loginService.LoginCheck(&userLogin)

	if errToken != nil {
		logger.ErrorLogger.Println(err)
		return errToken
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
	return err

}
