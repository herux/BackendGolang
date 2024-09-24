package jobs

import (
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/herux/indegooweather/service"
)

func FetchWeatherJob(scheduler gocron.Scheduler, apiKey string) (gocron.Job, error) {
	weatherJobDone := make(chan bool, 1)
	weatherJobDone <- true
	job, err := scheduler.NewJob(
		gocron.DurationJob(time.Duration(60)*time.Second),
		gocron.NewTask(
			func() {
				select {
				case <-weatherJobDone:
					log.Println("Running jobs: ", FetchWeatherJob)
					runFW(weatherJobDone, apiKey)
					log.Println("Finished jobs: ", FetchWeatherJob)
				}
			},
		),
	)

	if err != nil {
		return nil, err
	}

	return job, nil
}

func runFW(weatherJobDone chan<- bool, apiKey string) {
	service.FetchWeather(apiKey)
	weatherJobDone <- true
	return
}
