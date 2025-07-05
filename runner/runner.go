package runner

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/riandyhasan/imperio/config"
)

// Executor defines the interface for executing operations.
//
//go:generate mockery --name=Executor --output=./mocks
type Executor interface {
	Execute(op string) error
}

// Runner manages the simulation lifecycle.
type Runner struct {
	Config   *config.Config
	Executor Executor
	Duration time.Duration // negative means "run forever"
}

// NewRunner constructs a new Runner with injected dependencies.
func NewRunner(cfg *config.Config, exec Executor, duration time.Duration) *Runner {
	return &Runner{
		Config:   cfg,
		Executor: exec,
		Duration: duration, // if zero => infinite mode
	}
}

// Start launches the simulation using the provided configuration.
func (r *Runner) Start() error {
	opsChan := make(chan string)
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Launch operation generator
	go r.generateOperations(ctx, opsChan)

	// Launch workers
	for i := range r.Config.Concurrency {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for op := range opsChan {
				if err := r.Executor.Execute(op); err != nil {
					log.Printf("[Worker %d] Failed to execute '%s': %v", workerID, op, err)
				}
			}
		}(i + 1)
	}

	// Decide when to stop
	if r.Duration > 0 {
		time.Sleep(r.Duration)
		cancel()                    // cancel context to stop generator
		time.Sleep(1 * time.Second) // give workers time to drain queue
		close(opsChan)              // stop all workers
	}

	wg.Wait()
	return nil
}

// generateOperations emits operations at the configured rate, until canceled.
func (r *Runner) generateOperations(ctx context.Context, out chan<- string) {
	ticker := time.NewTicker(time.Second / time.Duration(r.Config.OpsPerSecond))
	defer ticker.Stop()

	opIndex := 0
	ops := r.Config.Operations

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			out <- ops[opIndex]
			opIndex = (opIndex + 1) % len(ops)
		}
	}
}
