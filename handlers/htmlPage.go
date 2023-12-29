package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	model "github.com/Faizal-Asep/test-tech/models"
	repositorie "github.com/Faizal-Asep/test-tech/repositories"
	util "github.com/Faizal-Asep/test-tech/utils"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InputPage(c *gin.Context) {
	// c.Redirect(http.StatusFound, "/signin")
	c.HTML(http.StatusOK, "input.html", gin.H{
		"title":      "Input No Handphone",
		"encriptKey": util.Config.EncriptKey,
		"encriptIv":  util.Config.EncriptIv,
	})
}

func OutputPage(c *gin.Context) {

	var ganjil, genap []repositorie.HandPhone

	client := repositorie.NewHandPhone()
	handphones, err := client.Get(c.Request.Context())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.HttpResponse{Message: err.Error()})
		return
	}

	for _, handphone := range handphones {
		x, _ := strconv.Atoi(handphone.Number[len(handphone.Number)-1:])
		if x%2 == 0 {
			genap = append(genap, handphone)
		} else {
			ganjil = append(ganjil, handphone)
		}
	}

	c.HTML(http.StatusOK, "output.html", gin.H{
		"title":  "Output Data",
		"genap":  genap,
		"ganjil": ganjil,
	})
}

func SigninPage(c *gin.Context) {
	b := make([]byte, 32)
	rand.Read(b)

	state := base64.StdEncoding.EncodeToString(b)
	session := sessions.Default(c)
	session.Set("state", state)
	session.Save()
	c.HTML(http.StatusOK, "signin.html", gin.H{
		"title":     "Signin website",
		"googleUrl": util.OauthConfGl.AuthCodeURL(state),
	})
}

func SigninRedirect(c *gin.Context) {
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	if retrievedState != c.Query("state") {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("invalid session state: %s", retrievedState))
		return
	}

	tok, err := util.OauthConfGl.Exchange(context.Background(), c.Query("code"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := util.OauthConfGl.Client(context.Background(), tok)
	email, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer email.Body.Close()

	data, _ := io.ReadAll(email.Body)
	var userInfo model.UserInfo

	if err := json.Unmarshal(data, &userInfo); err != nil {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("invalid user info: %s", retrievedState))
		return
	}

	session.Set("sub", userInfo.Sub)
	session.Set("picture", userInfo.Picture)
	session.Set("email", userInfo.Email)
	session.Set("emailVerified", userInfo.EmailVerified)
	session.Save()

	c.Redirect(http.StatusFound, "/output")
}
