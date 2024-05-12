package main

import (
	"main/pool"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infof("application started")

	p, err := pool.New(pool.Config{
		WorkersCount: 5,
	})
	if err != nil {
		log.Fatalf("failed to create new worker pool: %v", err)
	}
	log.Infof("new worker pool created")

	log.Infof("task handling started")
	p.Start()

	for idx := 1; idx <= 5; idx++ {
		idx := idx
		log.Infof("pushed task %d to worker pool", idx)

		p.Push(func() error {
			log.Infof("task %d completed", idx)
			return nil
		})
	}
	time.Sleep(1 * time.Second)

	p.Stop()
	log.Infof("worker pool stopped")

	log.Infof("try pushing task to worker pool after stopping")

	p.Push(func() error {
		log.Infof("task completed after worker pool stopping")
		return nil
	})

	time.Sleep(1 * time.Second)

	log.Infof("application finished")
}
