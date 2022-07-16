package pool

import "context"

type opt struct {
	Jobs    int32
	Logging bool
	Context context.Context
}

type Option func(o *opt)

func WithSetJobs(n int32) Option {
	return func(o *opt) {
		o.Jobs = n
	}
}

func WithLogging() Option {
	return func(o *opt) {
		o.Logging = true
	}
}

func WithContext(ctx context.Context) Option {
	return func(o *opt) {
		o.Context = ctx
	}
}
