package db

import (
	"common/consts"
	"context"
	"time"
	"user/biz/dal/db/ent"

	"github.com/cloudwego/kitex/pkg/klog"
)

// SlowQueryHook
func SlowQueryHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			start := time.Now()
			v, err := next.Mutate(ctx, m)
			duration := time.Since(start)

			defer func() {
				klog.CtxInfof(ctx, "Op=%s Type=%s Time=%s ConcreteType=%T", m.Op(), m.Type(), duration, m)
			}()

			if duration > consts.SlowQueryThreshold {
				klog.CtxInfof(ctx, "ent: [SLOW QUERY] Op=%s Type=%s Duration=%s Error=%v", m.Op(), m.Type(), duration, err)
			}
			return v, err
		})
	}
}

// SlowQueryInterceptor
func SlowQueryInterceptor() ent.Interceptor {

	return ent.InterceptFunc(func(next ent.Querier) ent.Querier {
		return ent.QuerierFunc(func(ctx context.Context, q ent.Query) (ent.Value, error) {

			start := time.Now()
			v, err := next.Query(ctx, q)
			duration := time.Since(start)
			defer func() {
				klog.CtxInfof(ctx, "Op=Query Time=%s ConcreteType=%T", duration, q)
			}()

			if duration > consts.SlowQueryThreshold {
				klog.CtxInfof(ctx, "ent: [SLOW QUERY] Op=Query Duration=%s ConcreteType=%T Error=%v", duration, q, err)
			}

			return v, err
		})
	})

}
