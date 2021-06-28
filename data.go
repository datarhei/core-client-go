package client

import (
	"encoding/json"
	"io"
	"net/url"

	"github.com/datarhei/core-client-go/api"
)

const (
	SORT_DEFAULT  = "none"
	SORT_NONE     = "none"
	SORT_NAME     = "name"
	SORT_SIZE     = "size"
	SORT_LASTMOD  = "lastmod"
	ORDER_DEFAULT = "asc"
	ORDER_ASC     = "asc"
	ORDER_DESC    = "desc"
)

func (r *restclient) DataList(sort, order string) ([]api.FileInfo, error) {
	var files []api.FileInfo

	values := url.Values{}
	values.Set("sort", sort)
	values.Set("order", order)

	data, err := r.call("GET", "/data?"+values.Encode(), "", nil)
	if err != nil {
		return files, err
	}

	err = json.Unmarshal(data, &files)

	return files, err
}

func (r *restclient) DataHasFile(path string) bool {
	_, err := r.call("GET", "/data"+path, "", nil)

	return err == nil
}

func (r *restclient) DataDeleteFile(path string) error {
	_, err := r.call("DELETE", "/data"+path, "", nil)

	return err
}

func (r *restclient) DataAddFile(path string, data io.Reader) error {
	_, err := r.call("PUT", "/data"+path, "application/data", data)

	return err
}