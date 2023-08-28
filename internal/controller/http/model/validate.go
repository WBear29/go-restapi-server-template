package model

import (
	"github.com/WBear29/go-restapi-server-template/pkg/apperr"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ValidateID(c *gin.Context) (uuid.UUID, apperr.Err) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.UUID{}, apperr.New(apperr.ERROR_INVALID_REQUEST_BODY, err.Error())
	}
	return id, nil
}
