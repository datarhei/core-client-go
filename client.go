package coreclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/datarhei/core-client-go/api"

	"github.com/Masterminds/semver/v3"
)

const (
	coreapp     = "datarhei-core"
	coreversion = "^12.0.0"
	apiversion  = "^3.0.0"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RestClient interface {
	String() string
	ID() string
	Address() string

	About() api.About

	Config() (api.Config, error)
	ConfigSet(config api.ConfigData) error
	ConfigReload() error

	DataList(sort, order string) ([]api.FileInfo, error)
	DataHasFile(path string) bool
	DataDeleteFile(path string) error
	DataAddFile(path string, data io.Reader) error

	Log() (api.Log, error)

	MemFSList(sort, order string) ([]api.FileInfo, error)
	MemFSHasFile(path string) bool
	MemFSDeleteFile(path string) error
	MemFSAddFile(path string, data io.Reader) error

	Metadata(id, key string) (api.Metadata, error)
	MetadataSet(id, key string, metadata api.Metadata) error

	ProcessList(id, filter []string) ([]api.Process, error)
	Process(id string, filter []string) (api.Process, error)
	ProcessAdd(p api.ProcessConfig) error
	ProcessDelete(id string) error
	ProcessCommand(id, command string) error
	ProcessProbe(id string) (api.Probe, error)
	ProcessConfig(id string) (api.ProcessConfig, error)
	ProcessReport(id string) (api.ProcessReport, error)
	ProcessState(id string) (api.ProcessState, error)
	ProcessMetadata(id, key string) (api.Metadata, error)
	ProcessMetadataSet(id, key string, metadata api.Metadata) error

	Skills() (api.Skills, error)
	SkillsReload() error

	Sessions(collectors []string) (api.SessionsSummary, error)
	SessionsActive(collectors []string) (api.SessionsActive, error)
}

type Config struct {
	Address  string
	Username string
	Password string
	Client   HTTPClient
}

type restclient struct {
	address  string
	prefix   string
	token    string
	expire   time.Time
	username string
	password string
	client   HTTPClient
	about    api.About
}

func New(config Config) (RestClient, error) {
	r := &restclient{
		address:  config.Address,
		prefix:   "/api",
		username: config.Username,
		password: config.Password,
		client:   config.Client,
	}

	if r.client == nil {
		r.client = &http.Client{
			Timeout: 15 * time.Second,
		}
	}

	about, err := r.info()
	if err != nil {
		return nil, err
	}

	r.about = about

	if r.about.App != coreapp {
		return nil, fmt.Errorf("didn't receive the expected API response (got: %s, want: %s)", r.about.Name, coreapp)
	}

	if len(r.about.ID) == 0 {
		if err := r.login(); err != nil {
			return nil, err
		}
	}

	c, _ := semver.NewConstraint(coreversion)
	v, err := semver.NewVersion(r.about.Version.Number)
	if err != nil {
		return nil, err
	}

	if !c.Check(v) {
		return nil, fmt.Errorf("the core version (%s) is not supported (%s)", r.about.Version.Number, coreversion)
	}

	found := false
	for _, version := range r.about.Version.API {
		c, _ := semver.NewConstraint(apiversion)
		v, err := semver.NewVersion(version)
		if err != nil {
			return nil, err
		}

		if c.Check(v) {
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("this client (%s) doesn't support the availabe API versions: %s", apiversion, strings.Join(r.about.Version.API, ", "))
	}

	return r, nil
}

func (r restclient) String() string {
	return fmt.Sprintf("%s %s (%s) %s @ %s", r.about.Name, r.about.Version.Number, r.about.Version.Arch, r.about.ID, r.address)
}

func (r *restclient) ID() string {
	return r.about.ID
}

func (r *restclient) Address() string {
	return r.address
}

func (r *restclient) About() api.About {
	return r.about
}

func (r *restclient) login() error {
	login := api.Login{
		Username: r.username,
		Password: r.password,
	}

	var buf bytes.Buffer

	e := json.NewEncoder(&buf)
	e.Encode(login)

	req, err := http.NewRequest("POST", r.address+r.prefix+"/login", &buf)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	status, body, err := r.request(req)
	if err != nil {
		return err
	}

	if status != 200 {
		return fmt.Errorf("wrong username and/or password")
	}

	jwt := api.JWT{}

	json.Unmarshal(body, &jwt)

	r.token = jwt.Token
	r.expire.UnmarshalText([]byte(jwt.Expire))

	about, err := r.info()
	if err != nil {
		return err
	}

	if len(about.ID) == 0 {
		return fmt.Errorf("login to API failed")
	}

	r.about = about

	return nil
}

func (r *restclient) info() (api.About, error) {
	req, err := http.NewRequest("GET", r.address+r.prefix, nil)
	if err != nil {
		return api.About{}, err
	}

	if len(r.token) != 0 {
		req.Header.Add("Authorization", "Bearer "+r.token)
	}

	status, body, err := r.request(req)
	if err != nil {
		return api.About{}, err
	}

	if status != 200 {
		return api.About{}, fmt.Errorf("access to API failed (%d)", status)
	}

	about := api.About{}

	json.Unmarshal(body, &about)

	return about, nil
}

func (r *restclient) call(method, path, contentType string, data io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, r.address+r.prefix+"/v3"+path, data)
	if err != nil {
		return nil, err
	}

	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", contentType)
	}

	if len(r.token) != 0 {
		req.Header.Add("Authorization", "Bearer "+r.token)
	}

	status, body, err := r.request(req)
	if status == http.StatusUnauthorized {
		if err := r.login(); err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+r.token)
		status, body, err = r.request(req)
	}

	if err != nil {
		return nil, err
	}

	if status < 200 || status >= 300 {
		e := api.Error{}

		json.Unmarshal(body, &e)

		return nil, fmt.Errorf("%w", e)
	}

	return body, nil
}

func (r *restclient) request(req *http.Request) (int, []byte, error) {
	resp, err := r.client.Do(req)
	if err != nil {
		return -1, nil, err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return resp.StatusCode, body, nil
}
