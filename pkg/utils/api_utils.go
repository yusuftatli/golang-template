package utils

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/internal/models"
	"go.uber.org/zap"
)

func ApiError(c *gin.Context, err *models.APIError) {
	zap.S().Error(err.ErrorMessage)
	resp, _ := err.MarshalBinary()
	c.Header("Content-Type", "application/json")
	c.Writer.Write(resp)
}

func ApiDeleted(c *gin.Context, deletedMessage string) {
	var deleted = make(map[string]string)
	deleted["message"] = deletedMessage
	binaryApiMessage, _ := json.Marshal(deleted)
	c.Header("Content-Type", "application/json")
	c.Writer.Write(binaryApiMessage)
}

func ResponseWrite(c *gin.Context, data interface{}) {
	responseBytes, _ := json.Marshal(data)
	c.Header("Content-Type", "application/json")
	c.Writer.Write(responseBytes)
}
