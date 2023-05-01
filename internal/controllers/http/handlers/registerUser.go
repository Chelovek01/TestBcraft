package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func Register(c *gin.Context, p *pgxpool.Pool) {

	c.JSON(http.StatusOK, gin.H{"data": "hello from handler"})

}
