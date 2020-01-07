package api

import "github.com/followme1987/wcwhen/model"

type Processor interface {
	GetData() ([]byte, error)
	UnmarshalData(data []byte) []model.Matches
}

func GetMatches(p Processor) ([]byte, error) {
	data, err := p.GetData()
	if err != nil {
		return nil, err
	}
	return data, nil
}
