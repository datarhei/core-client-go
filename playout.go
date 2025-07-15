package coreclient

import (
	"bytes"
	"net/url"

	"encoding/json"

	"github.com/datarhei/core-client-go/v16/api"
)

func (r *restclient) PlayoutStatus(id ProcessID, inputID string) (api.PlayoutStatus, error) {
	var status api.PlayoutStatus

	path := "/v3/process/" + url.PathEscape(id.ID) + "/playout/" + url.PathEscape(inputID) + "/status"

	values := &url.Values{}
	values.Set("domain", id.Domain)

	data, err := r.call("GET", path, values, nil, "", nil)
	if err != nil {
		return status, err
	}

	err = json.Unmarshal(data, &status)

	return status, err
}

func (r *restclient) PlayoutReopen(id ProcessID, inputID string) (bool, error) {
	path := "/v3/process/" + url.PathEscape(id.ID) + "/playout/" + url.PathEscape(inputID) + "/reopen"

	values := &url.Values{}
	values.Set("domain", id.Domain)

	data, err := r.call("GET", path, values, nil, "", nil)
	if err != nil {
		return false, err
	}

	bytes.HasPrefix(data, []byte("OK"))

	return bytes.HasPrefix(data, []byte("OK")), nil
}
