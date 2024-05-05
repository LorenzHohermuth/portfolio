package mdparser

import (
	"context"
	"strings"

	"github.com/lorenzhohermuth/portfolio/pkg/github"
	"github.com/lorenzhohermuth/portfolio/view/component"
)

func GetWork() []component.Event {
	projectsFile, ghErr := github.FetchGHFile(context.Background(), "/interactive/work.md")
	if ghErr != nil {
		panic(ghErr)
	}
	components := make([]component.Event, 0)

	lines := strings.Split(projectsFile, "\n")
	tmpTitle := ""
	tmpPeriod := ""
	for _, l := range lines {
		isFilled, elm := parseMdWork(l, &tmpTitle, &tmpPeriod)
		if isFilled {
			components = append(components, elm)
		}
	}

	return components
}

func parseMdWork(line string, title *string, periode *string) (bool, component.Event){
	char := getChar(line, 0)	

	if char == "" {
		elm := component.Event{
			TimePeriode: *periode,
			Title: *title,
		}
		*title = ""
		*periode = ""
		return true, elm
		
	} else if char == "#" {
		*title = strings.TrimSpace(line[1:])
	} else {
		*periode = line
	}
	return false, component.Event{}
}
