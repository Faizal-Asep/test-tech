package handler

import (
	"log"
	"net/http"

	model "github.com/Faizal-Asep/test-tech/models"
	repositorie "github.com/Faizal-Asep/test-tech/repositories"
	"github.com/gin-gonic/gin"
)

func ApiGetId(c *gin.Context) {
	id := c.Param("id")

	client := repositorie.NewHandPhone()
	p, err := client.GetById(c.Request.Context(), id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.HttpResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)

}
