package repo

import (
	"context"
	"time"

	"github.com/WBear29/go-restapi-server-template/internal/entity"
	"github.com/WBear29/go-restapi-server-template/pkg/apperr"
	"github.com/google/uuid"
)

type sample struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s sample) toEntity() entity.Sample {
	return entity.Sample{
		ID:   s.ID,
		Name: s.Name,
		TimeStamps: entity.TimeStamps{
			CreatedAt: s.CreatedAt,
			UpdatedAt: s.UpdatedAt,
		},
	}
}

func (r *Repository) InsertSample(ctx context.Context, enSample entity.Sample) (entity.Sample, apperr.Err) {
	sample := sample{
		Name: enSample.Name,
	}
	if err := r.DBHandler.Conn.WithContext(ctx).Create(&sample).Error; err != nil {
		return entity.Sample{}, apperr.New(apperr.ERROR_INTERNAL_SERVER_ERROR, err.Error())
	}
	return sample.toEntity(), nil
}
