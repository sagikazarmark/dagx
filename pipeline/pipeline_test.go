package pipeline_test

import (
	"context"
	"testing"

	"github.com/sagikazarmark/dagx/pipeline"
)

type syncable struct{}

func (s *syncable) Sync(_ context.Context) (*syncable, error) {
	return s, nil
}

func TestPipeline(t *testing.T) {
	p := pipeline.New(context.Background())

	pipeline.AddStep(p, func(_ context.Context) error {
		return nil
	})

	s := &syncable{}

	pipeline.AddSyncStep(p, s)

	err := p.Run()
	if err != nil {
		t.Fatal(err)
	}
}
