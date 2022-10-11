package coreclient

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/datarhei/core-client-go/v16/api"
)

func (r *restclient) Metadata(id, key string) (api.Metadata, error) {
	var m api.Metadata

	path := "/v3/process/" + url.PathEscape(id) + "/metadata"
	if len(key) != 0 {
		path += "/" + url.PathEscape(key)
	}

	data, err := r.call("GET", path, "", nil)
	if err != nil {
		return m, err
	}

	err = json.Unmarshal(data, &m)

	return m, err
}

func (r *restclient) MetadataSet(id, key string, metadata api.Metadata) error {
	var buf bytes.Buffer

	e := json.NewEncoder(&buf)
	e.Encode(metadata)

	_, err := r.call("PUT", "/v3/process/"+url.PathEscape(id)+"/metadata/"+url.PathEscape(key), "application/json", &buf)
	if err != nil {
		return err
	}

	return nil
}
