package migrator

import (
	"database/sql"
	"embed"
	"github.com/anondigriz/mogan-core/pkg/gooselogger"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

type Migrator struct {
	lg    *zap.Logger
	dsn   string
	embed *embed.FS
}

func New(lg *zap.Logger, dsn string, embed *embed.FS) *Migrator {
	return &Migrator{lg: lg, dsn: dsn, embed: embed}
}

func (m *Migrator) Migrate() error {
	db, err := sql.Open("pgx", m.dsn)
	if err != nil {
		m.lg.Error("unable to connect to database")
		return err
	}
	defer db.Close()

	goose.SetLogger(gooselogger.New(m.lg))
	goose.SetBaseFS(m.embed)

	if err = goose.SetDialect("postgres"); err != nil {
		m.lg.Error("SetDialect was failed", zap.Error(err))
		return err
	}

	if err = goose.Up(db, "migrations"); err != nil {
		m.lg.Error("init database was failed", zap.Error(err))
		return err
	}

	return nil
}
