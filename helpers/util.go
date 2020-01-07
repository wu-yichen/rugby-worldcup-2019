package helpers

import (
	"log"

	"github.com/followme1987/wcwhen/model"

	"github.com/buger/goterm"

	"github.com/followme1987/wcwhen/api"
)

func RetrieveMatches(p api.Processor) ([]model.Matches, error) {
	data, err := api.GetMatches(p)
	if err != nil {
		return nil, err
	}
	matches := p.UnmarshalData(data)
	return matches, nil
}

func PrintTable(totals *goterm.Table) {
	_, err := goterm.Println(totals)
	if err != nil {
		log.Fatal(err)
	}
}
