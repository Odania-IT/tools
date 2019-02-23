package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func buildTabStringFor(keys []string, pipeline *PipelineState) string {
	result := pipeline.Name + "\t"

	for _, key := range keys {
		result += stageByName(key, pipeline).State + "\t"
	}

	return result
}

func buildTabStringForLastChanged(keys []string, pipeline *PipelineState) string {
	result := "\t"

	for _, key := range keys {
		result += stageByName(key, pipeline).LastChanged + "\t"
	}

	return result
}

func printResult(pipelines []*PipelineState, startDate string) {
	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 0, 8, 0, '\t', 0)

	fmt.Println()
	fmt.Println()
	fmt.Println("Current States", startDate)
	fmt.Println()
	stageNames := allStageNames(pipelines)

	// Print Header
	header := "Pipeline Name\t"
	for _, key := range stageNames {
		header += key + "\t"
	}
	_, err := fmt.Fprintln(writer, header)

	if err != nil {
		fmt.Println("[printResult] Error printing header", err)
	}

	// Print Pipeline State
	for _, pipeline := range pipelines {
		line := buildTabStringFor(stageNames, pipeline)
		_, err = fmt.Fprintln(writer, line)

		if err != nil {
			fmt.Println("[printResult] Error printing pipeline result", err)
		}

		line = buildTabStringForLastChanged(stageNames, pipeline)
		_, err = fmt.Fprintln(writer, line)

		if err != nil {
			fmt.Println("[printResult] Error printing pipeline result", err)
		}
	}

	err = writer.Flush()

	if err != nil {
		fmt.Println("[printResult] Error flushing tab writer", err)
	}

	fmt.Println()
	fmt.Println()
}
