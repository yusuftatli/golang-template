package initial

import (
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/internal/handler"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/pkg/middleware"
)

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func InitRoutes(r *gin.Engine) {
	// Handlers
	TemplateHandler := handler.NewHandlerExample()

	authMiddleware, _ := middleware.Auth()

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
		auth.POST("/delete:id", TemplateHandler.Delete)
		auth.GET("/list", TemplateHandler.FindAll)
	}

}
