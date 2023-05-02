package v2

import (
	"TestBcraft/internal/adapters/db/postgresadapter"
	"TestBcraft/internal/domain/service"
	"TestBcraft/internal/token"
	"TestBcraft/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func CurrentUser(c *gin.Context, p *pgxpool.Pool) {

	userId, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.ErrorLogger.Println(err)
		return
	}

	storageUser := postgresadapter.NewUserStorage(p)
	serviceUser := service.NewUserService(storageUser)

	user, errGetUser := serviceUser.GetUserById(int(userId))

	if errGetUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errGetUser.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})

}
