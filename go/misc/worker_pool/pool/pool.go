package pool

import (
	"fmt"
	"sync/atomic"

	log "github.com/sirupsen/logrus"
)

type Task func() error

type Pool struct {
	tasksCh      chan Task
	workersCount uint32
	isStopped    atomic.Uint32
	isStarted    atomic.Uint32
}

type Config struct {
	WorkersCount uint32
}

func (c *Config) Validate() error {
	if c.WorkersCount == 0 {
		return fmt.Errorf("workers count equal zero")
	}
	return nil
}

func New(config Config) (*Pool, error) {
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	return &Pool{
		tasksCh:      make(chan Task),
		workersCount: config.WorkersCount,
	}, nil
}

func (p *Pool) Push(task Task) {
	if p.isStopped.Load() == 1 {
		return
	}
	p.tasksCh <- task
}

func (p *Pool) Start() {
	if !p.isStarted.CompareAndSwap(0, 1) {
		return
	}
	for idx := 0; idx < int(p.workersCount); idx++ {
		go p.startWorker()
	}
}

func (p *Pool) startWorker() {
	for task := range p.tasksCh {
		if err := task(); err != nil {
			log.Errorf("worker pool: task error: %v", err)
		}
	}
}

func (p *Pool) Stop() {
	p.isStopped.CompareAndSwap(0, 1)
}
