package client

import (
	"encoding/json"
	"io"
	"net/url"

	"github.com/datarhei/core-client-go/api"
)

func (r *restclient) MemFSList(sort, order string) ([]api.FileInfo, error) {
	var files []api.FileInfo

	values := url.Values{}
	values.Set("sort", sort)
	values.Set("order", order)

	data, err := r.call("GET", "/memfs?"+values.Encode(), "", nil)
	if err != nil {
		return files, err
	}

	err = json.Unmarshal(data, &files)

	return files, err
}

func (r *restclient) MemFSHasFile(path string) bool {
	_, err := r.call("GET", "/memfs"+path, "", nil)

	return err == nil
}

func (r *restclient) MemFSDeleteFile(path string) error {
	_, err := r.call("DELETE", "/memfs"+path, "", nil)

	return err
}

func (r *restclient) MemFSAddFile(path string, data io.Reader) error {
	_, err := r.call("PUT", "/memfs"+path, "application/data", data)

	return err
}
