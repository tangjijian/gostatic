package goroutinePool

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Id      int
	randNum int
}

type Result struct {
	job *Job
	sum int
}

func Init() {

	jobChan := make(chan *Job, 128)

	resultChan := make(chan *Result, 128)
	createPool(64, jobChan, resultChan)
	go func(resultChan chan *Result) {
		for data := range resultChan {

			fmt.Printf("Job.id=%d.randnum=%d,res=%d\n", data.job.Id, data.job.randNum, data.sum)

		}
	}(resultChan)

	var id int
	for {
		id++

		randNum := rand.Int() // 随机数
		job := &Job{Id: id, randNum: randNum}
		jobChan <- job
	}
}
func createPool(num int, jobChan chan *Job, resChan chan *Result) {

	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resChan chan *Result) {
			for job := range jobChan {

				var sum int
				r_num := job.randNum
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				res := &Result{job: job, sum: sum}
				resChan <- res
			}

		}(jobChan, resChan)
	}
}
