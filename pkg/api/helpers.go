package api

import (
	"encoding/json"
	"net/url"
)

func (api *API) GetExpand(Url string) (*Expand, error) {
	var expand Expand

	urlPath := "v3/expand/?access_token=" + api.token + "&format=json&shortUrl=" + url.QueryEscape(Url)
	body, err := api.GetResponse(urlPath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &expand); err != nil {
		return nil, err
	}
	return &expand, nil
}

func (api *API) GetLinkSave(Url string) (*Expand, error) {
	var expand Expand

	urlPath := "v3/user/link_save?access_token=" + api.token + "&format=json&longUrl=" + url.QueryEscape(Url)
	body, err := api.GetResponse(urlPath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &expand); err != nil {
		return nil, err
	}
	return &expand, nil
}
