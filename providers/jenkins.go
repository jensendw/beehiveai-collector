package providers

import (
	"fmt"
	"github.com/jensendw/beehiveai-go"
	"github.com/yosida95/golang-jenkins"
)

// JenkinsJob - The information for the jenkins job
type JenkinsJob struct {
	Name      string
	URL       string
	LastBuild string
}

//JenkinsCollector - kicks off the process to collect data from jenkins
func JenkinsCollector() {
	auth := setCredentials(Config.JenkinsUsername, Config.JenkinsPassword)

	jenkins := gojenkins.NewJenkins(auth, Config.JenkinsURL)

	var jenkinsJobs []JenkinsJob

	jobs, err := jenkins.GetJobs()
	if err != nil {
		fmt.Println("Got an error: ", err)
	}
	for _, job := range jobs {
		jenkinsJob := JenkinsJob{}
		jenkinsJob.Name = job.Name
		jenkinsJob.URL = job.Url
		lastBuild, err := jenkins.GetLastBuild(job)
		if err != nil {
			Logger.Errorf("Unable to get last build information for jenkins job %v", job.Name)
		}
		jenkinsJob.LastBuild = lastBuild.Url
		jenkinsJobs = append(jenkinsJobs, jenkinsJob)
	}
	sendToBeehive(jenkinsJobs)
}

//Sets the credentials if any were configured
func setCredentials(jenkinsUsername string, jenkinsPassword string) *gojenkins.Auth {
	if (jenkinsUsername != "") || (jenkinsPassword != "") {
		auth := &gojenkins.Auth{
			Username: jenkinsUsername,
			ApiToken: jenkinsPassword,
		}
		return auth
	}
	var auth *gojenkins.Auth
	return auth
}

func sendToBeehive(jenkinsJobs []JenkinsJob) {
	bclient := bhive.NewClient(Config.BeehiveToken)

	Logger.Infof("Sending %v Jenkins records to BeehiveAI", len(jenkinsJobs))
	for _, job := range jenkinsJobs {
		descriptiveName := job.Name + " jenkins job"
		_, err := bclient.CreateIntegration(descriptiveName, assembleJenkinsText(job), job.Name)
		if err != nil {
			Logger.Errorf("Error creating BeehiveAI integration: %s", err)
		}
	}

}

func assembleJenkinsText(jenkinsJob JenkinsJob) string {
	// Just cleaner to declare things up front for debug purposes
	return fmt.Sprintf("<%v:%v>\nLastBuild: %v\n", jenkinsJob.URL, jenkinsJob.Name, jenkinsJob.LastBuild)
}
