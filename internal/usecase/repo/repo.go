package repo

import (
	"github.com/WBear29/go-restapi-server-template/pkg/logger"
	"github.com/WBear29/go-restapi-server-template/pkg/rdb"
)

type Repository struct {
	DBHandler *rdb.DBHandler
	l         *logger.Logger
}

func New(dbh *rdb.DBHandler, l *logger.Logger) *Repository {
	return &Repository{dbh, l}
}
