package db

import (
	"common/consts"
	"context"
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/cloudwego/kitex/pkg/klog"

	"time"
)

// SlowQueryDriver is a driver that logs all driver operations.
type SlowQueryDriver struct {
	dialect.Driver                               // underlying driver.
	Log            func(context.Context, ...any) // log function. defaults to log.Println.
	logLevel       string
}

// Exec logs its params and calls the underlying driver Exec method.
func (d *SlowQueryDriver) Exec(ctx context.Context, query string, args, v any) error {
	start := time.Now()
	err := d.Driver.Exec(ctx, query, args, v)
	duration := time.Since(start)
	if d.logLevel == "debug" {
		klog.CtxInfof(ctx, "driver.Exec: query=%v args=%v time=%v", query, args, duration)
	} else if duration > consts.SlowQueryThreshold {
		d.Log(ctx, fmt.Sprintf("driver.Exec: query=%v args=%v time=%v", query, args, duration))
	}
	return err
}

// ExecContext logs its params and calls the underlying driver ExecContext method if it is supported.
func (d *SlowQueryDriver) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	drv, ok := d.Driver.(interface {
		ExecContext(context.Context, string, ...any) (sql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	start := time.Now()
	result, err := drv.ExecContext(ctx, query, args...)
	duration := time.Since(start)
	if d.logLevel == "debug" {
		klog.CtxInfof(ctx, "driver.ExecContext: query=%v args=%v time=%v", query, args, duration)
	} else if duration > consts.SlowQueryThreshold {
		d.Log(ctx, fmt.Sprintf("driver.ExecContext: query=%v args=%v time=%v", query, args, duration))
	}
	return result, err
}

// Query logs its params and calls the underlying driver Query method.
func (d *SlowQueryDriver) Query(ctx context.Context, query string, args, v any) error {
	start := time.Now()
	err := d.Driver.Query(ctx, query, args, v)
	duration := time.Since(start)
	if d.logLevel == "debug" {
		klog.CtxInfof(ctx, "driver.Query: query=%v args=%v time=%v err=%v", query, args, duration, err)
	} else if duration > consts.SlowQueryThreshold {
		d.Log(ctx, fmt.Sprintf("driver.Query: query=%v args=%v time=%v err=%v", query, args, duration, err))
	}
	return err
}

// QueryContext logs its params and calls the underlying driver QueryContext method if it is supported.
func (d *SlowQueryDriver) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	drv, ok := d.Driver.(interface {
		QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	start := time.Now()
	rows, err := drv.QueryContext(ctx, query, args...)
	duration := time.Since(start)
	if d.logLevel == "debug" {
		klog.CtxInfof(ctx, "driver.QueryContext: query=%v args=%v time=%v", query, args, duration)
	} else if duration > consts.SlowQueryThreshold {
		d.Log(ctx, fmt.Sprintf("driver.QueryContext: query=%v args=%v time=%v", query, args, duration))
	}
	return rows, err
}
func (d *SlowQueryDriver) Tx(ctx context.Context) (dialect.Tx, error) {
	tx, err := d.Driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &slowQueryTx{tx, d.Log, d.logLevel}, nil
}

type slowQueryTx struct {
	dialect.Tx
	Log      func(context.Context, ...any)
	logLevel string
}

func (t *slowQueryTx) Query(ctx context.Context, query string, args, v interface{}) error {
	start := time.Now()
	err := t.Tx.Query(ctx, query, args, v)
	duration := time.Since(start)

	if t.logLevel == "debug" {
		klog.CtxInfof(ctx, "driver.QueryContext: query=%v args=%v time=%v", query, args, duration)
	} else if duration > consts.SlowQueryThreshold {
		t.Log(ctx, fmt.Sprintf("driver.QueryContext: query=%v args=%v time=%v", query, args, duration))
	}
	return err
}

func (t *slowQueryTx) Exec(ctx context.Context, query string, args, v interface{}) error {
	start := time.Now()
	err := t.Tx.Exec(ctx, query, args, v)
	duration := time.Since(start)

	if t.logLevel == "debug" {
		klog.CtxInfof(ctx, "driver.Exec: query=%v args=%v time=%v", query, args, duration)
	} else if duration > consts.SlowQueryThreshold {
		t.Log(ctx, fmt.Sprintf("driver.Exec: query=%v args=%v time=%v", query, args, duration))
	}
	return err
}
