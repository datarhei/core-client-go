package coreclient

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/datarhei/core-client-go/v16/api"
)

func (r *restclient) IdentitiesList() ([]api.IAMUser, error) {
	var users []api.IAMUser

	data, err := r.call("GET", "/v3/iam/user", nil, nil, "", nil)
	if err != nil {
		return users, err
	}

	err = json.Unmarshal(data, &users)

	return users, err
}

func (r *restclient) Identity(name string) (api.IAMUser, error) {
	var user api.IAMUser

	data, err := r.call("GET", "/v3/iam/user/"+url.PathEscape(name), nil, nil, "", nil)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(data, &user)

	return user, err
}

func (r *restclient) IdentityAdd(u api.IAMUser) error {
	var buf bytes.Buffer

	e := json.NewEncoder(&buf)
	e.Encode(u)

	_, err := r.call("POST", "/v3/iam/user", nil, nil, "application/json", &buf)
	if err != nil {
		return err
	}

	return nil
}

func (r *restclient) IdentityUpdate(name string, u api.IAMUser) error {
	var buf bytes.Buffer

	e := json.NewEncoder(&buf)
	e.Encode(u)

	_, err := r.call("PUT", "/v3/iam/user/"+url.PathEscape(name), nil, nil, "application/json", &buf)
	if err != nil {
		return err
	}

	return nil
}

func (r *restclient) IdentitySetPolicies(name string, p []api.IAMPolicy) error {
	var buf bytes.Buffer

	e := json.NewEncoder(&buf)
	e.Encode(p)

	_, err := r.call("PUT", "/v3/iam/user/"+url.PathEscape(name)+"/policy", nil, nil, "application/json", &buf)
	if err != nil {
		return err
	}

	return nil
}

func (r *restclient) IdentityDelete(name string) error {
	_, err := r.call("DELETE", "/v3/iam/user"+url.PathEscape(name), nil, nil, "", nil)

	return err
}
