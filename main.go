package main

import (
	"github.com/jensendw/beehiveai-collector/config"
	"github.com/jensendw/beehiveai-collector/logger"
)

var Config = config.LoadConfig().(*config.Configuration)
var Logger = *logger.Logger

func main() {
	ScheduleRunChecks()
}
