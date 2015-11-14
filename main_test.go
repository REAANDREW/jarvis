package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
	//"github.com/stretchr/testify/assert"
	"github.com/hoisie/mustache"
	"gopkg.in/yaml.v2"
)

type BuildFile struct {
	Language      string   `yaml:"language"`
	BeforeInstall []string `yaml:"before_install"`
	Install       []string `yaml:"install"`
	BeforeScript  []string `yaml:"before_script"`
	Script        []string `yaml:"script"`
	AfterScript   []string `yaml:"after_script"`
	AfterSuccess  []string `yaml:"after_success"`
	AfterFailure  []string `yaml:"after_failure"`
	BeforeDeploy  []string `yaml:"before_deploy"`
	Deploy        []string `yaml:"deploy"`
	AfterDeploy   []string `yaml:"after_deploy"`
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

func TestParsingTravisFile(t *testing.T) {
	var buildFile BuildFile

	data, err := ioutil.ReadFile("./example.travis.yml")
	if err != nil {
		panic(err)
	}

	yaml.Unmarshal(data, &buildFile)
	fmt.Println(fmt.Sprintf("%v", buildFile))
}

func TestRenderingABuildFile(t *testing.T) {
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

func TestCreatingTempDirectoryStructure(t *testing.T) {
	fmt.Println(os.TempDir())

	rootPath := path.Join(os.TempDir(),"jarvis")

	os.MkdirAll(rootPath,0777)

	tempDir, err := ioutil.TempDir(rootPath, "")
	if err != nil {
		panic(err)
	}

	outputPath := path.Join(tempDir, "/output")
	buildPath := path.Join(tempDir, "/build")
	os.MkdirAll(outputPath,0777)
	os.MkdirAll(buildPath,0777)

	fmt.Println(outputPath)
	fmt.Println(buildPath)
	fmt.Println(tempDir)
}

func TestInitiateABuild(t *testing.T){
	testServer := CreateRequestRecordingServer(TestPort)
	testServer.Start()



	testServer.Stop()
}
