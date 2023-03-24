package database

import (
	"database/sql"
	"ethernal/explorer-api/configuration"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// Initialize returns an initialized bun.DB instance.
func Initialize(configuration *configuration.Configuration) *bun.DB {

	sqlDB := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithAddr(configuration.DbHost+":"+configuration.DbPort),
		pgdriver.WithUser(configuration.DbUser),
		pgdriver.WithPassword(configuration.DbPassword),
		pgdriver.WithDatabase(configuration.DbName),
		pgdriver.WithInsecure(true),
		pgdriver.WithReadTimeout(time.Duration(int(configuration.CallTimeoutInSeconds))*time.Second),
	))

	return bun.NewDB(sqlDB, pgdialect.New())
}
