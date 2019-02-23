package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codepipeline"
	"github.com/aws/aws-sdk-go/service/sts"
	"os"
	"time"
)

func codePipelines(sess *session.Session) []*PipelineState {
	client := codepipeline.New(sess)
	ctx := context.Background()
	var result []*PipelineState

	paginator := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			req, _ := client.ListPipelinesRequest(&codepipeline.ListPipelinesInput{})
			req.SetContext(ctx)

			return req, req.Error
		},
	}

	for paginator.Next() {
		page := paginator.Page().(*codepipeline.ListPipelinesOutput)
		fmt.Println("Received", len(page.Pipelines), "objects in page")

		for _, obj := range page.Pipelines {
			fmt.Println("Name:", aws.StringValue(obj.Name))
			result = append(result, &PipelineState{
				Name: aws.StringValue(obj.Name),
			})
		}
	}

	return result
}

func codePipelineState(sess *session.Session, pipelines []*PipelineState, timeFormat string) []*PipelineState {
	client := codepipeline.New(sess)

	for _, pipeline := range pipelines {
		fmt.Println("Looking info for:", pipeline.Name)
		pipelineState, err := client.GetPipelineState(&codepipeline.GetPipelineStateInput{
			Name: &pipeline.Name,
		})

		if err != nil {
			fmt.Println("Error retrieving pipelineState state: ", err)
			os.Exit(1)
		}

		fmt.Println("Updates", pipelineState.Updated)

		for _, stageState := range pipelineState.StageStates {
			pipeline.Stages = append(pipeline.Stages, PipelineStageState{
				Name:        *stageState.StageName,
				State:       *stageState.LatestExecution.Status,
				HtmlState:   getHtmlStateFor(*stageState.LatestExecution.Status),
				LastChanged: getLastStateChange(stageState, timeFormat),
			})
		}
	}

	return pipelines
}

func main() {
	startDate := time.Now()
	arguments := parseArguments()
	config := newPipelineStatusConfig(arguments.config)

	sess, err := session.NewSession(&aws.Config{Region: aws.String(config.Region)})

	if err != nil {
		fmt.Println("Error creating aws session: ", err)
		os.Exit(1)
	}

	stsClient := sts.New(sess)
	identity, err := stsClient.GetCallerIdentity(&sts.GetCallerIdentityInput{})

	if err != nil {
		fmt.Println("Error retrieving caller identity: ", err)
		os.Exit(1)
	}

	fmt.Println("AWS Identity", identity)

	pipelines := codePipelines(sess)
	pipelines = codePipelineState(sess, pipelines, config.TimeFormat)
	printResult(pipelines, startDate.Format(config.TimeFormat))

	// Write HTML State File
	htmlStateFile := generateHtmlState(config, pipelines, startDate.Format("15:04"))
	if config.BucketName != "" && config.BucketKey != "" {
		err = AddToS3(sess, htmlStateFile, config)

		if err != nil {
			fmt.Println("Error uploading html state to s3: ", err)
			os.Exit(1)
		}
	}
}
