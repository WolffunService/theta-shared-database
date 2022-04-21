package playerstatsmodel

type PlayerGetInt64Response struct {
	Success bool    `json:"success"`
	Data    int64   `json:"data"`
	Status  string  `json:"status"`
}

type PlayerGetFloat64Response struct {
	Success bool    `json:"success"`
	Data    float64 `json:"data"`
	Status  string  `json:"status"`
}

type PlayerTrackStatResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Status  string      `json:"status"`
}

type TrackStatsRequest struct {
	UserId    string                  `json:"userId"`
	StatName  string                  `json:"statName"`
	StatValue int64                   `json:"statValue"`
}