package initial

import (
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	common "github.com/yusuftatli/go-template-with-redis-gorm-api/common"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/common/database"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/handler"
	midddleware "github.com/yusuftatli/go-template-with-redis-gorm-api/midddleware"
	repositories "github.com/yusuftatli/go-template-with-redis-gorm-api/repository"
)

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

var identityKey = "id"

func InitRoutes(r *gin.Engine) {
	env := common.GetEnvironment()
	// Database
	db := database.ConnectDB(env.Postgres)
	defer database.CloseDBConnection(db)
	database.AutoMigrate(db)
	// Repositories
	templateRepository := repositories.NewTemplateRepository(db)

	// Handlers
	TemplateHandler := handler.NewATemplateHandler(templateRepository)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authMiddleware, _ := midddleware.Auth()
	r.POST("/login", authMiddleware.LoginHandler)
	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/api")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.POST("/create", TemplateHandler.Insert)
		auth.POST("/update:id", TemplateHandler.Update)
		auth.POST("/delete:id", TemplateHandler.Detele)
		auth.GET("/list", TemplateHandler.FindAll)
	}

}
