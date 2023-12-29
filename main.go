package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	handler "github.com/Faizal-Asep/test-tech/handlers"
	middleware "github.com/Faizal-Asep/test-tech/middlewares"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {

	// Router.Run() // listen and serve on 0.0.0.0:8080
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8080),
		Handler:        Router(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// run server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown Server ...")
}

func Router() *gin.Engine {

	r := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.LoadHTMLGlob("web/templates/*")
	r.Static("/assets", "./web/static")

	r.GET("/", handler.SigninPage)
	r.GET("/oauth/redirect", handler.SigninRedirect)
	r.Use(middleware.ValidateSession())

	r.GET("/input", handler.InputPage)
	r.GET("/output", handler.OutputPage)

	r.GET("/echo", handler.Socket)

	api := r.Group("/api")
	api.GET("/handphone/auto", handler.ApiAuto)
	api.POST("/handphone", handler.ApiSave)
	api.PUT("/handphone/:id", handler.ApiUpdate)
	api.DELETE("/handphone/:id", handler.ApiDelete)
	api.GET("/handphone", handler.ApiGet)
	api.GET("/handphone/:id", handler.ApiGetId)

	return r
}
