package coreclient

import "github.com/datarhei/core-client-go/v16/api"

func (r *restclient) ClusterIdentitiesList(domain string) ([]api.IAMUser, error) {
	return r.identitiesList("cluster", domain)
}

func (r *restclient) ClusterIdentity(name, domain string) (api.IAMUser, error) {
	return r.identity("cluster", name, domain)
}

func (r *restclient) ClusterIdentityAdd(domain string, u api.IAMUser) error {
	return r.identityAdd("cluster", domain, u)
}

func (r *restclient) ClusterIdentityUpdate(name, domain string, u api.IAMUser) error {
	return r.identityUpdate("cluster", name, domain, u)
}

func (r *restclient) ClusterIdentitySetPolicies(name, domain string, p []api.IAMPolicy) error {
	return r.identitySetPolicies("cluster", name, domain, p)
}

func (r *restclient) ClusterIdentityDelete(name, domain string) error {
	return r.identityDelete("cluster", name, domain)
}

func (r *restclient) ClusterIAMReload() error {
	_, err := r.call("PUT", "/v3/cluster/iam/reload", nil, nil, "", nil)

	return err
}
