package fakes

import "github.com/followme1987/wcwhen/model"

type FakeScraper struct {
	GetDataFunc       func() ([]byte, error)
	UnmarshalDataFunc func(data []byte) []model.Matches
}

func (f FakeScraper) UnmarshalData(data []byte) []model.Matches {
	if f.UnmarshalDataFunc == nil {
		panic("scraper func not defined")
	}
	return f.UnmarshalDataFunc(data)
}

func (f FakeScraper) GetData() ([]byte, error) {
	if f.GetDataFunc == nil {
		panic("scraper func not defined")
	}
	return f.GetDataFunc()
}
