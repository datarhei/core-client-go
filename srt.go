package coreclient

import (
	"encoding/json"

	"github.com/datarhei/core-client-go/v16/api"
)

func (r *restclient) SRTChannels() (api.SRTChannels, error) {
	var m api.SRTChannels

	data, err := r.call("GET", "/v3/srt", "", nil)
	if err != nil {
		return m, err
	}

	err = json.Unmarshal(data, &m)

	return m, err
}
