package providers

import (
	"github.com/stretchr/testify/assert"
	//"gopkg.in/jarcoal/httpmock.v1"
	//"net/url"
	"testing"
)

type Configuration struct {
	MarathonURL     string
	BeehiveToken    string
	JenkinsURL      string
	JenkinsUsername string
	JenkinsPassword string
	Interval        string
}

// Need to move things to interfaces first
/*
func TestAssembleText(t *testing.T) {

	service := Service{
		ID:           "/testservice",
		Name:         "testservice",
		Instances:    []string{"instance1", "instance2"},
		TasksRunning: 2,
		Health:       true,
	}
	//assert.Equal(t, assembleText(service), "Name: testservice\nURL: Number of tasks: 2\nInstances:\n instance1\ninstance2\nHealthy: true\n", "assembleText should assemble text correctly")
}
*/

func TestConvertMarathonID(t *testing.T) {
	assert.Equal(t, convertMarathonID("/group1/service1"), "group1-service1", "convertMarathonID should convert / to - except first character")
	assert.Equal(t, convertMarathonID("group1/service1"), "group1-service1", "convertMarathonID should convert ids without / at beginning")
	assert.Equal(t, convertMarathonID("group1/service1/service2/service3"), "group1-service1-service2-service3", "convertMarathonID should convert really long app id's")
	assert.Equal(t, convertMarathonID("service1"), "service1", "convertMarathonID should convert single ids without groups")
}

func TestPathType(t *testing.T) {
	assert.Equal(t, pathType("/service1"), "apps", "pathType with less than 2 / should return apps")
	assert.Equal(t, pathType("/group1/service1"), "group", "pathType with more than 1 / should return group")
}

func TestGenerateMarathonURL(t *testing.T) {
	assert.Equal(t, generateMarathonURL("/group1/service1", "https://marathon.domain.com"), "https://marathon.domain.com/ui/#/apps/%2Fgroup1%2Fservice1", "generateMarathonURL should create correct path for apps with groups")
	assert.Equal(t, generateMarathonURL("/service1", "https://marathon.domain.com"), "https://marathon.domain.com/ui/#/apps/%2Fservice1", "generateMarathonURL should create correct path for apps without groups")
	assert.Equal(t, generateMarathonURL("/group1/group2/service1", "https://marathon.domain.com"), "https://marathon.domain.com/ui/#/apps/%2Fgroup1%2Fgroup2%2Fservice1", "generateMarathonURL should create correct path for apps with 2 levels of groups")

}
