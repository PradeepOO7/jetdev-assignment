package main

import (
	"blogs/config"
	"blogs/controller"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()
	config.InitDB()
	engine := initServer()
	PORT := viper.GetInt("SERVER_PORT")
	server := http.Server{
		Addr:         fmt.Sprintf(":%v", PORT),
		Handler:      engine,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  3 * time.Second,
	}

	log.Printf("Server started on port %d\n", PORT)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Error Starting Server on port %v %s \n", PORT, err)
	}
}
func initServer() *gin.Engine {
	engine := gin.New()
	engine.Use(cors.Default())
	engine.Use(gin.Recovery())
	engine.Use(gin.ErrorLogger())
	controller.AddRoutes(engine)
	return engine
}
