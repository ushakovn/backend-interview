package pool_v2

import (
  "sync"
  "sync/atomic"
)

// Task задача на выполнение
type Task func() error

// WorkerPool примитив конкурентного выполнения задач
type WorkerPool struct {
  count uint32

  tasksCh chan Task
  errCh   chan error

  isStarted atomic.Int32
  isStopped atomic.Int32
}

// New создает примитив WorkerPool
func New(count uint32) *WorkerPool {
  return &WorkerPool{
    count: count,

    tasksCh: make(chan Task),
    errCh:   make(chan error),
  }
}

// Push добавляет задачу на выполнение
func (p *WorkerPool) Push(task Task) {
  if p.isStarted.Load() == 0 {
    return
  }
  if p.isStopped.Load() == 1 {
    return
  }
  p.tasksCh <- task
}

// Start запускает выполнение задач
func (p *WorkerPool) Start() {
  if !p.isStarted.CompareAndSwap(0, 1) {
    return
  }

  wg := sync.WaitGroup{}
  wg.Add(int(p.count))

  for idx := 0; idx < int(p.count); idx++ {
    go func() {
      defer wg.Done()

      for task := range p.tasksCh {
        p.errCh <- task()
      }
    }()
  }

  go func() {
    wg.Wait()
    close(p.errCh)
  }()
}

// Stop останавливает выполнение новых задач
func (p *WorkerPool) Stop() {
  if p.isStarted.Load() == 0 {
    return
  }
  if !p.isStopped.CompareAndSwap(0, 1) {
    return
  }
  close(p.tasksCh)
}

// Errors возвращает канал с ошибками при выполнении задач
func (p *WorkerPool) Errors() <-chan error {
  return p.errCh
}
