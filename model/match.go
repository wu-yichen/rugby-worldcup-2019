package model

type Matches struct {
	MatchID     int         `json:"matchId"`
	AltID       interface{} `json:"altId"`
	Description string      `json:"description"`
	EventPhase  string      `json:"eventPhase"`
	Venue       struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"venue"`
	Time struct {
		Millis    int64   `json:"millis"`
		GmtOffset float64 `json:"gmtOffset"`
		Label     string  `json:"label"`
	} `json:"time"`
	Attendance  int         `json:"attendance"`
	Teams       []Teams     `json:"teams"`
	Scores      []int       `json:"scores"`
	Kc          interface{} `json:"kc"`
	Status      string      `json:"status"`
	Clock       interface{} `json:"clock"`
	Outcome     string      `json:"outcome"`
	Events      interface{} `json:"events"`
	Sport       string      `json:"sport"`
	Competition string      `json:"competition"`
	Weather     struct {
		MatchWeather         interface{} `json:"matchWeather"`
		MatchMinTemperature  interface{} `json:"matchMinTemperature"`
		MatchMaxTemperature  interface{} `json:"matchMaxTemperature"`
		MatchWindConditions  interface{} `json:"matchWindConditions"`
		MatchPitchConditions interface{} `json:"matchPitchConditions"`
	} `json:"weather"`
}
