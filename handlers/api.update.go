package handler

import (
	"log"
	"net/http"

	model "github.com/Faizal-Asep/test-tech/models"
	repositorie "github.com/Faizal-Asep/test-tech/repositories"
	"github.com/gin-gonic/gin"
)

func ApiUpdate(c *gin.Context) {
	var in repositorie.HandPhone

	id := c.Param("id")

	if err := c.ShouldBindJSON(&in); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.HttpResponse{Message: err.Error()})
		return
	}

	client := repositorie.NewHandPhone()
	p, err := client.Update(c.Request.Context(), id, in)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.HttpResponse{Message: err.Error()})
		return
	}

	go func() {
		model.Notif <- model.Notification{
			Command: "update",
			Data:    p,
		}
	}()

	c.JSON(http.StatusOK, p)

}
