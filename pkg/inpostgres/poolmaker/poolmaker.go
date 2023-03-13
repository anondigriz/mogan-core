package poolmaker

import (
	"context"
	"fmt"
	"github.com/anondigriz/mogan-core/pkg/logger"
	pgxUUID "github.com/jackc/pgx-gofrs-uuid"
	pgxZap "github.com/jackc/pgx-zap"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"go.uber.org/zap"
)

func convertLogLevel(logLevel logger.LogLevel) (tracelog.LogLevel, error) {
	switch logLevel {
	case logger.LogLevelNone:
		return tracelog.LogLevelNone, nil
	case logger.LogLevelTrace:
		return tracelog.LogLevelTrace, nil
	case logger.LogLevelDebug:
		return tracelog.LogLevelDebug, nil
	case logger.LogLevelInfo:
		return tracelog.LogLevelInfo, nil
	case logger.LogLevelWarn:
		return tracelog.LogLevelWarn, nil
	case logger.LogLevelError:
		return tracelog.LogLevelError, nil
	default:
		return 0, fmt.Errorf("unknown database logging level: %d", logLevel)
	}
}

func New(ctx context.Context, lg *zap.Logger, dsn string, logLevel logger.LogLevel) (*pgxpool.Pool, error) {
	dbCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		lg.Error("unable to parse config", zap.Error(err))
		return nil, err
	}

	traceLogLevel, err := convertLogLevel(logLevel)
	if err != nil {
		lg.Error("unknown database logging level", zap.Error(err))
		return nil, err
	}

	dbCfg.BeforeConnect = func(ctx context.Context, config *pgx.ConnConfig) error {
		config.Tracer = &tracelog.TraceLog{
			Logger:   pgxZap.NewLogger(lg),
			LogLevel: traceLogLevel,
		}
		return nil
	}
	dbCfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}
	dbPool, err := pgxpool.NewWithConfig(ctx, dbCfg)
	if err != nil {
		lg.Error("unable to create connection pool", zap.Error(err))
		return nil, err
	}
	return dbPool, nil
}
