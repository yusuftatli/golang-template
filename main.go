package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/internal/handler"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/internal/initial"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/pkg/config"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/pkg/database"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/pkg/logging"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/pkg/middleware"
)

var (
	TemplateHandler *handler.HandlerExample
)

var (
	router = gin.Default()
)

func main() {
	env := config.GetEnvironment()

	// gin.SetMode(gin.DebugMode)

	// Logger
	logger := logging.InitLogger(env.Debug)
	defer logger.Sync()

	// Database
	db := database.ConnectDB(env.Postgres)
	defer database.CloseDBConnection(db)
	database.AutoMigrate(db)

	//Middlewares
	router.Use(middleware.HttpLoggingMiddleware())
	router.Use(gin.Recovery())

	initial.InitRoutes(router)

	router.Run(fmt.Sprintf("%s:%s", env.ApplicationHost, env.Port))
}
