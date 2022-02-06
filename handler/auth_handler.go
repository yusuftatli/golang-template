package handler

import repositories "github.com/yusuftatli/go-template-with-redis-gorm-api/repository"

type AuthHandler struct {
	UserRepository repositories.UserRepositoryInterface
}

// func NewAuthHandler(userRepository repositories.UserRepositoryInterface) *TemplateHandler {
// 	return &AuthHandler{
// 		UserRepository: userRepository,
// 	}
// }
