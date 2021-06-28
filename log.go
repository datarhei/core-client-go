package client

import (
	"encoding/json"

	"github.com/datarhei/core-client-go/api"
)

func (r *restclient) Log() (api.Log, error) {
	var log api.Log

	data, err := r.call("GET", "/log", "", nil)
	if err != nil {
		return log, err
	}

	err = json.Unmarshal(data, &log)

	return log, err
}
