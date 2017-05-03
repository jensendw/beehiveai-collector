package providers

import (
	"fmt"
	marathon "github.com/gambol99/go-marathon"
	"github.com/jensendw/beehiveai-go"
	"net/url"
	"strconv"
	"strings"
)

// Service struct to store details we care about for each service
// then instantiate into a slice for each service
type Service struct {
	ID           string
	Name         string
	Instances    []string
	TasksRunning int
	Health       bool
}

//MarathonCollector kicks off the collection process for Marathon
// its FUGLY!
func MarathonCollector() {
	marathonURL := Config.MarathonURL
	config := marathon.NewDefaultConfig()
	config.URL = marathonURL
	client, err := marathon.NewClient(config)

	if err != nil {
		Logger.Errorf("Failed to create a client for marathon, error: %s", err)
	}

	applications, err := client.Applications(nil)
	if err != nil {
		Logger.Errorf("Failed to get Marathon applications: %s", err)
	}

	var myServices []Service
	for _, application := range applications.Apps {
		service := Service{}
		service.Name = convertMarathonID(application.ID)
		service.ID = application.ID

		details, err := client.Application(application.ID)
		if err != nil {
			Logger.Errorf("error getting application details")
		}
		service.TasksRunning = details.TasksRunning

		service.Health, err = client.ApplicationOK(details.ID)
		if err != nil {
			Logger.Errorf("Error while checking if application is healthy: %s", err)
		}

		for _, task := range details.Tasks {
			service.Instances = append(service.Instances, "http://"+task.Host+":"+strconv.Itoa(task.Ports[0]))
		}
		myServices = append(myServices, service)

	}
	bclient := bhive.NewClient(Config.BeehiveToken)

	Logger.Infof("Sending %v Marathon records to BeehiveAI", len(myServices))
	for _, serviceItem := range myServices {
		_, err := bclient.CreateIntegration(serviceItem.Name, assembleText(serviceItem), serviceItem.Name)
		if err != nil {
			Logger.Errorf("Error creating BeehiveAI integration: %s", err)
		}
	}
}

func assembleText(service Service) string {
	// Just cleaner to declare things up front for debug purposes
	instances := strings.Join(service.Instances, "\n")
	tasksRunning := strconv.Itoa(service.TasksRunning)
	health := strconv.FormatBool(service.Health)
	name := service.Name
	marathonURL := generateMarathonURL(service.ID, Config.MarathonURL)
	return fmt.Sprintf("Name: %v\nLink: %v\nNumber of tasks: %v\nInstances:\n %v\nHealthy: %v\n", name, marathonURL, tasksRunning, instances, health)
}

func convertMarathonID(id string) string {
	thestring := strings.Replace(id, "/", "-", -1)
	if thestring[0:1] == "-" {
		return thestring[1:len(thestring)]
	}
	return thestring
}

//Determines if path should contain group or apps as part of URL
func pathType(path string) string {
	pathCount := strings.Count(path, "/")
	switch {
	case pathCount > 1:
		return "group"
	default:
		return "apps"
	}
}

//Creates the url to the application in marathon
func generateMarathonURL(appId string, marathonURL string) string {
	return fmt.Sprintf("%v/ui/#/apps/%s", marathonURL, url.PathEscape(appId))
}
