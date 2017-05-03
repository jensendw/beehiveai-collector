package providers

import (
	"github.com/stretchr/testify/assert"
	//"gopkg.in/jarcoal/httpmock.v1"
	"fmt"
	"github.com/yosida95/golang-jenkins"
	"testing"
)

func TestSetCredentials(t *testing.T) {

	expectedAuth := &gojenkins.Auth{
		Username: "jenkinsuser",
		ApiToken: "jenkinspassword",
	}

	var emptyAuth *gojenkins.Auth

	assert.Equal(t, setCredentials("jenkinsuser", "jenkinspassword"), expectedAuth, "setCredentials should return struct when getting credentials")
	assert.Equal(t, setCredentials("", ""), emptyAuth, "setCredentials should return struct when getting credentials")

}

func TestAssembleJenkinsText(t *testing.T) {

	jenkinsJob := JenkinsJob{
		Name:      "testjenkinsjob",
		URL:       "https://my.awesomejenkins.com/job/testjenkinsjob",
		LastBuild: "https://my.awesomejenkins.com/job/testjenkinsjob/latestbuild",
	}

	assert.Equal(t, assembleJenkinsText(jenkinsJob), fmt.Sprintf("<%v:%v>\nLastBuild: %v\n", jenkinsJob.URL, jenkinsJob.Name, jenkinsJob.LastBuild), "assembleJenkinsText should return formateed text")

}

/*
func TestSendToBeehive(t *testing.T) {
	jenkinsJobs := []JenkinsJob
	jenkinsJob := JenkinsJob{
		Name:      "testjenkinsjob",
		URL:       "https://my.awesomejenkins.com/job/testjenkinsjob",
		LastBuild: "https://my.awesomejenkins.com/job/testjenkinsjob/latestbuild",
	}
	assert.Equal(t, sendToBeehive([]jenkinsJob), fmt.Sprintf("<%v:%v>\nLastBuild: %v\n", jenkinsJob.URL, jenkinsJob.Name, jenkinsJob.LastBuild), "assembleJenkinsText should return formateed text")

}
*/
