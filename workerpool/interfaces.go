package pool

import "context"

type Worker interface {
	Process(job Job)
	Start()
	Shutdown()
}

type Job interface {
	Process(ctx context.Context)
	Id() string
}
