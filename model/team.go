package model

type Teams struct {
	ID           int         `json:"id"`
	AltID        interface{} `json:"altId"`
	Name         string      `json:"name"`
	Abbreviation string      `json:"abbreviation"`
	Annotations  interface{} `json:"annotations"`
}
