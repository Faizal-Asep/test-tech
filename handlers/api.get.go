package handler

import (
	"log"
	"net/http"

	model "github.com/Faizal-Asep/test-tech/models"
	repositorie "github.com/Faizal-Asep/test-tech/repositories"
	"github.com/gin-gonic/gin"
)

func ApiGet(c *gin.Context) {

	client := repositorie.NewHandPhone()
	p, err := client.Get(c.Request.Context())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.HttpResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, p)
}
