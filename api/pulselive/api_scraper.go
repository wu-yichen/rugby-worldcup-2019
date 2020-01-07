package pulselive

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/followme1987/wcwhen/model"
)

type scraper struct {
	url string
}

func New() scraper {
	return scraper{
		url: os.Getenv("PULSELIVE_API_PATH"),
	}
}

func (s scraper) UnmarshalData(data []byte) []model.Matches {
	var content model.Rugby
	err := json.Unmarshal(data, &content)
	if err != nil {
		log.Fatal(err)
	}
	return content.Matches
}

func (s scraper) GetData() ([]byte, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return data, nil
}
