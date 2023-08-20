package model

import (
	"github.com/WBear29/go-restapi-server-template/pkg/apperr"
	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, appErr apperr.Err) {
	c.AbortWithStatusJSON(
		appErr.Status(),
		ModelError{
			ErrorCode: int32(appErr.Code()),
			ErrorInfo: ErrorInfo{
				Message: appErr.Message(),
			},
		},
	)
}
