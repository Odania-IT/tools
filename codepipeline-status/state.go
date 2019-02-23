package main

import (
	"github.com/aws/aws-sdk-go/service/codepipeline"
	"time"
)

type PipelineStageState struct {
	Name string
	State string
	HtmlState string
	LastChanged string
}

type PipelineState struct {
	Name string
	Stages []PipelineStageState
}

func AppendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func allStageNames(pipelineStates []*PipelineState) []string {
	var stages []string

	for _, pipelineState := range pipelineStates {
		for _, stage := range pipelineState.Stages {
			stages = AppendIfMissing(stages, stage.Name)
		}
	}

	return stages
}

func stageByName(name string, pipelineStates *PipelineState) PipelineStageState {
	for _, stage := range pipelineStates.Stages {
		if name == stage.Name {
			return stage
		}
	}

	return PipelineStageState{}
}

func getLastStateChange(stageState *codepipeline.StageState, timeFormat string) string {
	var result *time.Time

	for _, action := range stageState.ActionStates {
		if action.LatestExecution.LastStatusChange == nil {
			continue
		}

		if result != nil  && result.After(*action.LatestExecution.LastStatusChange) {
			continue
		}

		result = action.LatestExecution.LastStatusChange
	}

	if result == nil {
		now := time.Now()
		now.Add(-200000)

		return now.Format(timeFormat)
	}

	return result.Format(timeFormat)
}

func getHtmlStateFor(state string) string {
	if state == "Succeeded" {
		return "success"
	}

	if state == "InProgress" {
		return "dark"
	}

	if state == "Failed" {
		return "danger"
	}

	return "warning"
}
