package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/entities"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/models"
	repositories "github.com/yusuftatli/go-template-with-redis-gorm-api/repository"
	utils "github.com/yusuftatli/go-template-with-redis-gorm-api/utils"
)

type TemplateHandler struct {
	TemplateRepository repositories.TemplateRepositoryInterface
}

func NewATemplateHandler(templateRepository repositories.TemplateRepositoryInterface) *TemplateHandler {
	return &TemplateHandler{
		TemplateRepository: templateRepository,
	}
}

func (handler *TemplateHandler) Insert(c *gin.Context) {
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

	row := &entities.Template{Name: dto.Name, Size: dto.Size, Type: dto.Type}
	row, err = handler.TemplateRepository.Insert(row)

	if err != nil {
		// utils.ApiError(c, &models.APIError{ErrorCode: http.StatusInternalServerError, ErrorMessage: "Error unmarshalling request. Check your request."})
		return
	}

	utils.ResponseWrite(c, row)
}

func (handler *TemplateHandler) Update(c *gin.Context) {
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

	row := &entities.Template{ID: dto.ID}
	row, err = handler.TemplateRepository.Insert(row)

	if err != nil {
		// utils.ApiError(c, &models.APIError{ErrorCode: http.StatusInternalServerError, ErrorMessage: "Error unmarshalling request. Check your request."})
		return
	}

	utils.ResponseWrite(c, row)
}

func (handler *TemplateHandler) Detele(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		// utils.ApiError(c, &models.APIError{ErrorCode: http.StatusInternalServerError, ErrorMessage: "Tempalte id is required"})
		return
	}

	row := &entities.Template{ID: uint64(id)}
	err = handler.TemplateRepository.Delete(row)

	if err != nil {
		// utils.ApiError(c, &models.APIError{ErrorCode: http.StatusInternalServerError, ErrorMessage: "Error unmarshalling request. Check your request."})
		return
	}

	// s :=common.err_message_template

	// utils.ApiDeleted(c, common.err_message_template)
}

func (handler *TemplateHandler) FindAll(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		// utils.ApiError(c, &models.APIError{ErrorCode: http.StatusInternalServerError, ErrorMessage: "Tempalte id is required"})
		return
	}

	row := &entities.Template{ID: uint64(id)}
	err = handler.TemplateRepository.Delete(row)

	if err != nil {
		// utils.ApiError(c, &models.APIError{ErrorCode: http.StatusInternalServerError, ErrorMessage: "Error unmarshalling request. Check your request."})
		return
	}

	// s :=common.err_message_template

	// utils.ApiDeleted(c, common.err_message_template)
}
