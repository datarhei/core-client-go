package coreclient

import (
	"encoding/json"
	"io"
	"net/url"

	"github.com/datarhei/core-client-go/v16/api"
)

func (r *restclient) MemFSList(sort, order string) ([]api.FileInfo, error) {
	var files []api.FileInfo

	values := url.Values{}
	values.Set("sort", sort)
	values.Set("order", order)

	data, err := r.call("GET", "/v3/fs/mem?"+values.Encode(), "", nil)
	if err != nil {
		return files, err
	}

	err = json.Unmarshal(data, &files)

	return files, err
}

func (r *restclient) MemFSHasFile(path string) bool {
	_, err := r.call("HEAD", "/v3/fs/mem"+path, "", nil)

	return err == nil
}

func (r *restclient) MemFSGetFile(path string) (io.ReadCloser, error) {
	return r.stream("GET", "/v3/fs/mem"+path, "", nil)
}

func (r *restclient) MemFSDeleteFile(path string) error {
	_, err := r.call("DELETE", "/v3/fs/mem"+path, "", nil)

	return err
}

func (r *restclient) MemFSAddFile(path string, data io.Reader) error {
	_, err := r.call("PUT", "/v3/fs/mem"+path, "application/data", data)

	return err
}
