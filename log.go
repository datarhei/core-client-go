package coreclient

import (
	"encoding/json"

	"github.com/datarhei/core-client-go/v14/api"
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
