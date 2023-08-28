package model

import (
	"github.com/WBear29/go-restapi-server-template/internal/entity"
	"github.com/WBear29/go-restapi-server-template/pkg/apperr"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

func BindSample(c *gin.Context) (Sample, apperr.Err) {
	var sample Sample
	if err := c.BindJSON(&sample); err != nil {
		return Sample{}, apperr.New(apperr.ERROR_INVALID_REQUEST_BODY, err.Error())
	}
	return sample, nil
}

func (s Sample) toEntity() entity.Sample {
	return entity.Sample{
		Name: s.Name,
	}
}

func (s Sample) Validate() (entity.Sample, apperr.Err) {
	err := validation.Errors{
		"name": validation.Validate(s.Name, validation.Required, validation.RuneLength(1, 255)),
	}.Filter()
	if err != nil {
		return entity.Sample{}, apperr.New(apperr.ERROR_INVALID_REQUEST_BODY, err.Error())
	}
	return s.toEntity(), nil
}

func ValidateSample(c *gin.Context) (entity.Sample, apperr.Err) {
	sample, err := BindSample(c)
	if err != nil {
		return entity.Sample{}, err
	}
	return sample.Validate()
}

func ResSampleFrom(enSample entity.Sample) ResSample {
	return ResSample{
		Id:         enSample.ID.String(),
		Name:       enSample.Name,
		TimeStamps: TimeStampsFrom(enSample.TimeStamps),
	}
}

func ResSamplesFrom(enSamples []entity.Sample) []ResSample {
	res := make([]ResSample, len(enSamples))
	for i, enSample := range enSamples {
		res[i] = ResSampleFrom(enSample)
	}
	return res
}
