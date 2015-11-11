package main

import (
	"fmt"
	"testing"
	"encoding/json"
	"io/ioutil"
	//"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"github.com/hoisie/mustache"
)

type BuildFile struct{
	Language string `yaml:"language"`
	BeforeInstall []string `yaml:"before_install"`
	Install []string `yaml:"install"`
	BeforeScript []string `yaml:"before_script"`
	Script []string `yaml:"script"`
	AfterScript []string `yaml:"after_script"`
	AfterSuccess []string `yaml:"after_success"`
	AfterFailure []string `yaml:"after_failure"`
	BeforeDeploy []string `yaml:"before_deploy"`
	Deploy []string `yaml:"deploy"`
	AfterDeploy []string `yaml:"after_deploy"`
}

type GithubRepository struct {
	Url string `json:"url"`
}

type GithubPushEvent struct {
	Repository GithubRepository `json:"repository"`
}

func TestParsingGithubPayload(t *testing.T) {
	var event GithubPushEvent

	data, err := ioutil.ReadFile("./sample-github-webhook.payload")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &event)
	fmt.Println(fmt.Sprintf("%v", event))
}

func TestParsingTravisFile(t *testing.T){
	var buildFile BuildFile

	data, err := ioutil.ReadFile("./example.travis.yml")
	if err != nil {
		panic(err)
	}

	yaml.Unmarshal(data, &buildFile)
	fmt.Println(fmt.Sprintf("%v", buildFile))
}

func TestRenderingABuildFile(t *testing.T){
	var buildFile BuildFile

	travisData, err := ioutil.ReadFile("./example.travis.yml")
	if err != nil {
		panic(err)
	}

	yaml.Unmarshal(travisData, &buildFile)

	fmt.Println(fmt.Sprintf("%v", buildFile))

	buildTemplate, err := ioutil.ReadFile("./template-build.sh")
	if err != nil {
		panic(err)
	}

	data := mustache.Render(string(buildTemplate), buildFile)

	println(data)
}
