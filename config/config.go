package config

import (
	"github.com/jensendw/beehiveai-collector/logger"
	"github.com/kelseyhightower/envconfig"
	"os"
)

var Logger = *logger.Logger

// Configuration stores the configuration data
type Configuration struct {
	MarathonURL     string `envconfig:"marathon_url"`
	BeehiveToken    string `required:"true" envconfig:"token"`
	JenkinsURL      string `envconfig:"jenkins_url"`
	JenkinsUsername string `envconfig:"jenkins_username"`
	JenkinsPassword string `envconfig:"jenkins_password"`
	Interval        string `envconfig:"check_interval"`
}

var Config = LoadConfig().(*Configuration)

// LoadConfig simply loads the configuration
func LoadConfig() interface{} {
	var s Configuration
	err := envconfig.Process("beehiveai", &s)
	if err != nil {
		Logger.Error(err.Error())
		os.Exit(2)

	}
	return &s
}
