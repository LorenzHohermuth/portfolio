package mdparser

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/lorenzhohermuth/portfolio/pkg/github"
	"github.com/lorenzhohermuth/portfolio/view/component"
)

const pathProjects string = "projects.md"
const pathWork string = "work.md"

func GetProjects() []component.CarouselEntry {
	github.FetchGHFile(context.Background(), "/cmd")
	components := make([]component.CarouselEntry, 0)

	dat, err := os.ReadFile("interactive/projects.md")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	tmpTitle := ""
	tmpText := ""
	tmpImg := ""
	for _, l := range lines {
		isFilled, elm := parseMd(l, &tmpTitle, &tmpText, &tmpImg)
		if isFilled {
			components = append(components, elm)
		}
	}

	fmt.Println(components)
	return components
}

func getChar(text string, pos int) string {
	text = strings.TrimSpace(text)
	if text == "" {
		return ""
	}
	return string([]rune(text)[pos])
}

func parseMd(line string, title *string, text *string, img *string) (bool, component.CarouselEntry){
	char := getChar(line, 0)	

	if char == "" {
		elm := component.CarouselEntry{
			ImgPath: *img, 
			Title: *title,
			Text: *text,
		}
		*img = ""
		*title = ""
		*text = ""
		return true, elm
		
	} else if char == "#" {
		*title = strings.TrimSpace(line[1:])
	
	} else if char == "!" {
		rgx := regexp.MustCompile(`\((.*?)\)`)
		*img = rgx.FindStringSubmatch(line)[1]

	} else {
		*text += line
	}
	return false, component.CarouselEntry{}
}
