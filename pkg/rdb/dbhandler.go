package rdb

import (
	"fmt"
	"time"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBHandler struct {
	Conn *gorm.DB
}

func connect(host, port, user, name, pass string, count uint) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	conn, err := gorm.Open(mysql.Open(dsn))
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
