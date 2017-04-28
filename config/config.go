package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

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
		fmt.Println(err.Error())
		//lshoudl really exit here
	}
	return &s
}
