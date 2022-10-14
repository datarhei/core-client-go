package coreclient

import (
	"bytes"
	"encoding/json"

	"github.com/datarhei/core-client-go/v16/api"
)

type configVersion struct {
	Config struct {
		Version int64 `json:"version"`
	} `json:"config"`
}

func (r *restclient) Config() (int64, api.Config, error) {
	version := configVersion{}

	data, err := r.call("GET", "/v3/config", "", nil)
	if err != nil {
		return 0, api.Config{}, err
	}

	if err := json.Unmarshal(data, &version); err != nil {
		return 0, api.Config{}, err
	}

	config := api.Config{}

	if err := json.Unmarshal(data, &config); err != nil {
		return 0, api.Config{}, err
	}

	return version.Config.Version, config, nil
}

func (r *restclient) ConfigSet(config interface{}) error {
	var buf bytes.Buffer

	e := json.NewEncoder(&buf)
	e.Encode(config)

	_, err := r.call("PUT", "/v3/config", "application/json", &buf)

	return err
}

func (r *restclient) ConfigReload() error {
	_, err := r.call("GET", "/v3/config/reload", "", nil)

	return err
}
