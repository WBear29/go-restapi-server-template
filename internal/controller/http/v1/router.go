package v1

import (
	"github.com/WBear29/go-restapi-server-template/internal/usecase"
	"github.com/WBear29/go-restapi-server-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

// NewRouter -.
func NewRouter(handler *gin.Engine, l *logger.Logger, s usecase.Sample) {
	// Routers
	h := handler.Group("/v1")
	{
		newSample(h, s, l)
	}
}
