package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/internal/entities"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/internal/models"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/internal/repositories"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/pkg/utils"
)

type LoginHandler struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewLoginHandler(userRepository repositories.UserRepositoryInterface) *LoginHandler {
	return &LoginHandler{
		UserRepository: userRepository,
	}
}

func (handler *LoginHandler) Login(c *gin.Context) {
	dto := &models.PostTemplate{}
	err := c.Bind(dto)
	if err != nil {
		// utils.ApiError(c, &models.APIError{ErrorCode: http.StatusBadGateway, ErrorMessage: "model incorrect"})
		return
	}

	missingRequiredFields := dto.Validate(strfmt.Default)
	if missingRequiredFields != nil {
		// utils.ApiError(c, &models.APIError{ErrorCode: http.StatusInternalServerError, ErrorMessage: fmt.Sprintf("Required parameters are missing. %s", missingRequiredFields)})
		return
	}

	row := &entities.User{Name: dto.Name, Password: ""}
	row, err = handler.UserRepository.FindFirst(row)

	if err != nil {
		// utils.ApiError(c, &models.APIError{ErrorCode: http.StatusInternalServerError, ErrorMessage: "Error unmarshalling request. Check your request."})
		return
	}

	utils.ResponseWrite(c, row)
}
