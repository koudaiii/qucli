package quay

import (
	"fmt"
	"net/url"
	"os"
)

type QuayPermission struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type QuayPermissions struct {
	Items []QuayPermission
}

type QuayPermissionsResponse struct {
	Items map[string]interface{} `json:"permissions"`
}

type QuayRepositories struct {
	Items []ResponseRepository
}

type ResponseRepositories struct {
	Items []map[string]interface{} `json:"repositories"`
}

type QuayRepository struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

type ResponseRepository struct {
	Namespace   string `json:"namespace"`
	IsPublic    bool   `json:"is_public"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RequestRepository struct {
	Namespace   string `json:"namespace"`
	Visibility  string `json:"visibility"`
	Repository  string `json:"repository"`
	Description string `json:"description"`
}

func QuayURLParse(hostname string) *url.URL {
	u, err := url.Parse("https://" + hostname + "/api/v1/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	return u
}
