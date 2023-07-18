package coreclient

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"

	"github.com/datarhei/core-client-go/v16/api"
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

func (r *restclient) FilesystemList(name, pattern, sort, order string) ([]api.FileInfo, error) {
	var files []api.FileInfo

	query := &url.Values{}
	query.Set("glob", pattern)
	query.Set("sort", sort)
	query.Set("order", order)

	data, err := r.call("GET", "/v3/fs/"+url.PathEscape(name), query, nil, "", nil)
	if err != nil {
		return files, err
	}

	err = json.Unmarshal(data, &files)

	return files, err
}

func (r *restclient) FilesystemHasFile(name, path string) bool {
	if !filepath.IsAbs(path) {
		path = "/" + path
	}

	_, err := r.call("HEAD", "/v3/fs/"+url.PathEscape(name)+path, nil, nil, "", nil)

	return err == nil
}

func (r *restclient) FilesystemGetFile(name, path string) (io.ReadCloser, error) {
	return r.FilesystemGetFileOffset(name, path, 0)
}

type ContextReadCloser struct {
	io.ReadCloser
	cancel context.CancelFunc
}

func NewContextReadCloser(r io.ReadCloser, cancel context.CancelFunc) *ContextReadCloser {
	return &ContextReadCloser{
		ReadCloser: r,
		cancel:     cancel,
	}
}

func (r *ContextReadCloser) Close() error {
	r.cancel()
	return r.ReadCloser.Close()
}

func (r *restclient) FilesystemGetFileOffset(name, path string, offset int64) (io.ReadCloser, error) {
	if !filepath.IsAbs(path) {
		path = "/" + path
	}

	var header http.Header = nil

	if offset > 0 {
		header = make(http.Header)
		header.Set("Range", "bytes="+strconv.FormatInt(offset, 10)+"-")
	}

	return r.stream(context.Background(), "GET", "/v3/fs/"+url.PathEscape(name)+path, nil, header, "", nil)
}

func (r *restclient) FilesystemDeleteFile(name, path string) error {
	if !filepath.IsAbs(path) {
		path = "/" + path
	}

	_, err := r.call("DELETE", "/v3/fs/"+url.PathEscape(name)+path, nil, nil, "", nil)

	return err
}

func (r *restclient) FilesystemAddFile(name, path string, data io.Reader) error {
	if !filepath.IsAbs(path) {
		path = "/" + path
	}

	_, err := r.call("PUT", "/v3/fs/"+url.PathEscape(name)+path, nil, nil, "application/data", data)

	return err
}
