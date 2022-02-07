package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/internal/entities"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/internal/models"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/internal/repositories"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/pkg/database"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/pkg/utils"
)

type HandlerExample struct {
	RepositoryExample repositories.IRepositoryExample
}

func NewHandlerExample() *HandlerExample {
	return &HandlerExample{
		RepositoryExample: repositories.NewRepositoryExample(database.GetPostgresDB()),
	}
}

func (handler *HandlerExample) Insert(c *gin.Context) {
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

	row := &entities.Example{Name: dto.Name, Size: dto.Size, Type: dto.Type}
	row, err = handler.RepositoryExample.Insert(row)

	if err != nil {
		// utils.ApiError(c, &models.APIError{ErrorCode: http.StatusInternalServerError, ErrorMessage: "Error unmarshalling request. Check your request."})
		return
	}

	utils.ResponseWrite(c, row)
}

func (handler *HandlerExample) Update(c *gin.Context) {
	dto := &models.PacthTemplate{}
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

	row := &entities.Example{ID: dto.ID}
	row, err = handler.RepositoryExample.Insert(row)

	if err != nil {
		// utils.ApiError(c, &models.APIError{ErrorCode: http.StatusInternalServerError, ErrorMessage: "Error unmarshalling request. Check your request."})
		return
	}

	utils.ResponseWrite(c, row)
}

func (handler *HandlerExample) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		// utils.ApiError(c, &models.APIError{ErrorCode: http.StatusInternalServerError, ErrorMessage: "Tempalte id is required"})
		return
	}

	row := &entities.Example{ID: uint64(id)}
	err = handler.RepositoryExample.Delete(row)

	if err != nil {
		// utils.ApiError(c, &models.APIError{ErrorCode: http.StatusInternalServerError, ErrorMessage: "Error unmarshalling request. Check your request."})
		return
	}

	// s :=common.err_message_template

	// utils.ApiDeleted(c, common.err_message_template)
}

func (handler *HandlerExample) FindAll(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		// utils.ApiError(c, &models.APIError{ErrorCode: http.StatusInternalServerError, ErrorMessage: "Tempalte id is required"})
		return
	}

	row := &entities.Example{ID: uint64(id)}
	err = handler.RepositoryExample.Delete(row)

	if err != nil {
		// utils.ApiError(c, &models.APIError{ErrorCode: http.StatusInternalServerError, ErrorMessage: "Error unmarshalling request. Check your request."})
		return
	}

	// s :=common.err_message_template

	// utils.ApiDeleted(c, common.err_message_template)
}
