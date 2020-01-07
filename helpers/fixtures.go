package helpers

import (
	"fmt"
	"log"
	"strings"

	"github.com/followme1987/wcwhen/api/pulselive"

	"github.com/buger/goterm"
	"github.com/followme1987/wcwhen/model"
	"github.com/urfave/cli"
)

func GetFixturesByTeam(c *cli.Context) {
	RetrieveFixturesByTeam(c)
}

func RetrieveFixturesByTeam(c *cli.Context) (content string) {
	matches, err := RetrieveMatches(pulselive.New())
	if err != nil {
		log.Fatal(err)
	}

	teamName := c.String("name")
	totals := goterm.NewTable(0, 20, 5, ' ', 0)
	if teamName != "" {
		printFixturesHeader(totals)
		var found = false
		var str strings.Builder
		for _, m := range matches {
			var team []model.Teams
			for _, t := range m.Teams {
				if strings.EqualFold(teamName, t.Name) {
					team = m.Teams
					break
				}
			}
			for _, t := range team {
				if !strings.EqualFold(teamName, t.Name) {
					result := printFixtureContent(totals, m, t)
					str.WriteString(result + " ")
					found = true
					break
				}
			}
		}
		if !found {
			content = fmt.Sprintf("Sorry, there is no game for %s ", teamName)
			_, err := goterm.Println(goterm.Color(content, goterm.RED))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			PrintTable(totals)
			content = str.String()
		}
	} else {
		content = "Please provide team name"
		_, err := goterm.Println(goterm.Color(content, goterm.RED))
		if err != nil {
			log.Fatal(err)
		}
	}
	goterm.Flush()
	return
}

func printFixtureContent(totals *goterm.Table, m model.Matches, t model.Teams) (result string) {
	result = fmt.Sprintf("%s\t%s\t%s\n", m.Time.Label, m.Venue.Name, t.Name)
	_, err := fmt.Fprint(totals, result)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func printFixturesHeader(totals *goterm.Table) {
	_, err := fmt.Fprintf(totals, "TIME\tVENUE\tOPPONENT\n")
	if err != nil {
		log.Fatal(err)
	}
}
