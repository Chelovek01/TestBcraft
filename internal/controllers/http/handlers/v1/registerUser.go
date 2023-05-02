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

func Register(c *gin.Context, p *pgxpool.Pool) {

	var inputData dto.RegisterInput
	var user entity.User
	userStorage := postgresadapter.NewUserStorage(p)
	userService := service.NewUserService(userStorage)

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.ErrorLogger.Println(err)
		return
	}

	user.Username = inputData.Username
	user.Password = inputData.Password

	errHashPass := userService.HashPass(&user)
	if err != nil {
		logger.ErrorLogger.Println(errHashPass)
	}

	errCreateUser := userService.CreateUser(&user)
	if errCreateUser != nil {
		logger.ErrorLogger.Println(errCreateUser)
	}

	c.JSON(http.StatusOK, gin.H{"data": "registered"})

}
