package server

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	port   string
	engine *gin.Engine
}

func NewHttpServer(port string) *HttpServer {
	engine := gin.Default()

	engine.Use(cors.New(cors.Config{
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowAllOrigins: true,
	}))

	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	server := &HttpServer{
		port:   port,
		engine: engine,
	}

	return server
}

func (s *HttpServer) Run() error {
	return s.engine.Run(fmt.Sprintf(":%s", s.port))
}
