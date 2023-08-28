package rdb

import (
	"fmt"
	"time"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBHandler struct {
	Conn *gorm.DB
}

func connect(host, port, user, name, pass string, count uint) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, name, port)
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}))
	if err != nil {
		if count == 0 {
			return nil, err
		}
		time.Sleep(time.Second)
		count--
		return connect(host, port, user, name, pass, count)
	}
	if otelErr := conn.Use(otelgorm.NewPlugin()); otelErr != nil {
		return nil, otelErr
	}
	return conn, nil
}

func NewDBHandler(host, port, user, name, pass string, count uint) (*DBHandler, error) {
	conn, err := connect(host, port, user, name, pass, count)
	if err != nil {
		return nil, err
	}
	return &DBHandler{Conn: conn}, nil
}
