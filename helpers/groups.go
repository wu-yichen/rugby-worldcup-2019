package helpers

import (
	"fmt"
	"log"
	"strings"

	"github.com/followme1987/wcwhen/api/pulselive"
	"github.com/followme1987/wcwhen/model"

	"github.com/buger/goterm"
	"github.com/urfave/cli"
)

func GetTeamsByGroupName(c *cli.Context) {
	RetrieveTeamsByGroupName(c)
}

func RetrieveTeamsByGroupName(c *cli.Context) (content string) {
	gName := c.String("name")
	totals := goterm.NewTable(0, 20, 5, ' ', 0)
	if gName != "" {
		matches, err := RetrieveMatches(pulselive.New())
		if err != nil {
			log.Fatal(err)
		}

		m := getGrpMap(matches, gName)
		var contentBuilder strings.Builder
		if len(m) > 0 {
			printGroupHeader(totals)
			for key, value := range m {
				result := fmt.Sprintf("%s\t%s\t%s\n", gName, key, value)
				_, err := fmt.Fprint(totals, result)
				if err != nil {
					log.Fatal(err)
				}
				contentBuilder.WriteString(result + " ")
			}
			content = contentBuilder.String()
			PrintTable(totals)
		} else {
			content = fmt.Sprintf("Sorry, group name %s is invalid ", gName)
			_, err := goterm.Println(goterm.Color(content, goterm.RED))
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		content = "Please provide group name"
		_, err := goterm.Println(goterm.Color(content, goterm.RED))
		if err != nil {
			log.Fatal(err)
		}
	}
	goterm.Flush()
	return
}

func printGroupHeader(totals *goterm.Table) {
	_, err := fmt.Fprintf(totals, "GROUP\tFULL NAME\tSHORT NAME\n")
	if err != nil {
		log.Fatal(err)
	}
}

func getGrpMap(matches []model.Matches, gName string) map[string]string {
	tMap := make(map[string]string)
	for _, m := range matches {
		if l := len(strings.Fields(m.EventPhase)); l > 1 {
			pool := strings.Fields(m.EventPhase)[1]
			if strings.EqualFold(pool, gName) {
				for _, team := range m.Teams {
					tMap[team.Name] = team.Abbreviation
				}
			}
		}
	}
	return tMap
}
