package grpool

import (
	"fmt"
	"github.com/ivpusic/grpool"
	"testing"
	"time"
)

func TestExample(t *testing.T) {
	// number of workers, and size of job queue
	pool := grpool.NewPool(0, 1)

	// release resources used by pool
	defer pool.Release()

	// submit one or more jobs to pool
	for i := 0; i < 10; i++ {
		count := i
		pool.JobQueue <- func() {
			fmt.Printf("I am worker! Number %d\n", count)
		}
	}

	// dummy wait until jobs are finished
	time.Sleep(10 * time.Second)
}

func TestJobFinish(t *testing.T) {
	// number of workers, and size of job queue
	pool := grpool.NewPool(1, 1)
	defer pool.Release()

	// how many jobs we should wait
	pool.WaitCount(10)

	// submit one or more jobs to pool
	for i := 0; i < 10; i++ {
		count := i

		pool.JobQueue <- func() {
			// say that job is done, so we can know how many jobs are finished
			defer pool.JobDone()

			fmt.Printf("hello %d\n", count)
		}
	}

	// wait until we call JobDone for all jobs
	pool.WaitAll()
}
