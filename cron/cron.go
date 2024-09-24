package cron

import (
	"fmt"

	"github.com/go-co-op/gocron/v2"
)

type Cron struct {
	schedulers []gocron.Scheduler
}

func NewCron() *Cron {
	return &Cron{
		schedulers: make([]gocron.Scheduler, 0),
	}
}

func (c *Cron) NewScheduler(NewScheduler gocron.Scheduler) {
	c.schedulers = append(c.schedulers, NewScheduler)
}

func (c *Cron) Start() {
	for i, scheduler := range c.schedulers {
		scheduler.Start()
		fmt.Println("scheduler ", i+1, " started.")
	}
}
