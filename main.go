package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	common "github.com/yusuftatli/go-template-with-redis-gorm-api/common"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/handler"
	initial "github.com/yusuftatli/go-template-with-redis-gorm-api/initial"
)

var (
	TemplateHandler *handler.TemplateHandler
	loginHandler    *handler.LoginHandler
)

var (
	router = gin.Default()
)

func main() {
	env := common.GetEnvironment()
	// gin.SetMode(gin.DebugMode)
	// Logger
	logger := common.InitLogger(env.Debug)
	defer func() {
		_ = logger.Sync()
	}()

	initial.InitRoutes(router)

	router.Run(fmt.Sprintf("%s:%s", env.ApplicationHost, env.Port))
}
