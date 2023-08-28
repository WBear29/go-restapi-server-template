package v1

import (
	"net/http"

	"github.com/WBear29/go-restapi-server-template/internal/controller/http/model"
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
		handler.POST("/samples", r.postSample)
		handler.GET("/samples", r.getSamples)
	}
}

func (r *sampleRoutes) postSample(c *gin.Context) {
	enSample, appErr := model.ValidateSample(c)
	if appErr != nil {
		r.l.Error(appErr.Log())
		model.ErrorResponse(c, appErr)
		return
	}
	sample, appErr := r.uc.PostSample(c, enSample)
	if appErr != nil {
		r.l.Error(appErr.Log())
		model.ErrorResponse(c, appErr)
		return
	}

	c.JSON(http.StatusOK, model.ResSampleFrom(sample))
}

func (r *sampleRoutes) getSamples(c *gin.Context) {
	samples, appErr := r.uc.GetSamples(c)
	if appErr != nil {
		r.l.Error(appErr.Log())
		model.ErrorResponse(c, appErr)
		return
	}

	res := model.ResSamplesFrom(samples)
	c.JSON(http.StatusOK, res)
}
