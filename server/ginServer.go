package server

import (
	"be-authen/config"
	"be-authen/database"
	"be-authen/di"
	"log"

	"github.com/gin-gonic/gin"
)

type ginServer struct {
	app  *gin.Engine
	db   database.Database
	conf *config.Config
}

func NewGinServer(conf *config.Config, db database.Database, container *di.Container) Server {
	ginApp := gin.Default()

	RegisterRoutes(ginApp, container)

	return &ginServer{
		app:  ginApp,
		db:   db,
		conf: conf,
	}
}

func (s *ginServer) Start() {
	// Use Middleware Recovery and Logger of Gin
	s.app.Use(gin.Recovery())
	s.app.Use(gin.Logger())

	// เรียกใช้งาน HTTP Handlers

	s.initializeAuthenHttpHandler()

	// เริ่มเซิร์ฟเวอร์บน port 8080
	serverUrl := ":8080"
	if err := s.app.Run(serverUrl); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (s *ginServer) initializeAuthenHttpHandler() {
	// นี่คือที่ที่คุณจะเพิ่ม handler ที่เกี่ยวกับ Authen ในอนาคต
}
