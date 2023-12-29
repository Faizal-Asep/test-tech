package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	model "github.com/Faizal-Asep/test-tech/models"
	repositorie "github.com/Faizal-Asep/test-tech/repositories"
	util "github.com/Faizal-Asep/test-tech/utils"
	"github.com/gin-gonic/gin"
)

func ApiSave(c *gin.Context) {
	var in repositorie.HandPhone

	if c.GetHeader("encripted") == "true" {
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, model.HttpResponse{Message: err.Error()})
			return
		}
		cipher := util.NewCipherClient(util.Config.EncriptKey, util.Config.EncriptIv)
		b, err := cipher.GetAESDecrypted(string(jsonData))
		if err != nil {
			c.JSON(http.StatusBadRequest, model.HttpResponse{Message: err.Error()})
			return
		}
		if err := json.Unmarshal(b, &in); err != nil {
			c.JSON(http.StatusBadRequest, model.HttpResponse{Message: err.Error()})
			return
		}
	} else {
		if err := c.ShouldBindJSON(&in); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, model.HttpResponse{Message: err.Error()})
			return
		}
	}

	client := repositorie.NewHandPhone()
	p, err := client.Save(c.Request.Context(), in)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.HttpResponse{Message: err.Error()})
		return
	}

	go func() {
		model.Notif <- model.Notification{
			Command: "save",
			Data:    p,
		}
	}()

	c.JSON(http.StatusCreated, p)

}
