package main

import (
	"os"
	"path/filepath"
)

type HtmlConfig struct {
	PageTitle string
}

type PipelineStatusConfig struct {
	Region string
	TargetPath string
	Html HtmlConfig
	BucketName string
	BucketKey string // key inside the s3
	TimeFormat string
}

func newPipelineStatusConfig() *PipelineStatusConfig {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	return &PipelineStatusConfig{
		Region: "eu-central-1",
		TargetPath: filepath.Join(dir, "target"),
		Html: HtmlConfig{
			PageTitle: "Codepipeline Status",
		},
		TimeFormat: "2006-01-02 15:04",
	}
}
