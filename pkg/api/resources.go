package api

type Answer struct {
	StatusCode int    `json:"status_code"`
	StatusTxt  string `json:"status_txt"`
}

type Expand struct {
	Answer
	Data struct {
		Expand []struct {
			ShortURL   string `json:"short_url"`
			LongURL    string `json:"long_url"`
			UserHash   string `json:"user_hash"`
			GlobalHash string `json:"global_hash"`
			Error      string `json:"error"`
		} `json:"expand"`
		LinkSave struct {
			Link          string `json:"link"`
			AggregateLink string `json:"aggregate_link"`
			LongURL       string `json:"long_url"`
			NewLink       int    `json:"new_link"`
		} `json:"link_save"`
	} `json:"data"`
}
