package model

type Rugby struct {
	Event struct {
		ID    int         `json:"id"`
		AltID interface{} `json:"altId"`
		Label string      `json:"label"`
		Sport string      `json:"sport"`
		Start struct {
			Millis    int64   `json:"millis"`
			GmtOffset float64 `json:"gmtOffset"`
			Label     string  `json:"label"`
		} `json:"start"`
		End struct {
			Millis    int64   `json:"millis"`
			GmtOffset float64 `json:"gmtOffset"`
			Label     string  `json:"label"`
		} `json:"end"`
		RankingsWeight float64     `json:"rankingsWeight"`
		Abbr           interface{} `json:"abbr"`
		WinningTeam    interface{} `json:"winningTeam"`
		ImpactPlayers  interface{} `json:"impactPlayers"`
	} `json:"event"`
	Matches []Matches `json:"matches"`
}
