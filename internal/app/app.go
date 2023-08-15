package app

import (
	"context"

	"github.com/WBear29/go-restapi-server-template/config"
	"github.com/WBear29/go-restapi-server-template/pkg/logger"
	"github.com/WBear29/go-restapi-server-template/pkg/rdb"
	"github.com/WBear29/go-restapi-server-template/pkg/tracing"
	"go.opentelemetry.io/otel"
)

// Run application
func Run(cfg *config.Config, version string) {
	l := logger.New(cfg.Log.Level)

	// tracing
	tp, tErr := tracing.NewTraceProvider(
		cfg.App.Name,
		version,
		cfg.App.Env,
		cfg.Tracing.Endpoint,
	)
	if tErr != nil {
		l.Warn("Tracing error: %v", tErr)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			l.Fatal("Tracing shutdown error: %v", err)
		}
	}()
	otel.SetTracerProvider(tp)

	// database
	dbHandler, dbErr := rdb.NewDBHandler(
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Name,
		cfg.DB.Pass,
		cfg.DB.Retry,
	)
	if dbErr != nil {
		l.Fatal("Database error: %v", dbErr)
	}
	// migrate
	l.Info("Migrating database... from %s", cfg.DB.MigrationDir)
	if err := dbHandler.Migrate(cfg.DB.MigrationDir); err != nil {
		l.Error("Migration error: %v", err)
	}
}
