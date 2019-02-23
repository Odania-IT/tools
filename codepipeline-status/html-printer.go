package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
)

type LayoutData struct {
	PageTitle string
	StartDate string
	MainContent template.HTML
}

func renderPipeline(pipeline *PipelineState, buffer *bytes.Buffer) {
	fmt.Println("Rendering Pipeline State for", pipeline.Name)
	tmpl := template.Must(template.ParseFiles("templates/pipeline.html"))
	_ = tmpl.Execute(buffer, pipeline)
}

func generateHtmlState(config PipelineStatusConfig, pipelines []*PipelineState, startDate string) string {
	fmt.Println("Writing HTML State to", config.TargetPath)
	_ = os.MkdirAll(config.TargetPath, os.ModePerm)

	// Pipelines
	var pipelineBuffer bytes.Buffer
	pipelineBuffer.Write([]byte("<div class=\"row\">"))
	for idx, pipeline := range pipelines {
		if idx % 3 == 0 && idx > 2 {
			pipelineBuffer.Write([]byte("</div><div class=\"row\">"))
		}

		renderPipeline(pipeline, &pipelineBuffer)
	}
	pipelineBuffer.Write([]byte("</div>"))

	// Main Layout
	var layoutBuffer bytes.Buffer
	tmpl := template.Must(template.ParseFiles("templates/layout.html"))

	data := &LayoutData{
		PageTitle: config.Html.PageTitle,
		StartDate: startDate,
		MainContent: template.HTML(pipelineBuffer.String()),
	}
	_ = tmpl.Execute(&layoutBuffer, data)

	// Write file
	filePath := path.Join(config.TargetPath, "index.html")

	err := ioutil.WriteFile(filePath, layoutBuffer.Bytes(), 0644)

	if err != nil {
		fmt.Println("Error writing html output", err)
		os.Exit(1)
	}

	return filePath
}
