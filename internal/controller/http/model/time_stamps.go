package model

import "github.com/WBear29/go-restapi-server-template/internal/entity"

func TimeStampsFrom(enTimestamp entity.TimeStamps) TimeStamps {
	return TimeStamps{
		CreatedAt: enTimestamp.CreatedAt,
		UpdatedAt: enTimestamp.UpdatedAt,
	}
}
