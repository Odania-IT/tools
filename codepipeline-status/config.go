package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

const (
	VERSION = "0.1.1"
)

type HtmlConfig struct {
	PageTitle string `yaml:"PageTitle"`
}

type PipelineStatusConfig struct {
	Region string             `yaml:"Region"`
	TargetPath string `yaml:"TargetPath"`
	Html HtmlConfig `yaml:"Html"`
	BucketName string `yaml:"BucketName"`
	BucketKey string `yaml:"BucketKey"` // key inside the s3
	TimeFormat string `yaml:"TimeFormat"`
}

type Arguments struct {
	config string
}

func parseArguments() Arguments {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	arguments := &Arguments{
		config: path.Join(dir, "config.yml"),
	}
	flag.StringVar(&arguments.config, "config", arguments.config, "Config file")
	flag.Usage = flagUsage
	flag.Parse()

	return *arguments
}

func flagUsage() {
	_, _ = fmt.Fprintf(os.Stderr, "Version of %s: %s\n", os.Args[0], VERSION)
	_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func newPipelineStatusConfig(configPath string) PipelineStatusConfig {
	bytes, err := ioutil.ReadFile(configPath)

	if err != nil {
		fmt.Println("Could not read config", err)
		os.Exit(1)
	}

	config := &PipelineStatusConfig{}

	if err := yaml.Unmarshal(bytes, config); err != nil {
		fmt.Println("Could not parse config", err)
		os.Exit(1)
	}

	return setDefaultValues(*config)
}

func setDefaultValues(config PipelineStatusConfig) PipelineStatusConfig {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	if config.Region == "" {
		config.Region = "eu-central-1"
	}

	if config.TargetPath == "" {
		config.TargetPath = filepath.Join(dir, "target")
	}

	if config.Html.PageTitle == "" {
		config.Html.PageTitle = "Codepipeline Status"
	}

	if config.TimeFormat == "" {
		config.TimeFormat = "2006-01-02 15:04"
	}

	return config
}
