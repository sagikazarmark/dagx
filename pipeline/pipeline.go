package pipeline

import (
	"context"

	"github.com/sourcegraph/conc/pool"
)

// Pool is a low-level construct running functions in parallel.
//
// It's up to the implementation to decide when to interrupt execution (ie. on shared context cancellation)
// and what errors to return.
//
// See [pool.ContextPool] for default implementation and available options.
type Pool interface {
	Go(f func(ctx context.Context) error)
	Wait() error
}

// New returns a new [Pipeline] accepting steps to be executed concurrently.
func New(ctx context.Context) *Pipeline {
	return &Pipeline{
		pool: pool.New().WithErrors().WithContext(ctx),
	}
}

// NewWithPool returns a new [Pipeline] using a custom [Pool] implementation.
//
// See [pool.ContextPool] for default implementation and available options.
func NewWithPool(pool Pool) *Pipeline {
	return &Pipeline{
		pool: pool,
	}
}

// Pipeline accepts steps and runs them concurrently.
type Pipeline struct {
	pool Pool
}

// Run executes the [Pipeline].
func (p *Pipeline) Run() error {
	return p.pool.Wait()
}

// Run executes the [Pipeline].
func Run(p *Pipeline) error {
	return p.pool.Wait()
}

// AddStep adds a new step to the [Pipeline].
func (p *Pipeline) AddStep(step func(ctx context.Context) error) {
	p.pool.Go(step)
}

// AddStep adds a new step to the [Pipeline].
func AddStep(p *Pipeline, step func(ctx context.Context) error) {
	p.AddStep(step)
}

// Syncable is a generic type covering all types in the Dagger SDK implementing the Sync function.
type Syncable[T any] interface {
	Sync(ctx context.Context) (T, error)
}

// AddSyncStep adds a step to the [Pipeline] calling [Syncable.Sync] on the passed object.
func AddSyncStep[T Syncable[T]](p *Pipeline, step T) {
	p.pool.Go(func(ctx context.Context) error {
		_, err := step.Sync(ctx)

		return err
	})
}

// AddSyncStep adds a list of steps to the [Pipeline] calling [Syncable.Sync] on the passed objects.
func AddSyncSteps[T Syncable[T]](p *Pipeline, steps ...T) {
	for _, step := range steps {
		AddSyncStep(p, step)
	}
}
