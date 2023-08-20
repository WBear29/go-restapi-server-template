package usecase

// SampleUseCase -,
type SampleUseCase struct {
	repo Repository
}

// New -,
func NewSample(r Repository) *SampleUseCase {
	return &SampleUseCase{r}
}
