package tenant

import (
	"context"
	"fmt"
	"time"

	"github.com/influxdata/influxdb"
	"go.uber.org/zap"
)

type OrgLogger struct {
	logger     *zap.Logger
	orgService influxdb.OrganizationService
}

// NewOrgLogger returns a logging service middleware for the Organization Service.
func NewOrgLogger(log *zap.Logger, s influxdb.OrganizationService) *OrgLogger {
	return &OrgLogger{
		logger:     log,
		orgService: s,
	}
}

var _ influxdb.OrganizationService = (*OrgLogger)(nil)

func (l *OrgLogger) CreateOrganization(ctx context.Context, u *influxdb.Organization) (err error) {
	defer func(start time.Time) {
		dur := zap.Duration("took", time.Since(start))
		if err != nil {
			l.logger.Error("failed to create org", zap.Error(err), dur)
			return
		}
		l.logger.Info("org create", dur)
	}(time.Now())
	return l.orgService.CreateOrganization(ctx, u)
}

func (l *OrgLogger) FindOrganizationByID(ctx context.Context, id influxdb.ID) (u *influxdb.Organization, err error) {
	defer func(start time.Time) {
		dur := zap.Duration("took", time.Since(start))
		if err != nil {
			msg := fmt.Sprintf("failed to find org with ID %v", id)
			l.logger.Error(msg, zap.Error(err), dur)
			return
		}
		l.logger.Info("org find by ID", dur)
	}(time.Now())
	return l.orgService.FindOrganizationByID(ctx, id)
}

func (l *OrgLogger) FindOrganization(ctx context.Context, filter influxdb.OrganizationFilter) (u *influxdb.Organization, err error) {
	defer func(start time.Time) {
		dur := zap.Duration("took", time.Since(start))
		if err != nil {
			l.logger.Error("failed to find org matching the given filter", zap.Error(err), dur)
			return
		}
		l.logger.Info("org find", dur)
	}(time.Now())
	return l.orgService.FindOrganization(ctx, filter)
}

func (l *OrgLogger) FindOrganizations(ctx context.Context, filter influxdb.OrganizationFilter, opt ...influxdb.FindOptions) (orgs []*influxdb.Organization, n int, err error) {
	defer func(start time.Time) {
		dur := zap.Duration("took", time.Since(start))
		if err != nil {
			l.logger.Error("failed to find org matching the given filter", zap.Error(err), dur)
			return
		}
		l.logger.Info("orgs find", dur)
	}(time.Now())
	return l.orgService.FindOrganizations(ctx, filter)
}

func (l *OrgLogger) UpdateOrganization(ctx context.Context, id influxdb.ID, upd influxdb.OrganizationUpdate) (u *influxdb.Organization, err error) {
	defer func(start time.Time) {
		dur := zap.Duration("took", time.Since(start))
		if err != nil {
			l.logger.Error("failed to update org", zap.Error(err), dur)
			return
		}
		l.logger.Info("org update", dur)
	}(time.Now())
	return l.orgService.UpdateOrganization(ctx, id, upd)
}

func (l *OrgLogger) DeleteOrganization(ctx context.Context, id influxdb.ID) (err error) {
	defer func(start time.Time) {
		dur := zap.Duration("took", time.Since(start))
		if err != nil {
			msg := fmt.Sprintf("failed to delete org with ID %v", id)
			l.logger.Error(msg, zap.Error(err), dur)
			return
		}
		l.logger.Info("org delete", dur)
	}(time.Now())
	return l.orgService.DeleteOrganization(ctx, id)
}
