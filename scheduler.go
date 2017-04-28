package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/jensendw/beehiveai-collector/providers"
	"strconv"
)

func ScheduleRunChecks() {
	checkInterval, err := strconv.ParseUint(Config.Interval, 10, 64)
	if err != nil {
		Logger.Errorf("Couldn't convert interval configuration value to integer: %v", err)
	}

	Logger.Infof("Running checks every %v seconds", checkInterval)

	gocron.Every(checkInterval).Seconds().Do(runChecks)
	<-gocron.Start()
}

func runChecks() {
	if Config.MarathonURL != "" {
		Logger.Infof("Marathon provider enabled, using URL: %v", Config.MarathonURL)
		go providers.MarathonCollector()
	}

	if Config.JenkinsURL != "" {
		Logger.Infof("Jenkins provider enabled, using URL: %v", Config.JenkinsURL)
		go providers.JenkinsCollector()
	}
}
