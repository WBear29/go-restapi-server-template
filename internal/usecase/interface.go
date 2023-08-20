package usecase

import (
	"context"

	"github.com/WBear29/go-restapi-server-template/internal/entity"
	"github.com/WBear29/go-restapi-server-template/pkg/apperr"
)

type (
	// ImplDir: ./
	Sample interface {
		PostSample(ctx context.Context, sample entity.Sample) (entity.Sample, apperr.Err)
	}

	// ImplDir: ./repo
	Repository interface {
		InsertSample(ctx context.Context, sample entity.Sample) (entity.Sample, apperr.Err)
	}
)
