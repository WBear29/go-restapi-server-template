package usecase

import (
	"context"

	"github.com/WBear29/go-restapi-server-template/internal/entity"
	"github.com/WBear29/go-restapi-server-template/pkg/apperr"
)

// SampleUseCase -,
type SampleUseCase struct {
	repo Repository
}

// New -,
func NewSample(r Repository) *SampleUseCase {
	return &SampleUseCase{r}
}

// PostSample
func (uc *SampleUseCase) PostSample(ctx context.Context, sample entity.Sample) (entity.Sample, apperr.Err) {
	return uc.repo.InsertSample(ctx, sample)
}

// GetSamples
func (uc *SampleUseCase) GetSamples(ctx context.Context) ([]entity.Sample, apperr.Err) {
	return uc.repo.SelectSamples(ctx)
}

// UpdateSample
func (uc *SampleUseCase) PatchSample(ctx context.Context, sample entity.Sample) (entity.Sample, apperr.Err) {
	return uc.repo.UpdateSample(ctx, sample)
}
