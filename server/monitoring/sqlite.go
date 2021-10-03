package monitoring

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	monitoring "github.com/cloudradar-monitoring/rport/db/migration/monitoring"
	"github.com/cloudradar-monitoring/rport/db/sqlite"
	chshare "github.com/cloudradar-monitoring/rport/share"
	"github.com/cloudradar-monitoring/rport/share/models"
)

type DBProvider interface {
	CreateMeasurement(ctx context.Context, measurement *models.Measurement) error
	DeleteMeasurementsOlderThan(ctx context.Context, days int64) error
	Close() error
}

type SqliteProvider struct {
	db     *sqlx.DB
	logger *chshare.Logger
}

func NewSqliteProvider(dbPath string, logger *chshare.Logger) (DBProvider, error) {
	db, err := sqlite.New(dbPath, monitoring.AssetNames(), monitoring.Asset)
	if err != nil {
		return nil, fmt.Errorf("failed to create monitoring DB instance: %v", err)
	}

	logger.Infof("initialized database at %s", dbPath)

	return &SqliteProvider{db: db, logger: logger}, nil
}

func (p *SqliteProvider) CreateMeasurement(ctx context.Context, measurement *models.Measurement) error {
	_, err := p.db.NamedExecContext(
		ctx,
		"INSERT INTO measurements (client_id, timestamp, cpu_usage_percent, memory_usage_percent, io_usage_percent, processes, mountpoints) "+
			"VALUES (:client_id, :timestamp, :cpu_usage_percent, :memory_usage_percent, :io_usage_percent, :processes, :mountpoints)",
		measurement,
	)
	return err
}

func (p *SqliteProvider) DeleteMeasurementsOlderThan(ctx context.Context, compare int64) error {
	_, err := p.db.ExecContext(ctx, "DELETE FROM measurements WHERE  timestamp < ?", compare)
	return err
}
func (p *SqliteProvider) Close() error {
	return p.db.Close()
}
