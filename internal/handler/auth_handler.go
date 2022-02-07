package handler

import "github.com/yusuftatli/go-template-with-redis-gorm-api/internal/repositories"

type AuthHandler struct {
	UserRepository repositories.UserRepositoryInterface
}

// func NewAuthHandler(userRepository repositories.UserRepositoryInterface) *TemplateHandler {
// 	return &AuthHandler{
// 		UserRepository: userRepository,
// 	}
// }
