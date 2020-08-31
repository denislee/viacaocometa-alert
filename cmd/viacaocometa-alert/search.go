package main

type Search struct {
	BaseURL       string
	Date          string
	OriginID      string
	DestinationID string
}

func (search *Search) GetURL() string {
	finalURL := search.BaseURL
	finalURL += "?data=" + search.Date
	finalURL += "&origem=" + search.OriginID
	finalURL += "&destino=" + search.DestinationID
	return finalURL
}
