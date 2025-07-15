package coreclient

import (
	"bytes"
	"net/url"

	"encoding/json"

	"github.com/datarhei/core-client-go/v16/api"
)

func (r *restclient) identitiesList(where, domain string) ([]api.IAMUser, error) {
	var users []api.IAMUser

	path := "/v3/iam/user"
	if where == "cluster" {
		path = "/v3/cluster/iam/user"
	}

	query := &url.Values{}
	query.Set("domain", domain)

	data, err := r.call("GET", path, query, nil, "", nil)
	if err != nil {
		return users, err
	}

	err = json.Unmarshal(data, &users)

	return users, err
}

func (r *restclient) identity(where, name, domain string) (api.IAMUser, error) {
	var user api.IAMUser

	path := "/v3/iam/user/" + url.PathEscape(name)
	if where == "cluster" {
		path = "/v3/cluster/iam/user/" + url.PathEscape(name)
	}

	query := &url.Values{}
	query.Set("domain", domain)

	data, err := r.call("GET", path, query, nil, "", nil)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(data, &user)

	return user, err
}

func (r *restclient) identityAdd(where, domain string, u api.IAMUser) error {
	var buf bytes.Buffer

	path := "/v3/iam/user"
	if where == "cluster" {
		path = "/v3/cluster/iam/user"
	}

	query := &url.Values{}
	query.Set("domain", domain)

	e := json.NewEncoder(&buf)
	e.Encode(u)

	_, err := r.call("POST", path, query, nil, "application/json", &buf)
	if err != nil {
		return err
	}

	return nil
}

func (r *restclient) identityUpdate(where, name, domain string, u api.IAMUser) error {
	var buf bytes.Buffer

	path := "/v3/iam/user/" + url.PathEscape(name)
	if where == "cluster" {
		path = "/v3/cluster/iam/user/" + url.PathEscape(name)
	}

	query := &url.Values{}
	query.Set("domain", domain)

	e := json.NewEncoder(&buf)
	e.Encode(u)

	_, err := r.call("PUT", path, query, nil, "application/json", &buf)
	if err != nil {
		return err
	}

	return nil
}

func (r *restclient) identitySetPolicies(where, name, domain string, p []api.IAMPolicy) error {
	var buf bytes.Buffer

	path := "/v3/iam/user/" + url.PathEscape(name) + "/policy"
	if where == "cluster" {
		path = "/v3/cluster/iam/user/" + url.PathEscape(name) + "/policy"
	}

	query := &url.Values{}
	query.Set("domain", domain)

	e := json.NewEncoder(&buf)
	e.Encode(p)

	_, err := r.call("PUT", path, query, nil, "application/json", &buf)
	if err != nil {
		return err
	}

	return nil
}

func (r *restclient) identityDelete(where, name, domain string) error {
	path := "/v3/iam/user/" + url.PathEscape(name)
	if where == "cluster" {
		path = "/v3/cluster/iam/user/" + url.PathEscape(name)
	}

	query := &url.Values{}
	query.Set("domain", domain)

	_, err := r.call("DELETE", path, query, nil, "", nil)

	return err
}

func (r *restclient) IdentitiesList(domain string) ([]api.IAMUser, error) {
	return r.identitiesList("", domain)
}

func (r *restclient) Identity(name, domain string) (api.IAMUser, error) {
	return r.identity("", name, domain)
}

func (r *restclient) IdentityAdd(domain string, u api.IAMUser) error {
	return r.identityAdd("", domain, u)
}

func (r *restclient) IdentityUpdate(name, domain string, u api.IAMUser) error {
	return r.identityUpdate("", name, domain, u)
}

func (r *restclient) IdentitySetPolicies(name, domain string, p []api.IAMPolicy) error {
	return r.identitySetPolicies("", name, domain, p)
}

func (r *restclient) IdentityDelete(name, domain string) error {
	return r.identityDelete("", name, domain)
}
