package jobs

import (
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/herux/indegooweather/service"
)

func FetchBikestationJob(scheduler gocron.Scheduler) (gocron.Job, error) {
	bikestationJobDone := make(chan bool, 1)
	bikestationJobDone <- true
	job, err := scheduler.NewJob(
		gocron.DurationJob(time.Duration(60)*time.Second),
		gocron.NewTask(
			func() {
				select {
				case <-bikestationJobDone:
					log.Println("Running jobs: ", FetchBikestationJob)
					runFB(bikestationJobDone)
					log.Println("Finished jobs: ", FetchBikestationJob)
				}
			},
		),
	)

	if err != nil {
		return nil, err
	}

	return job, nil
}

func runFB(bikestationJobDone chan<- bool) {
	service.FetchAndStoreIndegoData()
	bikestationJobDone <- true
	return
}
