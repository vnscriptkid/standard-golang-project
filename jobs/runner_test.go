package jobs_test

import (
	"context"
	"testing"

	"github.com/matryer/is"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"

	"canvas/integrationtest"
	"canvas/jobs"
	"canvas/model"
)

type testRegistry map[string]jobs.Func

func (r testRegistry) Register(name string, fn jobs.Func) {
	r[name] = fn
}

func TestRunner_Start(t *testing.T) {
	integrationtest.SkipIfShort(t)

	t.Run("starts the runner and runs jobs until the context is cancelled", func(t *testing.T) {
		is := is.New(t)

		queue, cleanup := integrationtest.CreateQueue()
		defer cleanup()

		log, logs := newLogger()

		runner := jobs.NewRunner(jobs.NewRunnerOptions{
			Log:   log,
			Queue: queue,
		})

		ctx, cancel := context.WithCancel(context.Background())

		runner.Register("test", func(ctx context.Context, m model.Message) error {
			foo, ok := m["foo"]
			is.True(ok)
			is.Equal("bar", foo)

			cancel()
			return nil
		})

		err := queue.Send(context.Background(), model.Message{"job": "test", "foo": "bar"})
		is.NoErr(err)

		// This blocks until the context is cancelled by the job function
		runner.Start(ctx)

		is.Equal(3, logs.Len())
		is.Equal("Starting", logs.All()[0].Message)
		is.Equal("Successfully ran job", logs.All()[1].Message)
		is.Equal("Stopping", logs.All()[2].Message)
	})

	t.Run("receive does not return an error if the context is already cancelled", func(t *testing.T) {
		is := is.New(t)

		queue, cleanup := integrationtest.CreateQueue()
		defer cleanup()

		// Send first, to get the queue URL when the context is not cancelled
		err := queue.Send(context.Background(), model.Message{})

		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		m, _, err := queue.Receive(ctx)
		is.NoErr(err)
		is.Equal(nil, m)
	})
}

func newLogger() (*zap.Logger, *observer.ObservedLogs) {
	core, logs := observer.New(zapcore.InfoLevel)
	return zap.New(core), logs
}
