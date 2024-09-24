package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-co-op/gocron/v2"
	"github.com/herux/indegooweather/config"
	"github.com/herux/indegooweather/constant"
	"github.com/herux/indegooweather/cron"
	"github.com/herux/indegooweather/db"
	"github.com/herux/indegooweather/jobs"
	flag "github.com/spf13/pflag"
)

func main() {
	path := getConfigPath()
	_ = config.Load(path)

	db.Init()

	cron := cron.NewCron()
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		log.Fatalf("error creating scheduler: %v", err)
	}

	jobs.FetchWeatherJob(scheduler, config.OpenWeatherAPIKey())
	jobs.FetchBikestationJob(scheduler)

	cron.NewScheduler(scheduler)
	cron.Start()

	select {}
}

func getConfigPath() string {
	f := flag.NewFlagSet("indegooweather-cron", flag.ExitOnError)
	f.Usage = func() {
		fmt.Println(getFlagUsage(f))
		os.Exit(0)
	}
	config := f.String("config", constant.DefaultConfigFile, "configuration file path")
	f.Parse(os.Args[1:])

	return *config
}

func getFlagUsage(f *flag.FlagSet) string {
	usage := "indegooweather-cron\n\n"
	usage += "Options:\n"

	options := strings.ReplaceAll(f.FlagUsages(), "    ", "  ")
	usage += fmt.Sprintf("%s\n", options)

	return usage
}
