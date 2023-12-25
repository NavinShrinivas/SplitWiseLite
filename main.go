package main

import (
	"os"
	"strings"

	envLoader "SplitWiseLite/envLoader"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {

	log.Println("Starting user handle services...")

	envLoader.CheckAndSetVariables()

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	config.AllowCredentials = true
	r.Use(cors.New(config))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Server is active and listening on port" + os.Getenv("PORT"),
		},
		)
	})
	v1_grp := r.Group("/api/v1")
	{
		V1_Routes(v1_grp)
	}

	s := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	log.Println("Services started and running at port" + s.Addr)
	s.ListenAndServe()
}
