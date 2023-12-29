package handler

import (
	"log"
	"net/http"

	model "github.com/Faizal-Asep/test-tech/models"
	repositorie "github.com/Faizal-Asep/test-tech/repositories"
	"github.com/gin-gonic/gin"
)

func ApiDelete(c *gin.Context) {

	id := c.Param("id")

	client := repositorie.NewHandPhone()
	err := client.Delete(c.Request.Context(), id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.HttpResponse{Message: err.Error()})
		return
	}
	go func() {
		model.Notif <- model.Notification{
			Command: "delete",
			Data:    id,
		}
	}()
	c.JSON(http.StatusOK, model.HttpResponse{Message: "Sukses"})

}
