package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/herux/indegooweather/config"
	"github.com/herux/indegooweather/constant"
	"github.com/herux/indegooweather/db"
	"github.com/herux/indegooweather/route"
	"github.com/herux/indegooweather/server"
	flag "github.com/spf13/pflag"
)

func main() {
	path := getConfigPath()
	_ = config.Load(path)
	db.Init(false)

	srv := server.SetupService(config.Service(), route.RegisterAPI)
	srv.Run()
}

func getConfigPath() string {
	f := flag.NewFlagSet("indegooweather-apis", flag.ExitOnError)
	f.Usage = func() {
		fmt.Println(getFlagUsage(f))
		os.Exit(0)
	}
	config := f.String("config", constant.DefaultConfigFile, "configuration file path")
	f.Parse(os.Args[1:])

	return *config
}

func getFlagUsage(f *flag.FlagSet) string {
	usage := "indegooweather-apis\n\n"
	usage += "Options:\n"

	options := strings.ReplaceAll(f.FlagUsages(), "    ", "  ")
	usage += fmt.Sprintf("%s\n", options)

	return usage
}
