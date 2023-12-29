package handler

import (
	"math/rand"
	"net/http"

	repositorie "github.com/Faizal-Asep/test-tech/repositories"
	"github.com/gin-gonic/gin"
)

func ApiAuto(c *gin.Context) {
	phoneNumber := []string{"081234567890", "081234567891", "081234567892", "081234567893", "081234567894", "081234567895", "081234567896", "081234567897", "081234567898", "081234567899", "081234567901", "081234567902", "081234567903", "081234567904", "081234567905", "081234567906", "081234567907", "081234567908", "081234567909", "081234567910", "081234567911", "081234567912", "081234567913", "081234567914", "081234567915"}
	provider := []string{"Smartfren", "XL", "Telkomsel", "Indosat"}
	n1 := rand.Intn(len(phoneNumber))
	n2 := rand.Intn(len(provider))
	p := repositorie.HandPhone{
		Number:   phoneNumber[n1],
		Provider: provider[n2],
	}

	c.JSON(http.StatusOK, p)
}
