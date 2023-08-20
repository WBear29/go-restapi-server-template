package v1

import (
	"github.com/WBear29/go-restapi-server-template/internal/usecase"
	"github.com/WBear29/go-restapi-server-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

type sampleRoutes struct {
	uc usecase.Sample
	l  *logger.Logger
}

func newSample(handler *gin.RouterGroup, s usecase.Sample, l *logger.Logger) {
	r := &sampleRoutes{s, l}
	{
		handler.POST("/sample", r.postSample)
	}
}

func (r *sampleRoutes) postSample(c *gin.Context) {
}
